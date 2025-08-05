package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"geth/contract/store"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/T287_rO5bROjpOQYxzN84XpiaoodquJN")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	file, err := os.Open("/Users/0xsong/private.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	line = line[:len(line)-1]

	privateKey, err := crypto.HexToECDSA(line)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("加载私钥完成")
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xf70C0657B1133Ec184583C1A9f168144AFbf5bD2")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	key := [32]byte{}
	value := [32]byte{}
	//copy需要接受切片作为参数，而不是数组。
	//key[:]将数组转换为切片，切片是对数组的引用，所以可以直接使用。
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Println("验证是否写入成功")
	//写入之后立刻读取，读取不到数据
	//需要等待一个区块的时间，才能读取到数据
	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}
	//result [32]byte
	fmt.Println("result is ", string(result[:]))
}
