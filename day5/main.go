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

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	moves := 0
	var ind int
	fmt.Println(moves)
	values := []int{}
	for scanner.Scan() {
		number, err := strconv.Atoi(strings.Trim(scanner.Text(), "\n "))
		checkErr(err)
		values = append(values, number)
		
	}
	checkErr(scanner.Err())
	for ind < len(values) {
		oldVal := values[ind]
		if oldVal >= 3 {
			values[ind]--
		} else {
			values[ind]++
		}
		ind += oldVal
		moves++
	}
	fmt.Println(values)
	fmt.Println(moves)
}
