package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)


type Block struct {
	Hash     []byte
	//Data     []byte
	Transactions []*Transaction
	PrevHash []byte
	Nonce 	 int
}
/*
type BlockChain struct {
	Blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})

	hash := sha256.Sum256(info)

	b.Hash = hash[:]
}
*/
func CreateBlock(tsx []*Transaction , prevHash []byte) *Block {
	block := &Block{[]byte{}, tsx, prevHash, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
/*
func (chain *BlockChain) AddBlock(data string) {
	preBlock := chain.Blocks[len(chain.Blocks)-1]
	new := createBlock(data, preBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}
*/
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}
/*
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
*/
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (b *Block) Serialize() []byte {

	var res bytes.Buffer

	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

func Deserialize(data []byte) *Block {

	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func (b *Block) HashTransactions() []byte {

	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	
	return txHash[:]
}