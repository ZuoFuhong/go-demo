package main

import (
	"fmt"
)

func main() {
	blockchain := NewBlockchain()
	blockchain.AddBlock("Send 1 BTC to dazuo")
	blockchain.AddBlock("Send 1 EOS to dazuo")
	blockchain.Print()
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Prev.Hash：%s\n", block.PrevBlockHash)
		fmt.Printf("Curr.Hash：%s\n", block.Hash)
		fmt.Printf("Data：%s\n", block.Data)
		fmt.Printf("Timestamp：%d\n", block.Timestamp)
	}
}
