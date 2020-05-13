package main

type Blockchain struct {
	Blocks []*Block
}

// 创建区块链
func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{[]*Block{genesisBlock}}
	return &blockchain
}

// 创建区块
func (bc *Blockchain) AddBlock(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(preBlock.Hash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}
