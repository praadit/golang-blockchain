package main

import (
	"blockchain/blockchain"
	"fmt"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Pertama")
	chain.AddBlock("Kedua Kali")
	chain.AddBlock("Ketiga")

	for _, block := range chain.Blocks {
		fmt.Printf("Prev Hash : %x\n", block.PrevHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Println()
	}
}
