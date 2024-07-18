package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

	privateKey, err := crypto.HexToECDSA("879ec62a06c1aab868e67769c756a99ecf3c85562c4462276fb469afd7776a29")
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

	address := common.HexToAddress("0xE6736a7e389d6F277b2c3002059d505417ef4C0A")
	instance, err := store.NewStore(address, client)
	if err != nil {
		klog.Fatal(err)
	}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		klog.Fatal(err)
	}

	klog.Infof("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	result, err := instance.Items(nil, key)
	if err != nil {
		klog.Fatal(err)
	}

	klog.Info(string(result[:])) // "bar"
}
