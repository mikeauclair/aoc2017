package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
const (
	ascii_offset byte = 48
)

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func intSliceEq(a []int, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func checkSeen(seenSet [][]int, values []int) bool {
	for _, seen := range seenSet {
		if intSliceEq(seen, values) {
			fmt.Println(seenSet)
			fmt.Println(len(seenSet), seen, values)
			return true
		}
	}
	return false
}

func appendCopy(set [][]int, vals []int) [][]int {
	tmp := make([]int, len(vals))
	copy(tmp, vals)
	return append(set, tmp)
}

func findMaxInd(vals []int) int {
	var max, maxVal int
	for i, v := range vals {
		if v > maxVal {
			max = i
			maxVal = v
		}
	}
	return max
}

func redistribute(vals []int, index int) []int {
	amt := vals[index]
	vals[index] = 0
	newInd := index
	fmt.Println(amt)
	fmt.Println(newInd)
	for ; amt > 0; amt-- {
		if newInd < (len(vals) - 1) {
			newInd++
		} else {
			newInd = 0
		}
		vals[newInd] = vals[newInd] + 1
	}
	return vals
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	checkErr(scanner.Err())
	vals := strings.Split(strings.Trim(scanner.Text(), "\n "), "\t")
	values := make([]int, 0, len(vals))

	seen := [][]int{}

	var redis int
	var cycle bool
	for _, i := range vals {
		n, err := strconv.Atoi(i)
		checkErr(err)
		values = append(values, n)
	}
	seen = appendCopy(seen, values)

	fmt.Println(values)
	for redis=0; !cycle; redis++ {
		values = redistribute(values, findMaxInd(values))
		cycle = checkSeen(seen, values)
		if !cycle {
			seen = appendCopy(seen, values)
		}
	}
	fmt.Println(redis)

	
	
}
