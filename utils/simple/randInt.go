package simple

import (
	"math/rand"
	"time"
)

func RandInt(min, max int) int64 {
	rand.Seed(time.Now().UnixNano())
	if min >= max || min == 0 || max == 0 {
		return int64(max)
	}
	return int64(rand.Intn(max-min+1) + min)
}
