package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://localhost:8545/ws")
	if err != nil {
		klog.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x2B0C0B6E8c2246Bd90dD58BE6aF9e11D903e7DbF")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		klog.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			klog.Fatal(err)
		case vLog := <-logs:
			klog.Info(vLog) // pointer to event log
		}
	}
}
