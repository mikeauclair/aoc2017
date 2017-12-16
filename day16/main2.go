package main

import (
	"bufio"
	"fmt"
	"log"
	"io"
	"os"
	"strings"
	"strconv"
)
const (
	ascii_offset byte = 97
	cycles = 1000000000
	SPIN byte = 's'
	EXCHANGE byte = 'x'
	PARTNER byte = 'p'
	SEQSIZE = 16
)

func checkErr(err error) {
        if err != nil {
                log.Fatalf("Error: %v", err)
        }
}

func getIndex(set []byte, match byte) int {
	for i, b := range set {
		if b == match {
			return i
		}
	}
	return -1
}

func genSeq(size int) []byte {
	ret := make([]byte, 0, size)
	var i byte
	for i = 0; i < byte(size); i++ {
		ret = append(ret, i + ascii_offset)
	}
	return ret
}

func buildSpin(meta []string) move {
	size, err := strconv.Atoi(meta[0])
	checkErr(err)
	return move{Move: SPIN, SpinNum: size}
}

func processSpin(seq []byte, size int) []byte {
	return append(seq[len(seq)-size:], seq[:len(seq)-size]...)
}

func swap(seq []byte, l int, r int) []byte {
	tmp := seq[r]
	seq[r] = seq[l]
	seq[l] = tmp
	return seq
}

func buildExchange(meta []string) move {
	l, err := strconv.Atoi(meta[0])
	checkErr(err)
	r, err := strconv.Atoi(meta[1])
	checkErr(err)
	return move{Move: EXCHANGE, ExchangeNums: [2]int{l,r}}
}

func processExchange(seq []byte, indexes [2]int) []byte {
	return swap(seq, indexes[0], indexes[1])
}

func buildPartner(meta []string) move {
	l := []byte(meta[0])[0]
	r := []byte(meta[1])[0]
	return move{Move: PARTNER, ExchangeChars: [2]byte{l,r}}
}

func processPartner(seq []byte, chars [2]byte) []byte {
	l := getIndex(seq, chars[0])
	r := getIndex(seq, chars[1])
	return swap(seq, l, r)
}

type move struct {
	Move byte
	SpinNum int
	ExchangeNums [2]int
	ExchangeChars [2]byte
}

func seqEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aVal := range a {
		if aVal != b[i] {
			return false
		}
	}
	return true
}

func processMoves(seq []byte, moves []move) []byte {
	for _, runMove := range moves {
		switch runMove.Move {
		case SPIN:
			seq = processSpin(seq, runMove.SpinNum)
		case EXCHANGE:
			seq = processExchange(seq, runMove.ExchangeNums)
		case PARTNER:
			seq = processPartner(seq, runMove.ExchangeChars)
		}

	}
	return seq
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	reader := bufio.NewReader(f)

	var readErr error

	seq := genSeq(SEQSIZE)
	origSeq := genSeq(SEQSIZE)
	moves := []move{}
	
	for readErr == nil {
		moveBytes, readErr := reader.ReadBytes(',')
		if readErr != io.EOF {
			checkErr(readErr)
		}
		if len(moveBytes) == 0 {
			readErr = io.EOF
			break
		}
		var moveEnd int
		if moveBytes[len(moveBytes)-1] == ',' {
			moveEnd = len(moveBytes)-1
		} else {
			moveEnd = len(moveBytes)
		}
		meta := strings.Split(string(moveBytes[1:moveEnd]), "/")
		switch moveBytes[0] {
		case SPIN:
			moves = append(moves, buildSpin(meta))
		case EXCHANGE:
			moves = append(moves, buildExchange(meta))
		case PARTNER:
			moves = append(moves, buildPartner(meta))
		}
	}
	var fullCycle int
	for ; fullCycle < cycles; fullCycle++ {
		seq = processMoves(seq, moves)
		if seqEqual(seq, origSeq) {
			break
		}
	}
	
	for i := 0; i < cycles % (fullCycle + 1); i++{
		seq = processMoves(seq, moves)
	}
		
	
	fmt.Println(string(seq))
	
}
