package main

import (
	"aoc21/lib"
	"log"
	"os"
	"strings"
	"unicode"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/12

func main() {
	lines := lib.ReadLines(os.Stdin)
	graph := ParseGraph(lines)

	count := CountPaths(graph)
	log.Println("count: ", count)

	count2 := CountPaths2(graph)
	log.Println("count2: ", count2)
}

func ParseGraph(lines []string) map[string][]string {
	graph := make(map[string][]string)
	for _, line := range lines {
		ar := strings.Split(line, "-")
		graph[ar[0]] = append(graph[ar[0]], ar[1])
		graph[ar[1]] = append(graph[ar[1]], ar[0])
	}
	return graph
}

type State struct {
	visited      map[string]bool
	allowRevisit bool
	revisited    string
}

func NewState() State {
	return State{
		visited: make(map[string]bool),
	}
}

func CountPaths(graph map[string][]string) int {
	state := NewState()
	return CountPathsFrom(graph, "start", state)
}

func CountPaths2(graph map[string][]string) int {
	state := NewState()
	state.allowRevisit = true
	return CountPathsFrom(graph, "start", state)
}

func CountPathsFrom(graph map[string][]string, node string, state State) int {
	n := 0

	//log.Println(pad, node)
	if node == "end" {
		return 1
	}

	if unicode.IsLower(rune(node[0])) {
		if state.visited[node] {
			if state.allowRevisit {
				if state.revisited != "" || node == "start" || node == "end" {
					return 0
				} else {
					state.revisited = node
				}
			} else {
				return 0
			}
		} else {
			state.visited[node] = true
		}
	}

	for _, next := range graph[node] {
		n += CountPathsFrom(graph, next, state)
	}

	if state.revisited == node {
		state.revisited = ""
	} else {
		state.visited[node] = false
	}

	return n
}
