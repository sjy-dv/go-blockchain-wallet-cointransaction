package blockchain

import (
	"bytes"
	"crypto/sha256"
)


type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce 	 int
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})

	hash := sha256.Sum256(info)

	b.Hash = hash[:]
}

func createBlock(data string , prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	preBlock := chain.Blocks[len(chain.Blocks)-1]
	new := createBlock(data, preBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block {
	return createBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
