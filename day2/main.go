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
	sum := 0
	for scanner.Scan() {
		var min, max int
		values := strings.Split(scanner.Text(), "\t")
		for _, val := range values {
			i, err := strconv.Atoi(val)
			checkErr(err)
			if min == 0 || i < min {
				min = i
			}
			if max == 0 || i > max {
				max = i
			}
		}
		fmt.Println("Max: ", max)
		fmt.Println("Min: ", min)
		sum += max
		sum -= min
		fmt.Println("Sum: ", sum)
		
	}
	checkErr(scanner.Err())
}
