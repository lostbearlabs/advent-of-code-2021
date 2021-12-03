package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/2

// Usage:
//   go run . < input

func main() {
	horz, depth, horz2, depth2 := Count(os.Stdin)

	fmt.Printf("horz: %d\n", horz)
	fmt.Printf("depth: %d\n", depth)
	fmt.Printf("horz*depth: %d\n", horz*depth)

	fmt.Printf("horz2: %d\n", horz2)
	fmt.Printf("depth2: %d\n", depth2)
	fmt.Printf("horz2*depth2: %d\n", horz2*depth2)

}

func Count(rd io.Reader) (int, int, int, int) {
	horz1 := 0
	depth1 := 0

	aim2 := 0
	horz2 := 0
	depth2 := 0

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			fmt.Printf("bad line: %s\n", line)
			os.Exit(1)
		}

		delta, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("bad line: %s\n", line)
			os.Exit(1)
		}

		if parts[0] == "down" {
			depth1 += delta

			aim2 += delta
		} else if parts[0] == "up" {
			depth1 -= delta

			aim2 -= delta
		} else if parts[0] == "forward" {
			horz1 += delta

			horz2 += delta
			depth2 += aim2 * delta
		} else {
			fmt.Printf("bad line: %s\n", line)
			os.Exit(1)
		}

	}

	return horz1, depth1, horz2, depth2
}
