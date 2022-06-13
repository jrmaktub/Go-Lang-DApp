package main

import (
	// "context"
	"fmt"
	"log"

	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	//connecting to the Ethereum Network using Infura node
	// conn, err := ethclient.Dial("https://mainnet.infura.io/v3/2ea67de6c0fe4745b93a7bf99ba89c86")

	// if err != nil {
	// 	log.Fatal("Failed to connect to Eth node", err)
	// }

	//create context which allows to create communication between the concurrent Eth node communication  and the main application
	// cntxt := context.Background()

	//using a combination of the connection established to query the blockchain for transactions by  a specific hasg
	// txn, pending, _ := conn.TransactionByHash(cntxt, common.HexToHash("0xa533e1aa8d96fad7ce21397e28b5c19ea9c646a0f28ef2e639141aa06b94d7d8"))

	// if pending {
	// 	fmt.Println("Txn is pending", txn)
	// } else {
	// 	fmt.Println("Txn is not pending", txn)
	// }

	//reading wallet balannce
	// account := common.HexToAddress("0x4CE66b68404Bb3775202C1F78a87150E925B53C2")

	// balance, err := conn.BalanceAt(context.Background(), account, nil)

	// if err != nil {
	// 	log.Fatal("unable to get balance")
	// } else {
	// 	fmt.Println(balance)
	// }

	//creating wallets
	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal("unable to generate new keys")
	}

	//convert the privatekey from ECDSA format to Bytes format
	privateKeyBytes := crypto.FromECDSA(privateKey)
	//hexadecimal private key
	fmt.Println(hexutil.Encode(privateKeyBytes))

	//creating a new wallet, which converts it into Bytes and then into Hexadecimal
}
