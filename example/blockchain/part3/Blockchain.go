package main

import (
	"fmt"
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// tip 这个词本身有事物尖端或尾部的意思，这里指的是存储最后一个块的哈希在链的末端可能出现短暂分叉的情况，
// 所以选择 tip 其实就是选择了哪条链 db 存储数据库连接
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

// 创建区块链
func NewBlockchain() *Blockchain {
	var tip []byte
	db, e := bolt.Open(dbFile, 0600, nil)
	if e != nil {
		panic(e)
	}
	e = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new on ...")
			genesisBlock := GenerateGenesisBlock()

			b, e := tx.CreateBucket([]byte(blocksBucket))
			if e != nil {
				panic(e)
			}
			e = b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if e != nil {
				panic(e)
			}
			e = b.Put([]byte("l"), genesisBlock.Hash)
			if e != nil {
				panic(e)
			}
			tip = genesisBlock.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if e != nil {
		panic(e)
	}

	return &Blockchain{tip, db}
}

// 加入区块时，需要将区块持久化到数据库中
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	// 首先获取最后一个块的哈希用于生成新快的哈希
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		panic(err)
	}

	newBlock := NewBlock(lastHash, data)
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put([]byte(newBlock.Hash), newBlock.Serialize())
		if err != nil {
			panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			panic(err)
		}
		bc.tip = newBlock.Hash
		return nil
	})
	if err != nil {
		panic(err)
	}
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.tip, bc.db}
}

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
