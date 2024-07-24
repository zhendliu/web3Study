package main

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	EncodeToBytes()
}

func Encode() error {

	val := "abc"

	b := new(bytes.Buffer)
	err := rlp.Encode(b, val)
	if err != nil {
		panic(err)
	}

	fmt.Println(b.String())

	return nil

}

func EncodeToBytes() error {
	val := "abc"
	b, err := rlp.EncodeToBytes(val)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)

	fmt.Println(string(b))

	return nil
}
