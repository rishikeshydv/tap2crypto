"use client"
import React from "react";
import LoggerNews from "@/components/LoggerNews";
import LoggerStocks from "@/components/LoggerStocks";
import KeyTable from "@/components/key-table";
import axios from "axios";
import { useState } from "react";
import { FaWallet } from "react-icons/fa";
import { MdImportExport } from "react-icons/md";
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"


export default function Home() {
  //retrieve the scraped stocks from the backend
  const [scrapedNews, setScrapedNews] = useState<string[]>([]);
  const [scrapedStocks, setScrapedStocks] = useState<any>({});
  const [stock, setStock] = useState<boolean>(false); 
  const [publicKey, setPublicKey] = useState<string>("");
  const [privateKey, setPrivateKey] = useState<string>("");
  const [walletCreation, setWalletCreation] = useState<boolean>(false);
  const [showImportWallet, setShowImportWallet] = useState<boolean>(false);
  const [importPrivateKey, setImportPrivateKey] = useState<string>("");



   async function getNews() {
    await axios.get("http://localhost:8080/api/vi/get-news")
    .then((response) => {
      setScrapedNews(response.data);
    })
   }

   async function getStocks() {
    await axios.get("http://localhost:8080/api/v1/get-companies")
    .then((response) => {
      setScrapedStocks(response.data);
    })
   }

   async function createWallet() {
    await axios.get("http://localhost:9000/api/v1/create-wallet")
    .then((response) => {
       setPublicKey(response.data.publicKey);
       setPrivateKey(response.data.privateKey);
    })
  }

  async function importWallet(){
    await axios.post("http://localhost:9000/api/v1/import-wallet", {
      privateKey: importPrivateKey as string
    })
    .then((response) => {
      setPublicKey(response.data.publicKey);
      setPrivateKey(response.data.privateKey);
      setWalletCreation(true);
    })
  }


  return (
<div>
  <div className="flex gap-4">
  <div className="flex">
  <button onClick={()=>{
    createWallet();
    setWalletCreation(true);
}} className="p-2 bg-gray-200">Create Wallet </button>
<FaWallet className="mt-3"/>
  </div>
  <div className="flex">
  <button onClick={()=>{
    setShowImportWallet(true);
}} className="p-2 bg-gray-200">Import Wallet </button>
<MdImportExport className="mt-3"/>
  </div>
</div>
{
  showImportWallet ? 
  <header className="flex flex-col items-center gap-4 p-6">
  <h2 className="text-2xl font-semibold">Import Wallet</h2>
  <div className="flex w-full max-w-sm items-center gap-2">
    <Input placeholder="Enter Private Key" className="flex-1" onChange={(e)=>setImportPrivateKey(e.target.value)}/>
    <Button onClick={()=>{
      importWallet();
      setShowImportWallet(true);
    }}>Import</Button>
  </div>
</header>
  : null
}
      {
        walletCreation ? <KeyTable publicKey={publicKey} privateKey={privateKey}/> : null
      }
  <div>
    <div>
    <button onClick={()=>{
      getNews();
      setStock(false);
      }} className="p-2 bg-gray-200">Get News</button>
    <button onClick={()=>{
      getStocks();
      setStock(true);
      }} className="p-2 bg-gray-200">Get Stocks</button>
    </div>
    {
      stock ? <LoggerStocks LoggerStatements={scrapedStocks} /> : <LoggerNews LoggerStatements={scrapedNews} />
    }
  </div>

      {
        walletCreation ? <KeyTable publicKey={publicKey} privateKey={privateKey}/> : null
      }
            <div>
        Enter Amount: <input type="text"/>
      </div>
  <div>
    <div>
    <button onClick={()=>{
      getNews();
      setStock(false);
      }} className="p-2 bg-gray-200">Get News</button>
    <button onClick={()=>{
      getStocks();
      setStock(true);
      }} className="p-2 bg-gray-200">Get Stocks</button>
    </div>
    {
      stock ? <LoggerStocks LoggerStatements={scrapedStocks} /> : <LoggerNews LoggerStatements={scrapedNews} />
    }
  </div>

</div>
  );
}
