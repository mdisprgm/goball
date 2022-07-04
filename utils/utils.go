package utils

import (
	"bufio"
	"math/rand"
	"os"
)

func Seed(seed int64) {
	rand.Seed(seed)
}
func Generate() int {
	return rand.Intn(10)
}

func IsSafe(err error) {
	if err != nil {
		panic(err)
	}
}

var __stdin__ = bufio.NewReader(os.Stdin)

func Flush() {
	__stdin__.ReadLine()
}
