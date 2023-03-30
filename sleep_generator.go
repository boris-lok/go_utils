package utils

import (
	"math/rand"
	"time"
)

type SleepGenerator struct {
	maximum    int64
	multiplier float64
	delay      int64
	extra      int64
}

func (s *SleepGenerator) Next() bool {
	return true
}

func (s *SleepGenerator) Item() time.Duration {
	tick := int64(float64(rand.Int31()) * s.multiplier)
	var extra int64
	if 1<<31-1-tick > 0 {
		extra = tick + s.extra
	} else {
		extra = 1<<31 - 1
	}
	if s.maximum-s.extra > s.delay {
		s.extra = extra
		s.delay = s.delay + s.extra
		return time.Duration(s.delay)
	} else {
		return time.Duration(s.maximum)
	}
}

// CreateSleepGenerator create a sleep generator
func CreateSleepGenerator(delay time.Duration, maximum time.Duration, multiplier float64) *SleepGenerator {
	d := delay.Nanoseconds()
	m := maximum.Nanoseconds()
	return &SleepGenerator{
		delay:      d,
		maximum:    m,
		multiplier: multiplier,
		extra:      0,
	}
}
