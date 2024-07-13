package main

import (
	"fmt"
	"tap2crypto-backend/walletFunctions"
	"tap2crypto-backend/walletTransactions"
)

func main() {
	//initialize the network
	walletFunctions.InitNetwork("https://cloudflare-eth.com")

	//get the balance of the wallet
	balance, ethValue := walletFunctions.WalletBalance("0x3aa2263480c9d84d4bf5fa8831de88200be898be")
	fmt.Println("Balance: ", balance)
	fmt.Println("Balance in Ether: ", ethValue)
	//here we can either create a new wallet or import an existing wallet
	//publicKey, privateKey := createWallet()
	//publicKey, privateKey := importWallet()
	//fmt.Println("Public Address: ", publicKey)
	//fmt.Println("Private Key: ", privateKey)
}
