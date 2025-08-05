package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
)

// parseDerivationPath 手动解析派生路径
func parseDerivationPath(path string) ([]uint32, error) {
	segments := strings.Split(strings.TrimPrefix(path, "m/"), "/")
	var result []uint32
	for _, segment := range segments {
		hardened := strings.HasSuffix(segment, "'")
		if hardened {
			segment = strings.TrimSuffix(segment, "'")
		}
		num, err := strconv.ParseUint(segment, 10, 31)
		if err != nil {
			return nil, err
		}
		if hardened {
			num += 0x80000000
		}
		result = append(result, uint32(num))
	}
	return result, nil
}

var globalMnemonic string

func hdwallet1() {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		log.Fatal(err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	globalMnemonic = mnemonic
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("助记词:", mnemonic)
	//通过助记词生成种子
	seed := bip39.NewSeed(mnemonic, "")
	fmt.Println("种子:", seed)
	//生成主密钥
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("主密钥:", masterKey)
	//生成子密钥
	derivationPath := "m/44'/60'/0'/0/0"
	path, err := parseDerivationPath(derivationPath)

	currentKey := masterKey
	for _, segment := range path {
		currentKey, err = currentKey.NewChildKey(segment)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("子密钥:", currentKey)
		privateKey, err := crypto.ToECDSA(currentKey.Key)
		if err != nil {
			log.Fatal(err)
		}
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
}

func hdwallet2() {
	// entropy, err := bip39.NewEntropy(128)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// mnemonic, err := bip39.NewMnemonic(entropy)
	mnemonic := globalMnemonic
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("助记词:", mnemonic)
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())

	path2 := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account2, err := wallet.Derive(path2, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account2.Address.Hex())
}

func main() {
	hdwallet1()
	fmt.Println("------------------------")
	hdwallet2()
}
