package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

//higher level setting in real-service
//const Difficultly = 6

const Difference = 9

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {

	target := big.NewInt(1)
	//target.Lsh(target, uint(256-Difficultly))

	target.Lsh(target, uint(256-Difference))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte {
			pow.Block.PrevHash,
			pow.Block.HashTransactions(),
			ToHex(int64(nonce)),
			ToHex(int64(Difference)),
		},
		[]byte{},
	)
	return data
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)

	binary.Write(buff, binary.BigEndian, num)

	return buff.Bytes()
}

func (pow *ProofOfWork) Validate() bool {

	var bigIntHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	bigIntHash.SetBytes(hash[:])

	return bigIntHash.Cmp(pow.Target) == -1
}

func (pow *ProofOfWork) Run() (int, []byte) {

	var bigIntHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		bigIntHash.SetBytes(hash[:])

		if bigIntHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

/*
 ver 2021-07-16 -ver2
*/