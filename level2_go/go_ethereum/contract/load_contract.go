package main

import (
	"fmt"
	"geth/contract/store"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/T287_rO5bROjpOQYxzN84XpiaoodquJN")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	address := common.HexToAddress("0xf70C0657B1133Ec184583C1A9f168144AFbf5bD2")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("instance is ", instance)
	//当一个结构体嵌入另一个结构体时，被嵌入结构体的方法会被提升到嵌入结构体上，
	//嵌入结构体的实例就能直接调用这些方法，这看起来类似于继承。
	//instance.StoreCaller.Version(nil)
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("version is ", version)

	key := [32]byte{}
	copy(key[:], []byte("foo"))

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result is ", string(result[:]))

}
