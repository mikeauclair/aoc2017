package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

type reducible struct {
	Opposite      int
	Clock         int
	ClockReduce   int
	Counter       int
	CounterReduce int
}

func getIndex(set []string, value string) int {
	for i, v := range set {
		if v == value {
			return i
		}
	}
	return -1
}

func getReducibles(directions []string, direction string) reducible {
	setLen := len(directions)
	dirInd := getIndex(directions, direction)
	opp := dirInd + setLen/2
	if opp >= setLen {
		opp -= setLen
	}
	clock := dirInd + 2
	if clock >= setLen {
		clock -= setLen
	}
	clockReduce := dirInd + 1
	if clockReduce >= setLen {
		clockReduce -= setLen
	}
	counter := dirInd - 2
	if counter < 0 {
		counter += setLen
	}
	counterReduce := dirInd - 1
	if counterReduce < 0 {
		counterReduce += setLen
	}
	return reducible{opp, clock, clockReduce, counter, counterReduce}
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	reader := bufio.NewReader(f)

	var readErr error

	directions := []string{"n", "ne", "se", "s", "sw", "nw"}

	directionReducibles := make(map[string]reducible)
	for _, dir := range directions {
		directionReducibles[dir] = getReducibles(directions, dir)
	}

	dirCounts := make(map[string]int)
	var dir []byte
	var max, total int
	for readErr == nil {
		total = 0
		dir, readErr = reader.ReadBytes(',')
		if readErr != io.EOF {
			checkErr(readErr)
		}
		direction := strings.Trim(string(dir), ",\n ")
		moves := directionReducibles[direction]
		if dirCounts[directions[moves.Opposite]] > 0 {
			dirCounts[directions[moves.Opposite]] = dirCounts[directions[moves.Opposite]] - 1
		} else if dirCounts[directions[moves.Clock]] > 0 {
			dirCounts[directions[moves.Clock]] = dirCounts[directions[moves.Clock]] - 1
			dirCounts[directions[moves.ClockReduce]] = dirCounts[directions[moves.ClockReduce]] + 1
		} else if dirCounts[directions[moves.Counter]] > 0 {
			dirCounts[directions[moves.Counter]] = dirCounts[directions[moves.Counter]] - 1
			dirCounts[directions[moves.CounterReduce]] = dirCounts[directions[moves.CounterReduce]] + 1
		} else {
			dirCounts[direction] = dirCounts[direction] + 1
		}
		for _, v := range dirCounts {
			total += v
		}
		if total > max {
			max = total
		}

	}

	fmt.Println(total)
	fmt.Println(max)
}
