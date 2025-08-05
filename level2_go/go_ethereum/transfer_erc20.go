package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
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
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("from:", fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(nonce)
	//接收方地址
	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	//代币合约地址 chainLink
	tokenAddress := common.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789")
	transferSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferSignature)
	methodId := hash.Sum(nil)[:4]
	fmt.Println("methodID::::", hexutil.Encode(methodId))
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println("paddedToAddress:", hexutil.Encode(paddedAddress))
	amount := new(big.Int)
	//10进制 转移数量 * 10 ^ decimals
	amount.SetString("1000000000000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println("padAmount:", hexutil.Encode(paddedAmount))
	var data []byte
	data = append(data, methodId...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gaslimit:", gasLimit)
	//将gasLimit设置为了预估的2倍才解决问题
	gasLimit = uint64(float64(gasLimit) * 2)
	fmt.Println("gaslimit:", gasLimit)
	//代币转账交易的本质便是调用ERC-20的transfer方法
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
