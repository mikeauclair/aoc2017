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
	maxvals = 5000000
	bitmask = 65535
	lmultiple = 4
	rmultiple = 8
)

func genNext(val int, factor int, multiple int) int {
	retVal := (val * factor) % divisor
	for (retVal % multiple) != 0 {
		retVal = (retVal * factor) % divisor
	}
	return retVal
}

func main() {
	sum := 0
	lval := lseed
	rval := rseed
	for i := 0; i < maxvals; i++ {
		lval = genNext(lval, lfactor, lmultiple)
		rval = genNext(rval, rfactor, rmultiple)
		if lval & bitmask == rval & bitmask {
			sum++
		}
	}
	fmt.Println(sum)
	
}
