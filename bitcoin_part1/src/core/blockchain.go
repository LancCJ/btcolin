package core

// Blockchain keeps s sequence of Blocks 区块链
type Blockchain struct {
	Blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain 添加区块
func (bc *Blockchain) AddBlock(data string)()  {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data,prevBlock.Hash)
	bc.Blocks = append(bc.Blocks,newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block //使用创世区块创建区块链
func NewBlockchain() *Blockchain  {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}