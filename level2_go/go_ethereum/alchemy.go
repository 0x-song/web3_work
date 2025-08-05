package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
//	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/T287_rO5bROjpOQYxzN84XpiaoodquJN")
	client, err := ethclient.Dial("http://localhost:7545")

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	fmt.Println(client)
	fmt.Println("connect to Alchemy Node successfully")
}
