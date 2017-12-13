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
	var cost int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		pos, err := strconv.Atoi(line[0])
		checkErr(err)
		height, err := strconv.Atoi(line[1])
		checkErr(err)
		if collision(pos, height) {
			cost += pos * height
		}
	}
	checkErr(scanner.Err())
	fmt.Println(cost)
}
