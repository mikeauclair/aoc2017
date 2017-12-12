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
	half := len(ba) / 2
	// prev := ba[len(ba)-1] - ascii_offset
	// fmt.Println(first)
	for i := 0; i < half; i++ {
		cur := ba[i] - ascii_offset
		match := ba[i + half] - ascii_offset
		if cur == match {
			sum = sum + (2 * int(cur))
		}
	}
	fmt.Printf("%v\n", sum)
}
