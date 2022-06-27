package datetime

import "time"

type Clock interface {
	Now() time.Time
}

type clock struct {
}

func (s *clock) Now() time.Time {
	return time.Now()
}

func NewClock() *clock {
	return &clock{}
}
