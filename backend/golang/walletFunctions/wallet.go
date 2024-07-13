package walletFunctions

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	//"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// defining the client as a global variable
var client *ethclient.Client

func InitNetwork(networkRPC string) {
	var err error
	client, err = ethclient.Dial(networkRPC)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the network")
}

func CreateWallet() (common.Address, *ecdsa.PrivateKey) {
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

func ImportWallet(privateKey string) (common.Address, *ecdsa.PrivateKey) {
	importedPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := importedPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return publicAddress, importedPrivateKey
}

func WalletBalance(address string) (*big.Int, *big.Float) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	balanceInEth := new(big.Float)
	balanceInEth.SetString(balance.String())
	//Finally, I multiply this float value by 10^18. This is because the smallest unit of Ether is called Wei and 1 Ether = 10^18 Wei.
	ethValue := new(big.Float).Quo(balanceInEth, big.NewFloat(1e18))
	return balance, ethValue
}
