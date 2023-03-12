package helper

import (
	"math/rand"
	"time"
)

func GenerateRandValue() int {
	const (
		start = 1
		end   = 100
	)

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end-start) + start
}
