package main

import (
	"fmt"
)

func main() {
	blockchain := NewBlockchain()
	blockchain.AddBlock("Send 1 ETC to Lucy")
	blockchain.Print()
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index：%d\n", block.Index)
		fmt.Printf("Prev.Hash：%x\n", block.PrevBlockHash)
		fmt.Printf("Curr.Hash：%x\n", block.Hash)
		fmt.Printf("Data：%s\n", block.Data)
		fmt.Printf("Timestamp：%d\n", block.Timestamp)
	}
}
