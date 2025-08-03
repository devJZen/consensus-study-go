package pow

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block은 블록체인의 기본 단위입니다.
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// SetHash는 블록의 해시를 계산하고 설정합니다.
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// Blockchain은 블록들의 포인터 배열입니다.
type Blockchain struct {
	Blocks []*Block
}

// AddBlock은 블록체인에 새로운 블록을 추가합니다.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlock은 주어진 데이터를 사용하여 새로운 블록을 생성합니다.
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewBlockchain은 제네시스 블록과 함께 새로운 블록체인을 생성합니다.
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// NewGenesisBlock은 제네시스 블록을 생성합니다.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
