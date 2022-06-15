// package main

// import (
// 	// "context"
// 	"fmt"
// 	"log"

// 	// "github.com/ethereum/go-ethereum/common"
// 	// "github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/ethereum/go-ethereum/common/hexutil"
// 	"github.com/ethereum/go-ethereum/crypto"
// )

// func main() {
//connecting to the Ethereum Network using Infura node
// conn, err := ethclient.Dial("https://mainnet.infura.io/v3/2ea67de6c0fe4745b93a7bf99ba89c86")

// if err != nil {
// 	log.Fatal("Failed to connect to Eth node", err)
// }

//1) create context which allows to create communication between the concurrent Eth node communication  and the main application
// cntxt := context.Background()

//using a combination of the connection established to query the blockchain for transactions by  a specific hash
// txn, pending, _ := conn.TransactionByHash(cntxt, common.HexToHash("0xa533e1aa8d96fad7ce21397e28b5c19ea9c646a0f28ef2e639141aa06b94d7d8"))

// if pending {
// 	fmt.Println("Txn is pending", txn)
// } else {
// 	fmt.Println("Txn is not pending", txn)
// }

// 2) reading wallet balannce
// account := common.HexToAddress("0x4CE66b68404Bb3775202C1F78a87150E925B53C2")

// balance, err := conn.BalanceAt(context.Background(), account, nil)

// if err != nil {
// 	log.Fatal("unable to get balance")
// } else {
// 	fmt.Println(balance)
// }

//3) creating wallets
// privateKey, err := crypto.GenerateKey()

// if err != nil {
// 	log.Fatal("unable to generate new keys")
// }

// 	//convert the privatekey from ECDSA format to Bytes format
// 	privateKeyBytes := crypto.FromECDSA(privateKey)

// 	//hexadecimal private key
// 	fmt.Println(hexutil.Encode(privateKeyBytes))

// 	 creating a new wallet, which converts it into Bytes and then into Hexadecimal
// }

//4) Validate Ethereum Addresses

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"regexp"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func main() {
// 	//this expression is checking for a string that starts with 0x, and that contains numeric values from  0-9, and letters that are capital or lowercase from A-F, and also making sure that character length is only 40 after the 0x
// 	regExpression := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

// 	fmt.Println("This is a valid address", regExpression.MatchString("0x4CE66b68404Bb3775202C1F78a87150E925B53C2"))
// 	fmt.Println("This is an invalid address", regExpression.MatchString("0xZCrt66b68404Bb3775202C1F78a87150E925B53C2"))

// 	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2ea67de6c0fe4745b93a7bf99ba89c86")

// 	if err != nil {
// 		log.Fatal("Unable to reach client")
// 	}

// 	//uniswap contract
// 	address := common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")

// CodeAt checks if there is code at given address
// 	bytecode, err := client.CodeAt(context.Background(), address, nil)

// 	if err != nil {
// 		log.Fatal("Unable to check if there's code", address, err)
// 	}

// 	//boolean. If lens is greater than one, the address has code
// 	isContract := len(bytecode) > 0

// 	fmt.Println(address, "is contract", isContract)

// }

// 5) Working with Ethereum blocks
// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math/big"

// 	// "github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	// "github.com/ethereum/go-ethereum/common"
// 	// "context"
// 	// "fmt"
// 	// "log"
// 	// "regexp"
// 	// "github.com/ethereum/go-ethereum/common"
// 	// "github.com/ethereum/go-ethereum/ethclient"
// )

// func main() {

// 	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2ea67de6c0fe4745b93a7bf99ba89c86")

// 	if err != nil {
// 		log.Fatal("Unable to reach client")
// 	}

// 	//get the header of a block by number. Using  nill which will return the latest block
// 	header, err := client.HeaderByNumber(context.Background(), nil)

// 	if err != nil {
// 		log.Fatal("unable to get block header")
// 	}

// 	fmt.Println(header.Number.String())

// 	//query the specified block
// 	blockNumber := big.NewInt(5231314)
// 	block, err := client.BlockByNumber(context.Background(), blockNumber)

// 	if err != nil {
// 		log.Fatal("Unable to get block by number,", blockNumber, err)
// 	}

// 	//will return the number of the block  and will convert it into Uint64
// 	fmt.Println(block.Number().Uint64())
// 	fmt.Println(block.Time())
// 	fmt.Println(block.Difficulty())
// 	fmt.Println(block.Hash().Hex())
// 	//because Transactions returns an array,
// 	fmt.Println(len(block.Transactions()))

// 	//return the chain ID
// 	// chainID, err := client.NetworkID(context.Background())
// 	if err != nil {
// 		log.Fatal("unable to get chain id")
// 	}

// 	//for every transaction withing the block returned
// 	// for _, tx := range block.Transactions() {

// 	// 	//to be able to see the Sender of each transaction
// 	// 	//will return the transaction in a message form
// 	// 	msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil)

// 	// 	if err != nil {
// 	// 		log.Fatal("Unable to convert tx(transaction) to msg")
// 	// 	}

// 	// 	fmt.Println(msg.From().Hex())

// 	// 	// fmt.Println(tx.Hash().Hex())
// 	// 	// fmt.Println(tx.Value().String())
// 	// 	// fmt.Println(tx.Gas())
// 	// 	// fmt.Println(tx.Nonce())
// 	// 	// fmt.Println(tx.Data())
// 	// 	// fmt.Println(tx.To().Hex())

// 	// }

// 	//getting transaction receipts
// 	for _, tx := range block.Transactions() {

// 		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())

// 		if err != nil {
// 			log.Fatal("Unable to get transaction receipt")
// 		}

// 		//status returns 1 if mined successfully
// 		fmt.Println(receipt.Status)

// 	}

// }
package main

//getting the transaction by using the hash
import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2ea67de6c0fe4745b93a7bf99ba89c86")

	if err != nil {
		log.Fatal("Unable to reach client")
	}

	//transform the Hex into Hash
	txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")

	//TransactionByHash returns three values. One is the transaction and the second is a boolean pending
	txn, isPending, err := client.TransactionByHash(context.Background(), txHash)

	if err != nil {
		log.Fatal("Unable to get txn by hash", txHash, err)
	}

	if isPending {
		fmt.Println("Txn is pending: ", txn.Hash())

	}

	fmt.Println(txn)

}
