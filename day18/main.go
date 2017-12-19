package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

type instruction struct {
	instrname string
	lreg      byte
	lval      int
	rreg      byte
	rval      int
}

func getRval(instr instruction, registers map[byte]int) int {
	if instr.rreg > 0 {
		return registers[instr.rreg]
	}
	return instr.rval
}

func getLval(instr instruction, registers map[byte]int) int {
	if instr.lreg > 0 {
		return registers[instr.lreg]
	}
	return instr.lval
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var rcvVal, curInstr int
	instructionSet := []instruction{}
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), " ")
		newInstr := instruction{instrname: vals[0]}
		if vals[1][0] >= 97 {
			newInstr.lreg = vals[1][0]
		} else {
			intVal, err := strconv.Atoi(vals[1])
			checkErr(err)
			newInstr.lval = intVal
		}
		if len(vals) > 2 {
			if vals[2][0] >= 97 {
				newInstr.rreg = vals[2][0]
			} else {
				intVal, err := strconv.Atoi(vals[2])
				checkErr(err)
				newInstr.rval = intVal
			}
		}
		instructionSet = append(instructionSet, newInstr)

	}
	registers := make(map[byte]int)
	var lsound int
	for rcvVal == 0 {
		instr := instructionSet[curInstr]
		switch instr.instrname {
		case "set":
			registers[instr.lreg] = getRval(instr, registers)
		case "mul":
			registers[instr.lreg] = registers[instr.lreg] * getRval(instr, registers)
		case "add":
			registers[instr.lreg] = registers[instr.lreg] + getRval(instr, registers)
		case "mod":
			registers[instr.lreg] = registers[instr.lreg] % getRval(instr, registers)
		case "jgz":
			if getLval(instr, registers) > 0 {
				curInstr += getRval(instr, registers)
				continue
			}
		case "snd":
			lsound = getLval(instr, registers)
		case "rcv":
			if getLval(instr, registers) > 0 {
				rcvVal = lsound
			}
		}
		curInstr++
	}

	fmt.Println(rcvVal)

}
