package lib

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"
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

func AssertEq(t *testing.T, name string, expected int, actual int) {
	if expected != actual {
		t.Errorf("%s: expected %d, got %d", name, expected, actual)
	}
}

func AssertEqAr(t *testing.T, name string, expected []int, actual []int) {
	match := true
	if len(expected) != len(actual) {
		t.Errorf("%s: len(expected) %d, got %d", name, len(expected), len(actual))
		match = false
	} else {
		for i, exp := range expected {
			if exp != actual[i] {
				t.Errorf("%s: expected[%d] %d, got %d", name, i, exp, actual[i])
				match = false
			}
		}
	}

	if !match {
		fmt.Println("expected: ", expected)
		fmt.Println("actual: ", actual)
	}
}

func LinesToNumbersSep(lines []string, base int, sep string) []int {
	var nums []int
	for _, line := range lines {

		var ar []string
		if sep == " " {
			// special case ... ignore repeated spaces
			ar = strings.Fields(line)
		} else {
			ar = strings.Split(line, sep)
		}

		for _, field := range ar {
			num, err := strconv.ParseInt(field, base, 32)
			if err != nil {
				fmt.Errorf("bad input: %s\n", line)
			}
			nums = append(nums, int(num))
		}
	}

	return nums
}
