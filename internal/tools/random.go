package tools

import (
	"math/rand"
	"time"
)

func RandomSelector(r []string) int {

	max := len(r)

	rand.Seed(time.Now().UnixNano())

	selector := rand.Intn(max)

	return selector

}
