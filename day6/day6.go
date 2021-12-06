package main

import (
	"aoc21/lib"
	"fmt"
	"os"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/6

func main() {
	lines := lib.ReadLines(os.Stdin)
	input := lib.LinesToNumbersSep(lines, 10, ",")
	counts := InputToCounts(input)

	// part 1
	for i := 0; i < 80; i++ {
		counts = NextDay(counts)
	}
	sum := SumCounts(counts)
	fmt.Printf("After 80 days: %d\n", sum)

	// part 2
	counts = InputToCounts(input)
	for i := 0; i < 256; i++ {
		counts = NextDay(counts)
	}
	sum = SumCounts(counts)
	fmt.Printf("After 256 days: %d\n", sum)
}

func InputToCounts(ar []int) []int {
	counts := make([]int, 9)
	for _, n := range ar {
		counts[n]++
	}
	return counts
}

func SumCounts(counts []int) int {
	sum := 0
	for _, n := range counts {
		sum += n
	}
	return sum
}

func NextDay(counts []int) []int {
	next := make([]int, 9)
	for i, n := range counts {
		if i == 0 {
			next[6] += n
			next[8] += n
		} else {
			next[i-1] += n
		}
	}

	return next
}
