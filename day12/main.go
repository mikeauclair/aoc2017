package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const ()

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

type node struct {
	number int
	edges  map[int]bool
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	nodeList := make(map[int]node)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " <-> ")
		nodeNum, err := strconv.Atoi(line[0])
		checkErr(err)
		edges := strings.Split(line[1], ", ")
		for _,edge := range edges {
			edgeNum, err := strconv.Atoi(edge)
			checkErr(err)
			edgeNode, ok := nodeList[edgeNum]
			if ok {
				edgeNode.edges[nodeNum] = true
			} else {
				edgeNode = node{edgeNum, map[int]bool{nodeNum: true}}
			}
			nodeList[edgeNum] = edgeNode
		}
	}
	visited := make(map[int]bool)
	stack := []int{0}
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		if visited[cur] {
			continue
		}
		visited[cur] = true
		curNode := nodeList[cur]
		for edge, _ := range curNode.edges {
			stack = append(stack, edge)
		}
			
	}
	checkErr(scanner.Err())
	fmt.Println(len(visited))
}
