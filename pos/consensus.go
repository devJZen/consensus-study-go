package pos

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"strconv"
	"time"
)

// Block은 PoS 블록체인의 블록을 나타냅니다.
type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	Data          []byte
	Hash          []byte
	Validator     string
}

// Blockchain은 PoS 블록체인을 나타냅니다.
type Blockchain struct {
	Blocks     []*Block
	Validators []Validator
}

// NewBlock은 새로운 블록을 생성합니다.
func NewBlock(data string, prevBlockHash []byte, validator string) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Validator:     validator,
	}
	block.SetHash()
	return block
}

// SetHash는 블록의 해시를 계산하고 설정합니다.
func (b *Block) SetHash() {
	// 1. Timestamp를 바이트 배열로 변환합니다.
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))

	// 2. PrevBlockHash, Data, Validator, Timestamp를 모두 합칩니다.
	headers := bytes.Join(
		[][]byte{
			b.PrevBlockHash,
			b.Data,
			[]byte(b.Validator),
			timestamp,
		},
		[]byte{},
	)
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// NewBlockchain은 제네시스 블록과 함께 새로운 블록체인을 생성합니다.
func NewBlockchain(validators []Validator) *Blockchain {
	genesisBlock := NewBlock("Genesis Block", []byte{}, validators[0].Address)
	return &Blockchain{
		Blocks:     []*Block{genesisBlock},
		Validators: validators,
	}
}

// ChooseValidator는 스테이크를 기반으로 검증인을 선택합니다.
func (bc *Blockchain) ChooseValidator() Validator {
	var totalStake int
	for _, v := range bc.Validators {
		totalStake += v.Stake
	}

	rand.Seed(time.Now().UnixNano())
	pick := rand.Intn(totalStake)
	var currentStake int
	for _, v := range bc.Validators {
		currentStake += v.Stake
		if pick < currentStake {
			return v
		}
	}
	return bc.Validators[0] // Fallback
}

// AddBlock은 블록체인에 새로운 블록을 추가합니다.
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	validator := bc.ChooseValidator()
	newBlock := NewBlock(data, prevBlock.Hash, validator.Address)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// IsValid는 블록체인의 유효성을 검사합니다.
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]

		if string(currentBlock.PrevBlockHash) != string(prevBlock.Hash) {
			return false
		}
	}
	return true
}
