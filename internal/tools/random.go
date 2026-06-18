package tools

import (
	"math/rand"
)

func RandomSelector(r []string) int {

	max := len(r)

	selector := rand.Intn(max)

	return selector

}
