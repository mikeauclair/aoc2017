package main

import (
	"fmt"
)

const (
	steps  = 303
	rounds = 50000000
)

func main() {
	var cur, found int
	for i := 1; i <= rounds; i++ {
		cur = (cur + steps) % i + 1
		if cur == 1 {
			found = i
		}
	}
	fmt.Println(found)
}
