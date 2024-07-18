package main

import (
	"context"
	"encoding/hex"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		klog.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xE6736a7e389d6F277b2c3002059d505417ef4C0A")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		klog.Fatal(err)
	}

	klog.Info(hex.EncodeToString(bytecode)) // 60806...10029
}
