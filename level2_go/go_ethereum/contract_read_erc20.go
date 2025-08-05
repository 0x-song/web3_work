package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	"geth/token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/T287_rO5bROjpOQYxzN84XpiaoodquJN")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// uni代币的地址
	tokenAddress := common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(instance)
	//持有uni代币第三个地址的余额
	address := common.HexToAddress("0x47173B170C64d16393a52e6C480b3Ad8c302ba1e")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	//23,047,480.513198522
	fmt.Println(bal)

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)         // "name: Uniswap"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: UNI"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f\n", value) // "balance: 23047480.513199"
}