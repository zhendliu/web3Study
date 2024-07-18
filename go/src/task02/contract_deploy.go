package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	store "task02/contracts"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		klog.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("3af7a7a7405e279548c7169db5094a2c70c569467b2783d837c8b4414eaeda57")
	if err != nil {
		klog.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		klog.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		klog.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		klog.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		klog.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance
}

/*
0xE6736a7e389d6F277b2c3002059d505417ef4C0A
0x53575f1b0c427f90812e4c6152d78482855f0f72a63df285571ed094fac899fc

*/
