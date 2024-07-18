package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ethereum/go-ethereum/whisper/shhclient"
)

func main() {
	client, err := shhclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		klog.Fatal(err)
	}

	_ = client // we'll be using this in the 下个章节
	fmt.Println("we have a whisper connection")
}
