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

func isDivis(n int, d int) bool {
	return d != 0 && n % d == 0
}

func getValue(values []int) int {
	for i, val := range values {
		for j := i+1; j < len(values); j++ {
			compVal := values[j]
			if isDivis(val, compVal) {return val / compVal}
			if isDivis(compVal, val) {return compVal / val}

		}
	}
	return 0
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		rawvalues := strings.Split(scanner.Text(), "\t")
		values := []int{}
		for _, val := range rawvalues {
			cell, err := strconv.Atoi(val)
			checkErr(err)
			values = append(values, cell)
		}
		fmt.Println(values)
		sum += getValue(values)
		
		fmt.Println("Sum: ", sum)
		
	}
	checkErr(scanner.Err())
}
