package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d"))
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d"))

	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/T287_rO5bROjpOQYxzN84XpiaoodquJN")
	if err != nil {
		log.Fatal(err)
	}
	//uni
	//address := common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")
	address := common.HexToAddress("0x4Dbcb40022f18D9A726a814d9742F752B77D806a")

	fmt.Println(address)
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract)

}
