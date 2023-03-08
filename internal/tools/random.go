package tools

import (
	"math/rand"
	"time"
)

func RandomSelector(r []string) int {
	
	min := 0
	max := len(r)

	rand.Seed(time.Now().UnixNano())

	selector := rand.Intn(max - min)
	
	return selector

}
