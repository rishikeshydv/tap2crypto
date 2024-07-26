from flask import Flask, request, jsonify
from flask_cors import CORS
from waitress import serve
from predictions import Headline
import numpy as np
app = Flask(__name__)
CORS(app)

@app.route("/api/vi/get-news",methods=['GET'])
def get_news():
    headline = Headline()
    news = headline.return_headlines()
    return jsonify(news)

@app.route('/api/v1/get-companies',methods=['GET'])
def get_companies():
    headline = Headline()
    headline.return_headlines()
    headline.preprocess()
    headline.predict()
    headline.return_entities()
    res_dict = headline.return_only_companies()
    # Convert int64 to int
    #res_dict = {k: int(v) if isinstance(v, np.int64) else v for k, v in res_dict.items()}
    return jsonify(res_dict)

if __name__ == '__main__':
        #debug mode
        app.run(debug=True, host='localhost', port=8080)
        #production mode
        #  app.logger.info("Starting the server...")
        #  serve(app, host="0.0.0.0", port=8080)
