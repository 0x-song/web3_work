package main

import (
	"context"
	"fmt"
	"geth/erc20/token"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/T287_rO5bROjpOQYxzN84XpiaoodquJN")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	//chainLink link token
	contractAddress := common.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}
	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logTransferApprovalHash := crypto.Keccak256Hash(LogApprovalSig)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			//打印当前交易的hash
			fmt.Printf("Transaction Hash: %s\n", vLog.TxHash.Hex())

			//log.Println(vLog)
			switch vLog.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				fmt.Println("LogTransfe:")
				var transferEvent LogTransfer
				//第一个参数用于存储解析过后的日志数据
				//第二个参数是日志事件的名称
				//第三个参数是日志事件的数据
				//这行代码的作用是依据 contractAbi 中定义的 Transfer 事件的 ABI 信息，
				// 将 vLog.Data 里的原始字节数据解析出来，并填充到 transferEvent 结构体中。
				err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

				fmt.Printf("From: %s\n", transferEvent.From.Hex())
				fmt.Printf("To: %s\n", transferEvent.To.Hex())
				fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())

			case logTransferApprovalHash.Hex():
				fmt.Println("LogApproval:")
				var approvalEvent LogApproval
				err := contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
				approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())

				fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
				fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
				fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())

			}
			fmt.Println("------------------------")

		}
	}
}
