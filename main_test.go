package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"consensus-algorithms-go/poh"
	"consensus-algorithms-go/pos"
	"consensus-algorithms-go/pow"
)

func TestPoW(t *testing.T) {
	bc := pow.NewBlockchain()

	bc.AddBlock("블록 생성 테스트")
	bc.AddBlock("Send 1 BTC to 홍길동")
	bc.AddBlock("블록 추가시 내용을 테스트 합니다.")

	fmt.Println("PoW Blockchain:")
	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		proof := pow.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(proof.Validate()))
		if !proof.Validate() {
			t.Errorf("PoW is not valid")
		}
		fmt.Println()
	}
}

func TestPoS(t *testing.T) {
	validators := []pos.Validator{
		{Address: "validator1", Stake: 10},
		{Address: "validator2", Stake: 20},
		{Address: "validator3", Stake: 30},
	}
	bc := pos.NewBlockchain(validators)

	fmt.Println("\n--- PoS Test ---")
	fmt.Println("Participating Validators:")
	for _, v := range bc.Validators {
		fmt.Printf("  Address: %s, Stake: %d\n", v.Address, v.Stake)
	}
	fmt.Println()

	bc.AddBlock("Block 1")
	bc.AddBlock("Block 2")

	fmt.Println("PoS Blockchain:")
	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Validator: %s\n", block.Validator)
		fmt.Println()
	}

	if !bc.IsValid() {
		t.Errorf("PoS blockchain is not valid")
	}
}

func TestPoH(t *testing.T) {
	verifier := poh.NewVerifier()
	verifier.Record("Event 1")

	// 이벤트 사이에 의도적으로 시간 간격을 둡니다.
	time.Sleep(10 * time.Millisecond)

	verifier.Record("Event 2")

	fmt.Println("\n--- PoH Test ---")
	fmt.Println("PoH Verifier Log:")
	for _, entry := range verifier.Entries {
		fmt.Printf("Data: %s\n", entry.Data)
		fmt.Printf("Timestamp: %d\n", entry.Timestamp)
		fmt.Printf("Hash: %x\n", entry.Hash)
		fmt.Println()
	}

	if !verifier.Verify() {
		t.Errorf("PoH log is not valid")
	}
}
