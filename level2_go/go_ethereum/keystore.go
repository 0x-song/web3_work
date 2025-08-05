package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
	// 第二个参数 keystore.StandardScryptN 是 Scrypt 密钥派生函数的 N 参数，用于指定内存和 CPU 消耗
	// 第三个参数 keystore.StandardScryptP 是 Scrypt 密钥派生函数的 P 参数，用于指定并行化因子
	// 这两个参数的作用是控制密钥派生函数的强度和安全性。
	// 具体来说，N 参数控制内存消耗，P 参数控制 CPU 消耗。
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	fmt.Println(ks)
	password := "123456"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}

func importKs() {
	file := "./wallets/UTC--2025-07-04T23-49-13.455511000Z--61e1e6b45de4e8dd4753c307fa9bdc24de589dbd"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	password := "123456"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}

func main() {
	importKs()
}
