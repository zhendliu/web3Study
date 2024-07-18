package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
)

func main() {

	//createKs()
	importKs()
}

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func importKs() {
	file := "./tmp/UTC--2024-07-15T14-11-14.686598000Z--f3122dbaf0e46b46d49f3386790021b10182c070"
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("异常了", err)
		log.Fatal(err)
	}

	keystore.

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		fmt.Println("Import 异常了", err)
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	/*if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}*/
}
