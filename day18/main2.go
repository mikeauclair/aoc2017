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

type proc struct {
	registers map[byte]int
	curInstr  int
	waitCount int
	queue     []int
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

	var pidOneSend int

	processes := []proc{
		proc{make(map[byte]int), 0, 0, []int{}},
		proc{make(map[byte]int), 0, 0, []int{}},
	}

	for i := 0; i < 2; i++ {
		processes[i].registers['p'] = i
	}

	for processes[0].waitCount < 1 || processes[1].waitCount < 1 {
		for i := 0; i < 2; i++ {
			process := processes[i]
			instr := instructionSet[process.curInstr]
			registers := process.registers
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
					process.curInstr += getRval(instr, registers)
					processes[i] = process
					continue
				}
			case "snd":
				if i == 0 {
					processes[1].queue = append(processes[1].queue, getLval(instr, registers))
				} else {
					processes[0].queue = append(processes[0].queue, getLval(instr, registers))
					pidOneSend++
				}
			case "rcv":
				if len(process.queue) > 0 {
					registers[instr.lreg] = process.queue[0]
					process.queue = process.queue[1:]
				} else {
					process.waitCount++
					processes[i] = process
					continue
				}

			}
			process.waitCount = 0
			process.registers = registers
			process.curInstr = process.curInstr + 1
			processes[i] = process
		}
	}

	fmt.Println(pidOneSend)

}
