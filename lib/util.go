package lib

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

func AssertEqStr(t *testing.T, name string, expected string, actual string) {
	if expected != actual {
		t.Errorf("%s: expected %s, got %s", name, expected, actual)
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

func AssertEqMap(t *testing.T, expected map[string]int, actual map[string]int) {
	match := true
	if len(expected) != len(actual) {
		t.Errorf("len(expected) %d, got %d", len(expected), len(actual))
		match = false
	} else {

		for k, v := range expected {
			if actual[k] != v {
				t.Errorf("expected[%d] %d, got %d", k, v, actual[k])
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

func ReadDigitsArray(rd io.Reader) [][]int {
	var ar [][]int
	lines := ReadLines(rd)
	for _, line := range lines {
		var br []int
		for _, c := range line {
			x := int(c) - int('0')
			if x < 0 || x > 9 {
				log.Fatal("nope")
			}
			br = append(br, x)
		}
		ar = append(ar, br)
	}

	return ar
}
