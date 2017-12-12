package main

import (
	"fmt"
	"os"
	"log"
)

const (
	number = 368078
)

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

/*

Given

147  142  133  122   59
304    5    4    2   57
330   10    1    1   54
351   11   23   25   26
362  747  806--->   ...

The trailer for 133 is 2, the leader is 5, and the inbound is 4
We special case the trailer for the cells after corners as n-2

*/

func main() {

	var number int
	_, err := fmt.Sscanf(os.Args[1], "%d", &number)
	if err != nil {
		panic(err)
	}
	
	var lastMultiple, trailer, inbound, leader int

	lastMultiplier := 1
	lastMultiple = 0
	seq := []int{1}
	corners := []int{}

	updateCorners := func(multiplier int) {
		per := 8 * multiplier / 4
		corners = []int{
			lastMultiple + per*1,
			lastMultiple + per*2,
			lastMultiple + per*3,
			// turn the corner
			lastMultiple + per*4 + 1,
		}
		lastMultiple = lastMultiple + 8*multiplier

	}

	contains := func(corners *[]int, value int) bool {
		for _, c := range *corners {
			if value == c {
				return true
			}
		}
		return false
	}
	updateCorners(lastMultiplier)
	leader = 1
	cur := len(seq)
	for foundNum := 0; foundNum < number; cur++ {
		var val int
		if cur != 1 {
			val = seq[cur-1]
		}

		val += seq[inbound]
		if contains(&corners, cur) {
		} else if contains(&corners, cur-1) {
			val += seq[cur-2]
			if !contains(&corners, cur+1) {
				val += seq[leader]
				leader++
				inbound++
			}

			// After rounding the corner, check to see if you are in a new layer
			// and update the corner set if you are

			if cur > lastMultiple {
				lastMultiplier = lastMultiplier + 1
				updateCorners(lastMultiplier)
			}

		} else if contains(&corners, cur+1) {
			if cur != 1 {
				val += seq[trailer]
				trailer++
			}
		} else {
			val = val + seq[trailer] + seq[leader]
			inbound++
			trailer++
			leader++
		}
		foundNum = val
		seq = append(seq, val)
	}

	fmt.Println(seq)
}
