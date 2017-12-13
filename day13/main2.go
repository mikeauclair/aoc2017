package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const ()

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func cyclePos(movePos int, pivot int) int {
	if movePos > pivot {
		return movePos - 2 * (movePos - pivot)
	} else {
		return movePos
	}
}

func collision(pos int, height int) bool {
	cycleSize := 2 * (height - 1)
	return 0 == cyclePos(pos % cycleSize, height - 1)
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	cost := -1

	var lines [][]int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		pos, err := strconv.Atoi(line[0])
		checkErr(err)
		height, err := strconv.Atoi(line[1])
		checkErr(err)
		lines = append(lines, []int{pos, height})
		
	}
	checkErr(scanner.Err())
	delay := 0
	for ; cost != 0; delay++ {
		cost = 0
		for _,line := range lines {
			pos, height := line[0], line[1]
			if collision(pos + delay, height) {
				cost ++
				break
			}
		}
	}
	fmt.Println(cost)
	fmt.Println("Delay: ", delay-1)
	
}
