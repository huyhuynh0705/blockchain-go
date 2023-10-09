package main

import (
	"blockchain-go/domain"
	"fmt"
	"strconv"
)

func main() {
	bc := domain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ngo Van Gioi")
	bc.AddBlock("Send 2 BTC to Tran Thanh Dong")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := domain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
