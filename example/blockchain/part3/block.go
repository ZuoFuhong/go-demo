package main

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	Timestamp     int64  // 区块时间戳
	PrevBlockHash []byte // 上一个区块的哈希值
	Hash          []byte // 当前区块的哈希值
	Data          []byte // 区块数据
	Nonce         int
}

// 将一个 Block 序列化为一个字节数组
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

// 将字节数组反序列化为一个 Block
func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))

	err := decoder.Decode(&block)
	if err != nil {
		panic(err)
	}
	return &block
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
