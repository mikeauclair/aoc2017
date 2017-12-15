package main

import (
	"fmt"
)

const (
	lseed = 516
	rseed = 190
	lfactor = 16807
	rfactor = 48271
	divisor = 2147483647
	maxvals = 40000000
	bitmask = 65535
)

func genNext(val int, factor int) int {
	return (val * factor) % divisor
}

func main() {
	sum := 0
	lval := lseed
	rval := rseed
	for i := 0; i < maxvals; i++ {
		lval = genNext(lval, lfactor)
		rval = genNext(rval, rfactor)
		if lval & bitmask == rval & bitmask {
			sum++
		}
	}
	fmt.Println(sum)
	
}
