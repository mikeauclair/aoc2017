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

func contains(list []string, candidate string) bool {
	for _, elem := range list {
		if elem == candidate {
			return true
		}
	}
	return false
}

func addChildren(list *[]string, newList []string) {
	for _, elem := range newList {
		if !contains(*list, elem) {
			*list = append(*list, elem)
		}
	}
}

func pruneCandidates(list []string, invalid []string) []string {
	candidates := []string{}
	for i, _ := range list {
		if !contains(invalid, list[i]) {
			candidates = append(candidates, list[i])
		}
	}
	return candidates
}

type node struct {
	Weight int
	TotalWeight *int
	Children []string
	Parent *node
}

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	nodes := make(map[string]node)
	for scanner.Scan() {
		halves := strings.Split(scanner.Text(), " -> ")
		
		nodeTokens := strings.Split(halves[0], " ")
		nodeName := nodeTokens[0]
		weightRaw := nodeTokens[1]
		weight, err := strconv.Atoi(weightRaw[1:len(weightRaw)-1])
		checkErr(err)
		newNode := node{Weight: weight, TotalWeight: &weight}
		if len(halves) > 1 {
			list := strings.Split(halves[1], ", ")
			newNode.Children = list
		}
		nodes[nodeName] = newNode
			
	}
	//crib root from previous part
	stack := []string{"airlri"}
	//visited := make(map[string]bool)
	var moves int
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		elem := nodes[cur]
		stack = stack[:len(stack)-1]
		
		n := elem
		for n.Parent != nil {
			n = *n.Parent
			*n.TotalWeight = *n.TotalWeight + elem.Weight
		}
		fmt.Println(n)
		for _, nodeName := range elem.Children {
			n := nodes[nodeName]
			n.Parent = &elem
			nodes[nodeName] = n
			stack = append(stack, nodeName)
		}
		moves++
	}
	cur := "airlri"
	//for cur != "" {
	for cur != "" {
		counts := make(map[int][]string)
		for _, n := range nodes[cur].Children {
			childNode := nodes[n]
			tw := *childNode.TotalWeight
			counts[tw] = append(counts[tw], n)
			fmt.Println(n, tw)
		}
		if len(counts) == 0 || len(counts) == 1 {
			cur = ""
		}
		for _, v := range counts {
			fmt.Println(v)
			if len(v) == 1 {
				fmt.Println("OOOOOO")
				cur = v[0]
			}
		}
		fmt.Println("----------------")
		if len(nodes[cur].Children) == 0 {
			cur = ""
		}
	}
	//}
}
