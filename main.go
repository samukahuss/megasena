package main

import (
	"fmt"
	"math/rand"
)

func main() {
	bet := make(map[int]int)

	for len(bet) < 6 {
		number := rand.Intn(60) + 1
		bet[number] = number
	}

	for k := range bet {
		fmt.Print(k, " ")
	}
}
