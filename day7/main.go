package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func main() {
	f, err := os.Open("input.txt")
	checkErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	notRoot := []string{}
	maybeRoot := []string{}
	for scanner.Scan() {
		halves := strings.Split(scanner.Text(), " -> ")
		
		nodeTokens := strings.Split(halves[0], " ")
		node := nodeTokens[0]
		weight := nodeTokens[1]
		fmt.Println(weight)
		if len(halves) > 1 {
			list := strings.Split(halves[1], ", ")
			addChildren(&notRoot, list)
		}
		maybeRoot = append(maybeRoot, node)
		maybeRoot = pruneCandidates(maybeRoot, notRoot)
	}
	fmt.Println(maybeRoot)
}
