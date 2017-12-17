package main

import (
	"fmt"
)

const (
	steps  = 303
	rounds = 2017
)

func main() {
	var cur int
	buffer := []int{0}
	for i := 1; i <= rounds; i++ {
		cur = (cur + steps) % len(buffer) + 1
		buffer = append(buffer, 0)
		copy(buffer[cur+1:], buffer[cur:])
		buffer[cur] = i	
	}
	cur = (cur + 1) % len(buffer)
	fmt.Println(buffer[cur])
}
