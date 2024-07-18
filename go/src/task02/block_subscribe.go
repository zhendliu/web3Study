package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		klog.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		klog.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			klog.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				klog.Fatal(err)
			}

			klog.Info(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			klog.Info(block.Number().Uint64())   // 3477413
			klog.Info(block.Time())              // 1529525947
			klog.Info(block.Nonce())             // 130524141876765836
			klog.Info(len(block.Transactions())) // 7
		}
	}
}
