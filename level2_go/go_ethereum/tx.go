package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/T287_rO5bROjpOQYxzN84XpiaoodquJN")
	if err != nil {
		log.Fatal(err)
	}
	blockNum := big.NewInt(22892268)
	block, err := client.BlockByNumber(context.Background(), blockNum)
	if err != nil {
		log.Fatal(err)
	}
	for _, tx := range block.Transactions() {
		if tx.Hash().Hex() == "0xcdb169cbede9154ba3e9bbee48f1115f2f320722426c3c20a163776a41ffae43" {
			fmt.Println("hash:", tx.Hash().Hex())
			fmt.Println("value:", tx.Value().String())
			fmt.Println("gas:", tx.Gas())
			fmt.Println("gasPrice:", tx.GasPrice().Uint64())
			fmt.Println("nonce:", tx.Nonce())
			fmt.Println("data:", tx.Data())
			fmt.Println("len:", len(tx.Data()))
			fmt.Println("to:", tx.To().Hex())
			fmt.Println("===========================")
			chainId, err := client.NetworkID(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("chainId:", chainId)
			//如果是域名，查询不到from信息
			from, err := types.Sender(types.NewEIP155Signer(chainId), tx)
			if err == nil {
				fmt.Println("from:", from.Hex())
			} else {
				log.Fatal(err)
			}

			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("status:", receipt.Status)
			fmt.Println("Log:", receipt.Logs)
		}
	}
	fmt.Println("=============tx by block hash==============")
	blockHash := block.Hash()
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count:", count)
	// for i := uint(0); i < count; i++ {
	// 	tx, err := client.TransactionInBlock(context.Background(), blockHash, i)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("tx:", tx.Hash().Hex())
	// }
	tx, err := client.TransactionInBlock(context.Background(), blockHash, 0)

	tx, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
	fmt.Println("tx:", tx.Hash().Hex())
	fmt.Println("isPending:", isPending)
}
