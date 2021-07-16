package main

import (
	"fmt"
	"gochain/blockchain"
	"strconv"
)

func main() {

	chain := blockchain.InitBlockChain()

	chain.AddBlock("first block")
	chain.AddBlock("second block")
	chain.AddBlock("third block")

	for _, block := range chain.Blocks {
		fmt.Printf("prev Hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("Pow : %s", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}