package main

import (
	"fmt"
	"math/big"
	"tap2crypto-backend/optionsFunctions"
	"tap2crypto-backend/walletFunctions"
	"tap2crypto-backend/walletTransactions"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func walletFunc() {
	//initialize the network
	walletFunctions.InitNetwork("https://cloudflare-eth.com")
	walletTransactions.InitNetwork("https://cloudflare-eth.com")

	//get the balance of the wallet
	balance, ethValue := walletFunctions.WalletBalance("0x3aa2263480c9d84d4bf5fa8831de88200be898be")
	fmt.Println("Balance: ", balance)
	fmt.Println("Balance in Ether: ", ethValue)
	//here we can either create a new wallet or import an existing wallet
	//publicKey, privateKey := createWallet()
	//publicKey, privateKey := importWallet()
	//fmt.Println("Public Address: ", publicKey)
	//fmt.Println("Private Key: ", privateKey)

	//config transactions
	var config walletTransactions.Configurations
	fromPublicKey := common.HexToAddress("0x3aA2263480c9d84D4BF5fa8831De88200be898BE")
	fromPrivateKey, err := crypto.HexToECDSA("47010d969c59e1e2ea43a9d2977df71b2c856909a1988164dc2f589578837cca")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	toPublicKey := common.HexToAddress("0xcf865d5114f8d9079d69ac54cf2e8bcd8eb5fccb")
	var weiVal = big.Int{}
	config = walletTransactions.TransferConfig(fromPublicKey)
	fmt.Println("Config: ", config)
	weiVal = *walletTransactions.ToWei(big.NewFloat(0.1))
	fmt.Println("Wei Value: ", weiVal)

	//send the ethers
	walletTransactions.SendEthers(fromPublicKey, fromPrivateKey, toPublicKey, big.NewFloat(0.1))
}

func optionsFunc() {
	//betting functions
	bettingFunctions.Bet()
}
func main() {

	//walletFunc()
	optionsFunc()

}
