package main

import (
	"context"
	"fmt"
	"geth/contract/store"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/T287_rO5bROjpOQYxzN84XpiaoodquJN")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	contractAddress := common.HexToAddress("0xf70C0657B1133Ec184583C1A9f168144AFbf5bD2")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(8882300),
		ToBlock:   big.NewInt(8882300),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))

	for _, vLog := range logs {
		//log.Println(vLog)
		//定义一个匿名结构体，并实例化对象   {}最后一个大括号表示实例化对象
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		//用于将智能合约的调用数据或者事件数据解析到指定的结构体中。
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(event.Key[:]))
		fmt.Println(string(event.Value[:]))
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())

		var topics [4]string
		for i, topic := range vLog.Topics {
			topics[i] = topic.Hex()
		}
		fmt.Println(topics[0])
	}

}
