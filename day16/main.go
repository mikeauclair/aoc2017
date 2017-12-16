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
	ascii_int_offset
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

func processSpin(seq []byte, meta []string) []byte {
	size, err := strconv.Atoi(meta[0])
	checkErr(err)
	return append(seq[len(seq)-size:], seq[:len(seq)-size]...)
}

func swap(seq []byte, l int, r int) []byte {
	tmp := seq[r]
	seq[r] = seq[l]
	seq[l] = tmp
	return seq
}

func processExchange(seq []byte, meta []string) []byte {
	l, err := strconv.Atoi(meta[0])
	checkErr(err)
	r, err := strconv.Atoi(meta[1])
	checkErr(err)
	return swap(seq, l, r)
}

func processPartner(seq []byte, meta []string) []byte {
	l := getIndex(seq, []byte(meta[0])[0])
	r := getIndex(seq, []byte(meta[1])[0])
	return swap(seq, l, r)
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	reader := bufio.NewReader(f)

	var readErr error

	seq := genSeq(16)
	
	for readErr == nil {
		move, readErr := reader.ReadBytes(',')
		if readErr != io.EOF {
			checkErr(readErr)
		}
		if len(move) == 0 {
			readErr = io.EOF
			break
		}
		var moveEnd int
		if move[len(move)-1] == ',' {
			moveEnd = len(move)-1
		} else {
			moveEnd = len(move)
		}
		meta := strings.Split(string(move[1:moveEnd]), "/")
		switch move[0] {
		case 's':
			seq = processSpin(seq, meta)
		case 'x':
			seq = processExchange(seq, meta)
		case 'p':
			seq = processPartner(seq, meta)
		}
	}
	fmt.Println(string(seq))
	
}
