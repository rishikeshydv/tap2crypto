package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func createWallet() (common.Address, *ecdsa.PrivateKey) {
	generatedPrivKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	//get the private key
	//privatekeyBytes := crypto.FromECDSA(generatedPrivKey)
	//log.Print("Private Key Bytes: ", privatekeyBytes)
	//privateKey := hexutil.Encode(privatekeyBytes)[2:]
	//log.Print("Private Key: ", privateKey)

	//get the public key
	publicKey := generatedPrivKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return publicAddress, generatedPrivKey
}
func main() {
	publicKey, privateKey := createWallet()
	fmt.Println("Public Address: ", publicKey)
	fmt.Println("Private Key: ", privateKey)
}
