package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"tap2crypto-backend/walletFunctions"
	"tap2crypto-backend/walletTransactions"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func createWallet(w http.ResponseWriter, r *http.Request) {
	//initialize the network
	walletFunctions.InitNetwork("https://cloudflare-eth.com")
	publicKey, privateKey := walletFunctions.CreateWallet()
	res := map[string]string{
		"publicKey":  (publicKey).String(),
		"privateKey": privateKey,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func importWallet(w http.ResponseWriter, r *http.Request) {
	//initialize the network
	walletFunctions.InitNetwork("https://cloudflare-eth.com")
	//getting the privateKey from the request
	//the request body has privateKey_: string
	var req struct {
		PrivateKey string `json:"privateKey"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)
	privateKey := req.PrivateKey
	publicKey, importedPrivateKey := walletFunctions.ImportWallet(privateKey)
	w.Header().Set("Content-Type", "application/json")
	res := map[string]string{
		"publicKey":          (publicKey).String(),
		"privateKey":         privateKey,
		"importedPrivateKey": importedPrivateKey.D.String(),
	}
	json.NewEncoder(w).Encode(res)
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	//initialize the network
	walletFunctions.InitNetwork("https://cloudflare-eth.com")
	//getting the publicKey from the request
	var req struct {
		PublicKey string `json:"publicKey"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)
	publicKey := req.PublicKey
	balance, ethValue := walletFunctions.WalletBalance(publicKey)
	w.Header().Set("Content-Type", "application/json")
	res := map[string]string{
		"balance":  balance.String(),
		"ethValue": ethValue.String(),
	}
	json.NewEncoder(w).Encode(res)
}

func walletFunc(w http.ResponseWriter, r *http.Request) {
	var publicKey string
	var privateKey string
	var toPublicKey_ string

	_ = json.NewDecoder(r.Body).Decode(&publicKey)
	_ = json.NewDecoder(r.Body).Decode(&privateKey)
	_ = json.NewDecoder(r.Body).Decode(&toPublicKey_)
	//initialize the network
	walletFunctions.InitNetwork("https://cloudflare-eth.com")
	walletTransactions.InitNetwork("https://cloudflare-eth.com")

	//config transactions
	var config walletTransactions.Configurations
	fromPublicKey := common.HexToAddress(publicKey)
	fromPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	toPublicKey := common.HexToAddress(toPublicKey_)
	var weiVal = big.Int{}
	config = walletTransactions.TransferConfig(fromPublicKey)
	fmt.Println("Config: ", config)
	weiVal = *walletTransactions.ToWei(big.NewFloat(0.1))
	fmt.Println("Wei Value: ", weiVal)

	//send the ethers
	walletTransactions.SendEthers(fromPublicKey, fromPrivateKey, toPublicKey, big.NewFloat(0.1))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Transaction Successful")
}

func main() {

	//handling requests
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/api/v1/create-wallet", createWallet).Methods("GET")
	myRouter.HandleFunc("/api/v1/import-wallet", importWallet).Methods("POST")
	myRouter.HandleFunc("/api/v1/get-balance", getBalance).Methods("POST")
	myRouter.HandleFunc("/api/v1/send-crypto", walletFunc).Methods("POST")

	// Create a CORS handler with the desired options
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3009"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"X-Total-Count"},
		AllowCredentials: true,
	})
	handler := c.Handler(myRouter)
	fmt.Printf("Server started at http://localhost:9000")
	log.Fatal(http.ListenAndServe(":9000", handler))

}
