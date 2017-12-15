package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)
const (
	dash byte = 45
	ascii_offset byte = 48
)

type regionResult struct {
	
}

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

func getBitArrayForHash(byteArray []byte) []bool {
	var skipVal, cursor int
	seq := fillSeq(255)
	for i := 0; i < 64; i++ {
		for _, move := range byteArray {
			cursor = processMove(seq, int(move), calcPos(len(seq), skipVal + cursor))
			skipVal++
		}
	}
	hash := make([]byte, 16)
	bits := make([]bool, 0, 128)
	for i:=0; i < 16; i++ {
		for j:=0; j < 16; j++ {
			hash[i] = hash[i] ^ byte(seq[i*16+j])
		}
		v := hash[i]
		var j uint
		for j = 0; j < 8 ; j++ {
			bit := ((v & (1 << (7-j))) != 0)
			bits = append(bits, bit)
		}
	}
	return bits
}

func dfs(x int, y int, field [][]bool, visited map[[2]int]bool) {
	directions := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
	stack := [][2]int{{x,y}}
	for len(stack) > 0 {
		pos := stack[0]
		stack = stack[1:]
		visited[[2]int{pos[0],pos[1]}] = true
		for _, direction := range directions {
			moveX, moveY := direction[0], direction[1]
			point := [2]int{pos[0]+moveX, pos[1]+moveY}
			if point[0] < 0 || point[0] > 127 || point[1] < 0 || point[1] > 127 {
				continue
			}
			if !visited[point] && field[point[0]][point[1]] {
				stack = append(stack, point)
			}
		}
	}
	
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	_ = scanner.Scan()
	baseMoves := scanner.Bytes()

	trailingBytes := []byte{17, 31, 73, 47, 23}

	arraySet := [][]bool{}
	
	for i := 0; i < 128; i++ {
		numberBytes := []byte(strconv.Itoa(i))
		calcMoves := append(baseMoves, dash)
		calcMoves = append(calcMoves, numberBytes...)
		calcMoves = append(calcMoves, trailingBytes...)
		arraySet = append(arraySet, getBitArrayForHash(calcMoves))
	}

	visited := make(map[[2]int]bool)
	var islands int
	for i, row := range arraySet {
		for j, point := range row {
			if !point {
				continue
			}
			if visited[[2]int{i,j}] {
				continue
			}
			islands++
			dfs(i, j, arraySet, visited)
		}
		
	}
	fmt.Println(islands)
	fmt.Println(len(visited))
		
	
}
