package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"strconv"
	"strings"
)

/* Axes scale like this

Up = 3
Source = 1

distance 0 =

1

distance 1 =

1 + 3

distance 2 =

1 + 3 + 3 + 8

distance 3 = 

1 + 3 + 3 + 3 + 8 + 8 + 8

So 3s (the direction) scale linearly, 8s scale as a triangular series

We also know that the corners split the difference between the axes, except for when rotating from down to right

This corresponds with an increase of distance, so we can model corner calculation as: 

(dist2 - dist1 + number2 + number1) / 2

So if we loop until we find the range of numbers, then figure out where the corner is, and lump the corners into

belonging to the initial axis, we can calculate position on a cartesian plane

*/



const (
	right = 1
	up = 3
	left = 5
	down = 7
	source = 1
	halfScale = 4 // numbers increase by factor of 8, half it for
)


func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func calcNumber(orig int, dist int) int {
	return 1 + (orig * dist) + (halfScale * dist * (dist - source))
}

func calcCorner(numone int, distone int, numtwo int, disttwo int ) int {
	return (disttwo - distone + numtwo + numone) / 2
}

func main() {
	d, err := ioutil.ReadFile("input.txt")
	checkErr(err)
	number, err := strconv.Atoi(strings.Trim(string(d), "\n "))
	checkErr(err)
	// number := 22

	var foundDist, prevNum int

	directions := []int{right, up, left, down}
	fmt.Println("AYYYYYY")
	for dist := 1; foundDist == 0; dist++ {
		for _, direction := range directions {
			axNum := calcNumber(direction, dist)
			fmt.Println(dist, direction, axNum)
			if axNum > number {
				var prevDist int
				if direction == right {
					prevDist = dist - 1
				} else {
					prevDist = dist
				}

				corner := calcCorner(prevNum, prevDist, axNum, dist)

				if number < corner {
					fmt.Println("prevDist: ", prevDist, " number: ", number, " prevNum: ", prevNum)
					foundDist = prevDist + number - prevNum
				} else {
					fmt.Println("dist: ", dist, " number: ", number, " axNum: ", axNum)
					foundDist = dist + axNum - number
				}
				

				break
			}
			prevNum = axNum
		}
	}

	
	fmt.Println(foundDist)
}
