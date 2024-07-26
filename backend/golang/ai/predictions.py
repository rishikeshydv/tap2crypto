#scraping modules
from bs4 import BeautifulSoup
import requests
import spacy
#import joblib
from openai import OpenAI
import numpy as np
import pandas as pd
import pickle
class Headline:

    def __init__(self):
        self.headlines = []
        self.preprocessed_headlines = []
        #self.loaded_knn_model = joblib.load('knn_model.pkl')
        self.loaded_knn_model = pickle.load(open('backend/golang/ai/knn_model.pkl', 'rb'))
        self.predictions_knn = []
        self.knn_res = {}
        self.new_knn_res = {}
        self.nlp = spacy.load('en_core_web_sm')
        self.client = OpenAI(api_key="your-api-key")

    def return_headlines(self):
        #retrieve the content of the website
        content = requests.get('https://finance.yahoo.com/')
        content = content.text
        soup = BeautifulSoup(content, 'html.parser')

        #retrieve the information from "hero-headlines hero-latest-news yf-13r5oof" class of the div tag
        self.headlines = soup.find_all('div', class_='hero-headlines hero-second-col yf-13r5oof')
        self.headlines = self.headlines[0].find_all('h3')
        self.headlines = [headline.text for headline in self.headlines]
        return self.headlines

    def create_embedding(self,text):
        response = self.client.embeddings.create(
            input=text,
            model="text-embedding-3-small"
        )
        return response.data[0].embedding

    def preprocess(self):
        for headline in self.headlines[1:]:
            doc = self.nlp(headline)
            preprocessed_headline = [token.lemma_ for token in doc if not token.is_stop and not token.is_punct]
            self.preprocessed_headlines.append(preprocessed_headline)

        self.preprocessed_headlines = [self.create_embedding(headline) for headline in self.preprocessed_headlines]
        self.preprocessed_headlines = np.stack(self.preprocessed_headlines)

    def predict(self):
        self.predictions_knn = self.loaded_knn_model.predict(self.preprocessed_headlines)

    def return_entities(self):
        for i, headline in enumerate(self.headlines[1:]):
            doc = self.nlp(headline)
            for ent in doc.ents:
                self.knn_res[ent.text] = self.predictions_knn[i]
        return self.knn_res
    
    def return_only_companies(self):
        df_sp500 = pd.read_csv('backend/golang/ai/sp500.csv')
        df_sp500.head()
        #using only Symbol and Security
        df_sp500 = df_sp500[['Symbol', 'Security']]
        df_sp500.head()
        df_sp500.dropna(inplace=True)

        for entity, sentiment in self.knn_res.items():
            for i, row in df_sp500.iterrows():
                if entity in row.Security:
                    self.new_knn_res[row.Symbol] = "High" if sentiment == 1 else "Low"

        return self.new_knn_res
    
#instanticate the class
headline = Headline()
print(headline.return_headlines())
headline.preprocess()
headline.predict()
headline.return_entities()
print(headline.return_only_companies())

    

