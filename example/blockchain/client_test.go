package blockchain

import (
	"fmt"
	"testing"
)

func TestBlockChain(t *testing.T) {
	blockchain := NewBlockchain()
	blockchain.SendData("Send 1 BTC to dazuo")
	blockchain.SendData("Send 1 EOS to dazuo")
	blockchain.Print()
}

func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index：%d\n", block.Index)
		fmt.Printf("Prev.Hash：%s\n", block.PrevBlockHash)
		fmt.Printf("Curr.Hash：%s\n", block.Hash)
		fmt.Printf("Data：%s\n", block.Data)
		fmt.Printf("Timestamp：%d\n", block.Timestamp)
	}
}
