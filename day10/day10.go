package main

import (
	"aoc21/lib"
	"log"
	"os"
	"sort"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/10

func main() {
	lines := lib.ReadLines(os.Stdin)
	errorScore, completionScore := ScoreLines(lines)
	log.Println("errorScore: ", errorScore)
	log.Println("completionScore", completionScore)
}

func ScoreLines(lines []string) (int, int) {
	n := 0
	var ar []int

	for _, line := range lines {
		a, b := ScoreLine(line)
		n += a
		if b > 0 {
			ar = append(ar, b)
		}
	}

	sort.Ints(ar)

	return n, ar[len(ar)/2]
}

func ScoreLine(line string) (int, int) {
	_, found, completion := ParseLine(line)

	errorScore := 0
	completionScore := 0

	switch found {
	case ")":
		errorScore = 3
	case "]":
		errorScore = 57
	case "}":
		errorScore = 1197
	case ">":
		errorScore = 25137
	}

	for _, c := range completion {
		completionScore *= 5
		switch c {
		case ')':
			completionScore += 1
		case ']':
			completionScore += 2
		case '}':
			completionScore += 3
		case '>':
			completionScore += 4
		}
	}

	return errorScore, completionScore
}

// ParseLine
// returns (expected, found, completion)
// expected/found indicates error in closing delimiter
// completion indicates missing closing delimiters
func ParseLine(line string) (string, string, string) {
	stack := make([]byte, 1000)
	pos := 0

	for _, c := range line {
		b := byte(c)
		switch b {
		case '{':
			stack[pos] = '}'
			pos++
			break

		case '(':
			stack[pos] = ')'
			pos++
			break

		case '[':
			stack[pos] = ']'
			pos++
			break

		case '<':
			stack[pos] = '>'
			pos++
			break

		case '}':
			fallthrough
		case ')':
			fallthrough
		case ']':
			fallthrough
		case '>':
			popped := stack[pos-1]
			pos--
			if popped != b {
				return string(popped), string(b), ""
			}

		default:
			log.Fatal("unexpected input: ", b)
		}

	}

	completion := ""
	for pos = pos - 1; pos >= 0; pos-- {
		completion += string(stack[pos])
	}

	return "", "", completion
}
