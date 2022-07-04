package utils

import "math/rand"

func Seed(seed int64) {
	rand.Seed(seed)
}
func Generate() int {
	return rand.Intn(10)
}
