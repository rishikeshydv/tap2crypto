package walletTransactions

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Configurations struct {
	Nonce    uint64
	GasPrice *big.Int
	ChainID  *big.Int
	GasLimit uint64
}

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
