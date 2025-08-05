package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/T287_rO5bROjpOQYxzN84XpiaoodquJN")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	account := common.HexToAddress("0x4Dbcb40022f18D9A726a814d9742F752B77D806a")
	balance, err := client.BalanceAt(context.Background(), account, nil)

	blockNum := big.NewInt(8364046)

	balanceAt, err := client.BalanceAt(context.Background(), account, blockNum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
	fmt.Println(balanceAt)

	fbalance := new(big.Float)
	fbalance.SetInt(balance)
	//fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(1e18))
	fmt.Println(ethValue)
}
