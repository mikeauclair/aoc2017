package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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
	buf := bytes.NewBuffer(nil)
	f, err := os.Open("input.txt")
	checkErr(err)
	io.Copy(buf, f)
	f.Close()
	ba := buf.Bytes()
	var sum int
	first := ba[0]
	prev := ba[len(ba)-1] - ascii_offset
	fmt.Println(first)
	for _, elem := range ba {
		cur := elem - ascii_offset
		if cur == prev {
			sum = sum + int(cur)
		}
		prev = cur
	}
	fmt.Println(string(first))
	fmt.Println(string(ba))
	fmt.Printf("%v\n", sum)
}
