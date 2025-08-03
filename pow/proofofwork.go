package pow

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math"
	"math/big"
)

const targetBits = 16 // 난이도를 낮춰서 실행 시간을 줄입니다.

// ProofOfWork는 작업 증명을 나타냅니다.
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork는 새로운 ProofOfWork를 생성하고 반환합니다.
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// Run은 작업 증명을 수행합니다.
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		pow.block.Nonce = nonce
		hash = sha256.Sum256(pow.prepareData(nonce))
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

// Validate는 작업 증명의 유효성을 검사합니다.
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// IntToHex는 int64를 바이트 슬라이스로 변환합니다.
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
