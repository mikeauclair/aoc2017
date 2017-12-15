package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
)
const (
	dash byte = 45
	ascii_offset byte = 48
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

func processMove(seq []int, num int, start int) int {
	seqLen := len(seq)
	var tmp int
	for i := 0; i < num/2; i++ {
		frontInd := calcPos(seqLen, start+i)
		backInd := calcPos(seqLen, start+num-i-1)
		tmp = seq[frontInd]
		seq[frontInd] = seq[backInd]
		seq[backInd] = tmp
	}
	return calcPos(seqLen, start+num-1)
}

func calcPos(l int, pos int) int {
	for pos >= l {
		pos -= l
	}
	return pos
}

func getSetBitsForHash(byteArray []byte) int {
	var skipVal, cursor int
	seq := fillSeq(255)
	for i := 0; i < 64; i++ {
		for _, move := range byteArray {
			cursor = processMove(seq, int(move), calcPos(len(seq), skipVal + cursor))
			skipVal++
		}
	}
	hash := make([]byte, 16)
	var sum byte
	for i:=0; i < 16; i++ {
		for j:=0; j < 16; j++ {
			hash[i] = hash[i] ^ byte(seq[i*16+j])
		}
		v := hash[i]
		for ; v != 0; v >>= 1 {
			sum += v & 1
		}
	}

	fmt.Println(hex.EncodeToString(hash), sum)
	
	return int(sum)
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	_ = scanner.Scan()
	baseMoves := scanner.Bytes()
	var sum int

	trailingBytes := []byte{17, 31, 73, 47, 23}
	
	for i := 0; i < 128; i++ {
		numberBytes := []byte(strconv.Itoa(i))
		calcMoves := append(baseMoves, dash)
		calcMoves = append(calcMoves, numberBytes...)
		calcMoves = append(calcMoves, trailingBytes...)
		sum += getSetBitsForHash(calcMoves)
		fmt.Println(string(calcMoves), sum)
	}

	fmt.Println(sum)
	
        
}
