package main

import (
	"time"
)

type Block struct {
	Timestamp     int64  // 区块时间戳
	PrevBlockHash []byte // 上一个区块的哈希值
	Hash          []byte // 当前区块的哈希值
	Data          []byte // 区块数据
	Nonce         int
}

// 生成区块
func NewBlock(preBlockHash []byte, data string) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: preBlockHash,
		Hash:          []byte{},
		Data:          []byte(data),
		Nonce:         0,
	}
	pow := NewProofOfWork(block)
	block.Nonce, block.Hash = pow.Run()
	return block
}

// 生成创世块
func GenerateGenesisBlock() *Block {
	return NewBlock([]byte{}, "Genesis Block")
}
