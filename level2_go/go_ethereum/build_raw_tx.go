package main

import (
	"bufio"
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
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
	amount := 0.02
	value := big.NewInt(int64(amount) * 1e18)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := "0xF8c3f049908D3E924845AB8b1CAEb10C96CE57fb"
	toAddr := common.HexToAddress(toAddress)
	tx := types.NewTransaction(nonce, toAddr, value, gasLimit, gasPrice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	//transaction切片
	ts := types.Transactions{signedTx}
	var buf bytes.Buffer
	//将切片中指定索引位置的交易编码到一个Write实现对象中
	ts.EncodeIndex(0, &buf)
	rawHex := hex.EncodeToString(buf.Bytes())
	fmt.Println(rawHex)
	rawTxBytes, err := hex.DecodeString(rawHex)
	if err != nil {
		log.Fatal(err)
	}
	tx2 := new(types.Transaction)
	// 使用 rlp.DecodeBytes 函数将 RLP 编码的字节数据 rawTxBytes 解码为 types.Transaction 类型的对象 tx2。
	// rlp.DecodeBytes 是以太坊 RLP 编码库提供的函数，第一个参数是待解码的字节切片，第二个参数是用于存储解码结果的目标对象指针。
	// 这里将 rawTxBytes 解码后存储到 tx2 中，tx2 是事先创建的 types.Transaction 指针。
	rlp.DecodeBytes(rawTxBytes, &tx2)
	// 检查解码过程中是否发生错误。如果 err 不为 nil，说明解码失败。
	// 在 Go 语言里，函数调用可能会返回错误信息，通常通过检查错误变量是否为 nil 来判断操作是否成功。
	// 若解码失败，使用 log.Fatal 函数输出错误信息并终止程序。
	err2 := client.SendTransaction(context.Background(), tx2)
	fmt.Println("发送交易完成")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("tx sent: %s\n", tx2.Hash().Hex())
}
