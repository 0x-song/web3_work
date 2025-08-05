package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	address := common.HexToAddress("0x4Dbcb40022f18D9A726a814d9742F752B77D806a")
	fmt.Println(address)
	fmt.Println(address.Hex())
	//fmt.Println(address.Hash().Hex())
	fmt.Println(address.Bytes())
}
