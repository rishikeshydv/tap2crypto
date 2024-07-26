"use client"
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card"
import { Table, TableHeader, TableRow, TableHead, TableBody, TableCell } from "@/components/ui/table"
import { Button } from "@/components/ui/button"
import { BiSolidHide } from "react-icons/bi";
import React from "react";
import { useState } from "react";
import axios from "axios";
const KeyTable = ({publicKey,privateKey}:{publicKey:string,privateKey:string}) => {
  const [balance, setBalance] = useState<string>("");
  const [eth, setEth] = useState<string>("");

  async function checkBalance() {
    await axios.post("http://localhost:9000/api/v1/get-balance",{
      publicKey: publicKey
    })
    .then((response) => {
      setBalance(response.data.balance);
      setEth(response.data.ethValue);
    })
  }
  return (
    <Card>
      <CardHeader>
        <CardTitle>Your Keys</CardTitle>
        <CardDescription>Manage your public and private keys.</CardDescription>
      </CardHeader>
      <CardContent>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Type</TableHead>
              <TableHead>Key</TableHead>
              <TableHead className="text-right">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell>Public</TableCell>
              <TableCell>
                <div className="truncate font-mono text-sm">{publicKey}</div>
              </TableCell>
              <TableCell className="text-right">
                <Button variant="outline" size="icon">
                  <BiSolidHide className="w-4 h-4" />
                  <span className="sr-only">hide public key</span>
                </Button>
              </TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Private</TableCell>
              <TableCell>
                <div className="truncate font-mono text-sm">{privateKey}</div>
              </TableCell>
              <TableCell className="text-right">
                <Button variant="outline" size="icon">
                  <BiSolidHide className="w-4 h-4" />
                  <span className="sr-only">hide private key</span>
                </Button>
              </TableCell>
            </TableRow>
          </TableBody>
          <button className="font-bold text-xl border-1" onClick={checkBalance}>Check Balance</button>
          <div>
          <p>Balance: {balance}</p>
          <p>Eth Value: {eth}</p>
          </div>
          
        </Table>
      </CardContent>
    </Card>
  )
}

export default KeyTable