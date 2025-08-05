package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(privateKey)
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyFinal := hexutil.Encode(privateKeyBytes)[2:]
	fmt.Println("私钥:", privateKeyFinal)
	publicKey := privateKey.Public()
	//类型断言 推断publicKey是 *ecdsa.PublicKey类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// 去掉前面的0x  04 代表压缩公钥
	publicKeyFinal := hexutil.Encode(publicKeyBytes)[4:]
	fmt.Println("公钥:", publicKeyFinal)
	//publicKeyECDSA是从上述推断出来的指针数据类型，使用*解引用后，得到的是ecdsa.PublicKey类型
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("地址:", address)
}
