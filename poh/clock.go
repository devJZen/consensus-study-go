package poh

import "time"

// Clock은 시스템의 시간을 나타냅니다.
type Clock struct {
	Time int64
}

// NewClock은 새로운 시계를 생성합니다.
func NewClock() *Clock {
	//다른 프로토콜과 달리 나노초 단위를 계산합니다.
	return &Clock{Time: time.Now().UnixNano()}
}

// Tick은 시계의 시간을 증가시킵니다.
func (c *Clock) Tick() {
	c.Time++
}
