package lib

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func ReadLines(rd io.Reader) []string {
	var lines []string

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func LinesToNumbers(lines []string, base int) []int {
	var nums []int
	for _, line := range lines {
		num, err := strconv.ParseInt(line, base, 32)
		if err != nil {
			fmt.Errorf("bad input: %s\n", line)
		}
		nums = append(nums, int(num))
	}

	return nums
}

func ReadNumbers(rd io.Reader, base int) []int {
	lines := ReadLines(rd)
	return LinesToNumbers(lines, base)
}
