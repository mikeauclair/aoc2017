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
	
)

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func fillSeq(max int) []int {
	fill := make([]int, 0, max+1)
	for i := 0; i <= max; i++ {
		fill = append(fill, i)
	}
	return fill
}

func processMove(seq *[]int, num int, start int) int {
	seqLen := len(*seq)
	var tmp int
	for i := 0; i < num/2; i++ {
		frontInd := calcPos(seqLen, start+i)
		backInd := calcPos(seqLen, start+num-i-1)
		tmp = (*seq)[frontInd]
		(*seq)[frontInd] = (*seq)[backInd]
		(*seq)[backInd] = tmp
	}
	return calcPos(seqLen, start+num-1)
}

func calcPos(l int, pos int) int {
	for pos >= l {
		pos -= l
	}
	return pos
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	_ = scanner.Scan()
	var skipVal, cursor int
	moves := strings.Split(scanner.Text(), ",")
	
        seq := fillSeq(255)

	for _, move := range moves {
		intVal, err := strconv.Atoi(move)
		checkErr(err)
		fmt.Println(seq, skipVal, cursor, intVal)
		cursor = processMove(&seq, intVal, calcPos(len(seq), skipVal + cursor))
		skipVal++
	}

	
	fmt.Println(seq)
}
