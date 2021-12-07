package main

import (
	"aoc21/lib"
	"fmt"
	"os"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/7

func main() {
	lines := lib.ReadLines(os.Stdin)
	input := lib.LinesToNumbersSep(lines, 10, ",")

	// part1
	best := BestPos(input, Cost1)
	fmt.Printf("best pos 1: %d\n", best)
	fmt.Printf("best cost 1: %d\n", Cost1(input, best))

	// part2
	best = BestPos(input, Cost2)
	fmt.Printf("best pos 2: %d\n", best)
	fmt.Printf("best cost 2: %d\n", Cost2(input, best))

}

// credit: https://stackoverflow.com/a/45976758/4540
func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func Cost1(ar []int, pos int) int {
	cost := 0

	for _, x := range ar {
		cost += AbsDiff(x, pos)
	}

	return cost
}

func Cost2(ar []int, pos int) int {
	cost := 0

	for _, x := range ar {
		cost += SumTo(AbsDiff(x, pos))
	}

	return cost
}

func SumTo(n int) int {
	return n * (n + 1) / 2
}

func AbsDiff(x int, y int) int {
	diff := x - y
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func BestPos(ar []int, cost func([]int, int) int) int {
	min, max := MinMax(ar)
	best := min
	best_cost := cost(ar, min)

	fmt.Printf("min=%d, max=%d\n", min, max)

	for i := min + 1; i <= max; i++ {
		cost := cost(ar, i)
		if cost < best_cost {
			best = i
			best_cost = cost
		}
	}

	return best
}
