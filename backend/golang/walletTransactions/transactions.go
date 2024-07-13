package walletTransactions

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
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

func TransferConfig(publicKey common.Address) Configurations {
	var config Configurations
	// get the nonce
	nonce, err := client.PendingNonceAt(context.Background(), publicKey)
	if err != nil {
		log.Fatal(err)
	}

	// get the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// get the chainID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	config.Nonce = nonce
	config.GasPrice = gasPrice
	config.ChainID = chainID
	config.GasLimit = uint64(21000)

	return config
}

func ToWei(amount *big.Float) *big.Int {
	wei := new(big.Int)
	amount.Mul(amount, big.NewFloat(1e18))
	amount.Int(wei)
	return wei
}

func Transaction (fromPublicAddress common.Address, fromPrivateAddress *ecdsa.PrivateKey, toPublicAddress common.Address, amount *big.Float) {
	weiAmount := ToWei(amount)
	config := TransferConfig(fromPublicAddress)
	tx := types.NewTransaction(config.Nonce, toPublicAddress, weiAmount, config.GasLimit, config.GasPrice,nil)
	signedTx,err := types.SignTx(tx,types.NewEIP155Signer(config.ChainID),fromPrivateAddress)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(),signedTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transfer: %s wei to %s", weiAmount.String(), toPublicAddress.String())
}