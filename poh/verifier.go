package poh

import (
	"crypto/sha256"
	"strconv"
	"time"
)

// Entry는 PoH 로그의 항목을 나타냅니다.
type Entry struct {
	Data      []byte
	Timestamp int64
	Hash      []byte
}

// Verifier는 PoH 로그를 검증합니다.
type Verifier struct {
	Entries []*Entry
}

// NewVerifier는 새로운 검증기를 생성합니다.
func NewVerifier() *Verifier {
	genesisEntry := &Entry{
		Data:      []byte("Genesis Entry"),
		Timestamp: time.Now().UnixNano(),
		Hash:      []byte{},
	}
	genesisEntry.Hash = CalculateHash(genesisEntry)
	return &Verifier{Entries: []*Entry{genesisEntry}}
}

// Record는 새로운 항목을 로그에 기록합니다.
func (v *Verifier) Record(data string) {
	prevEntry := v.Entries[len(v.Entries)-1]
	newEntry := &Entry{
		Data:      []byte(data),
		Timestamp: time.Now().UnixNano(),
	}
	newEntry.Hash = CalculateHash(newEntry, prevEntry.Hash)
	v.Entries = append(v.Entries, newEntry)
}

// Verify는 로그의 모든 항목을 검증합니다.
func (v *Verifier) Verify() bool {
	for i := 1; i < len(v.Entries); i++ {
		currentEntry := v.Entries[i]
		prevEntry := v.Entries[i-1]
		expectedHash := CalculateHash(currentEntry, prevEntry.Hash)
		if string(currentEntry.Hash) != string(expectedHash) {
			return false
		}
	}
	return true
}

// CalculateHash는 항목의 해시를 계산합니다.
func CalculateHash(entry *Entry, prevHash ...[]byte) []byte {
	headers := append(entry.Data, []byte(strconv.FormatInt(entry.Timestamp, 10))...)
	if len(prevHash) > 0 {
		headers = append(headers, prevHash[0]...)
	}
	hash := sha256.Sum256(headers)
	return hash[:]
}
