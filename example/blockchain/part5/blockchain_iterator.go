package main

import "github.com/boltdb/bolt"

// BlockchainIterator is used to iterate over blockchain blocks
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Iterator returns a BlockchainIterat
func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.tip, bc.db}
}

// Next returns next block starting from the tip
func (bi *BlockchainIterator) Next() *Block {
	var block *Block

	err := bi.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodeBlock := b.Get(bi.currentHash)
		block = DeserializeBlock(encodeBlock)
		return nil
	})
	if err != nil {
		panic(err)
	}
	bi.currentHash = block.PrevBlockHash
	return block
}
