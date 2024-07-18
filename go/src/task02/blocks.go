package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	klog.Info(header.Number.String()) // 5671744

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	klog.Info(block.Number().Uint64())
	klog.Info(block.Time())
	klog.Info(block.Difficulty().Uint64())
	klog.Info(block.Hash().Hex())
	klog.Info(len(block.Transactions()))

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	klog.Info(count) // 144
}
