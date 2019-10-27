package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//Block kepps block headers 区块定义
type Block struct {
	Timestamp int64       //区块创建时间戳
	Data []byte           //区块包含的数据
	PrevBlockHash []byte  //前一个区块的哈希值
	Hash []byte           //区块自身的哈希值，用于校验区块数据有效
}

//NewBlock creates and return Block 创建区块
func NewBlock(data string,prevBlockHash []byte) *Block  {
	block := &Block{time.Now().Unix(),[]byte(data),prevBlockHash,[]byte{}}
	block.SetHash()
	return block
}

//SetHash calculates and sets block hash 设置hash值
func (b *Block) SetHash()  {
	timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	headers := bytes.Join([][]byte{b.PrevBlockHash,b.Data,timestamp},[]byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// NewGenesisBlock creates and return s genesis Block 创世纪Block
func NewGenesisBlock() *Block  {
	return NewBlock("Genesis Block",[]byte{})
}
