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
	registers := make(map[string]int)
	var runOp bool
	var max int
	for scanner.Scan() {
		halves := strings.Split(scanner.Text(), " if ")

		condTokens := strings.Split(halves[1], " ")
		condReg := registers[condTokens[0]]
		condOp := condTokens[1]
		condVal, err := strconv.Atoi(condTokens[2])
		checkErr(err)
		
		switch condOp {
		case "!=":
			runOp = (condReg != condVal)
		case "==":
			runOp = (condReg == condVal)
		case "<":
			runOp = (condReg < condVal)
		case "<=":
			runOp = (condReg <= condVal)
		case ">":
			runOp = (condReg > condVal)
		case ">=":
			runOp = (condReg >= condVal)
		}

		if !runOp {
			continue
		}
		
		opTokens := strings.Split(halves[0], " ")
		opReg := registers[opTokens[0]]
		val, err := strconv.Atoi(opTokens[2])
		checkErr(err)
		op := opTokens[1]
		if op == "inc" {
			opReg += val
		} else {
			opReg -= val
		}
		if opReg > max {
			max = opReg
		}
		registers[opTokens[0]] = opReg
	}
	fmt.Println(registers)
	fmt.Println(max)
}
