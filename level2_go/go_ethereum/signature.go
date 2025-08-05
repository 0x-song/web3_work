package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func signature_generate() {
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
	data := []byte("hello,world")
	hash := crypto.Keccak256Hash(data)
	//0xab036729af8b8f9b610af4e11b14fa30c348f40c2c230cce92ef6ef37726fee7
	fmt.Println(hash.Hex())
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	//0x75b3813ed4440f162c69b68934b10d727ef82343d19fa9f600c2dadc090a6d8f0464a33f21c602724989871b712311ca0235c50847f7196fb8ead2a17708df6c01
	fmt.Println(hexutil.Encode(signature))

}

func signature_verify() {
	//我们需要有 3 件事来验证签名：签名，原始数据的哈希以及签名者的公钥。
	data := []byte("hello,world")
	hash := crypto.Keccak256Hash(data)
	signature, err := hexutil.Decode("0x75b3813ed4440f162c69b68934b10d727ef82343d19fa9f600c2dadc090a6d8f0464a33f21c602724989871b712311ca0235c50847f7196fb8ead2a17708df6c01")
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	recoverPubKey, err := crypto.UnmarshalPubkey(sigPublicKey)
	recoverAddress := crypto.PubkeyToAddress(*recoverPubKey)
	fmt.Println(recoverAddress.Hex())
}

func main() {
	signature_generate()
	signature_verify()
}
