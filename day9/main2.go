package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)
const (
	openGroup byte = 123
	closeGroup byte = 125
	openGarbage byte = 60
	closeGarbage byte = 62
	cancelChar byte = 33
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
	reader := bufio.NewReader(f)

	stackDepth := 0
	garbageOpen := false
	cancelNext := false
	total := 0
	var b byte
	for err == nil {
		b, err = reader.ReadByte()
		if err != nil {
			break
		}
		if cancelNext {
			cancelNext = false
			continue
		}
		switch b {
		case '{':
			if garbageOpen {
				total++
			} else {
				stackDepth++
			}
		case '}':
			if garbageOpen {
				total++
			} else {
				stackDepth--
			}
		case '<':
			if garbageOpen {
				total++
			} else {
				garbageOpen = true
			}
		case '>':
			garbageOpen = false
		case '!':
			cancelNext = true
		default:
			if garbageOpen {
				total++
			}
		}
	}
	if err != io.EOF {
		checkErr(err)
	}
	fmt.Println(total)
	
		
}
