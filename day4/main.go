package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		set := make(map[string]bool)
		hasDupe := false
		values := strings.Split(scanner.Text(), " ")
		
		for _, val := range values {
			fmt.Println(val)
			_, ok := set[val]
			if ok {
				hasDupe = true
				break
			} else {
				set[val] = true
			}
		}
		if !hasDupe {
			sum += 1
		}
		fmt.Println(sum)
		
	}
	checkErr(scanner.Err())
}
