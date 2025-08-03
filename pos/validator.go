package pos

// Validator는 블록체인의 검증인을 나타냅니다.
type Validator struct {
	Address string // 검증인의 주소
	Stake   int    // 검증인이 스테이킹한 금액
}
