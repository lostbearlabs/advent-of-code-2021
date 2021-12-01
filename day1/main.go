package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// https://adventofcode.com/2021/day/1

// Usage:
//   go run . < input

func main() {
	numIncrease, numWindowIncrease := Count(os.Stdin)

	fmt.Printf("numIncrease: %d\n", numIncrease)
	fmt.Printf("numWindowIncrease: %d\n", numWindowIncrease)
}

func Count(rd io.Reader) (int, int) {
	lineNum := 1
	numIncrease := 0
	numWindowIncrease := 0
	sum := 0
	prev := 0
	prevSum := 0
	window := [3]int{0, 0, 0}

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		line := scanner.Text()
		n, _ := strconv.Atoi(line)

		idx := lineNum % 3
		sum -= window[idx]
		window[idx] = n
		sum += n

		if lineNum > 1 && n > prev {
			numIncrease++
		}

		if lineNum > 3 && sum > prevSum {
			numWindowIncrease++
		}

		lineNum++
		prev = n
		prevSum = sum
	}

	return numIncrease, numWindowIncrease
}
