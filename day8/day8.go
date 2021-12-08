package main

import (
	"aoc21/lib"
	"fmt"
	"log"
	"os"
	"strings"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/8

func main() {
	lines := lib.ReadLines(os.Stdin)
	entries := ParseEntries(lines)

	easyCount := EasyCount(entries)
	log.Println("easy count ", easyCount)

	hardCount := GetTotalOutput(entries)
	log.Println("hard count ", hardCount)
}

type Entry struct {
	Input  []string // 10 digits
	Output []string // 4 digits
}

func ParseEntry(line string) Entry {
	fields := strings.Fields(line)

	if len(fields) != 15 || fields[10] != "|" {
		log.Fatal("bad line: ", line)
	}

	for i, st := range fields {
		if st != "|" {
			fields[i] = canonicalString(st)
		}
	}
	entry := Entry{
		Input:  fields[:10],
		Output: fields[11:],
	}
	return entry
}

func ParseEntries(lines []string) []Entry {
	var entries []Entry
	for _, line := range lines {
		entries = append(entries, ParseEntry(line))
	}
	return entries
}

func EasyCount(entries []Entry) int {
	sum := 0
	for _, entry := range entries {
		for _, st := range entry.Output {
			x := len(st)
			if x == 2 || x == 3 || x == 4 || x == 7 {
				sum++
			}
		}
	}
	return sum
}

func DeduceMapping(entry Entry) map[string]int {

	// The digits with unique numbers of segments are 1, 4, 7, and 8
	// For example, number 1 is the only digit with two segments, and those are c and f
	cf := findDigitWithGivenLength(entry, 2)  // 1
	acf := findDigitWithGivenLength(entry, 3) // 7

	// The 5-segment digits are 2 (acdeg), 3 (acdfg), and 6 (abdfg)
	adg := intersect(findDigitsWithGivenLength(entry, 5)...)

	// The 6-segment digits are 0 (abcdef), 6 (abdefg), and 9 (abcdef)
	abfg := intersect(findDigitsWithGivenLength(entry, 6)...)

	a := diff(acf, cf)
	c := diff(acf, abfg)
	f := diff(cf, c)
	cde := not(abfg)
	bcef := not(adg)
	b := diff(diff(bcef, cde), f)
	e := diff(diff(diff(bcef, b), c), f)
	d := diff(diff(cde, c), e)
	g := diff(diff(adg, a), d)

	// Now that we know how each segment is mapped, we can return a mapping
	// for each digit from the segments representing that digit to the digit itself.
	// The variable name here is the segment name in the original mapping, but
	// the variable value is the segment name in the new mapping.
	return map[string]int{
		union(a, b, c, e, f, g):    0,
		union(c, f):                1,
		union(a, c, d, e, g):       2,
		union(a, c, d, f, g):       3,
		union(b, c, d, f):          4,
		union(a, b, d, f, g):       5,
		union(a, b, d, e, f, g):    6,
		union(a, c, f):             7,
		union(a, b, c, d, e, f, g): 8,
		union(a, b, c, d, f, g):    9,
	}

}

func diff(a string, b string) string {
	return asStr(asInt(a) & ^asInt(b))
}

func intersect(ar ...string) string {
	n := 0b1111111
	for _, a := range ar {
		n = n & asInt(a)
	}
	return asStr(n)
}

func not(a string) string {
	return asStr(^asInt(a) & 0b1111111)
}

func union(ar ...string) string {
	n := 0
	for _, st := range ar {
		n |= asInt(st)
	}
	return asStr(n)
}

func asInt(st string) int {
	n := 0
	for _, c := range st {
		n |= 1 << (c - 'a')
	}

	return n
}

// convert input strings to canonical form
func canonicalString(st string) string {
	return asStr(asInt(st))
}

func asStr(n int) string {
	st := ""
	for i := 0; i < 8; i++ {
		if n&(1<<i) != 0 {
			st += string(rune('a' + i))
		}
	}
	return st
}

// find all unique words with the desired length
func findDigitsWithGivenLength(entry Entry, tgtLen int) []string {
	mp := make(map[string]int)

	for _, word := range entry.Input {
		if len(word) == tgtLen {
			mp[word]++
		}
	}
	for _, word := range entry.Output {
		if len(word) == tgtLen {
			mp[word]++
		}
	}

	if len(mp) == 0 {
		log.Fatal("NO EXAMPLE FOUND FOR LEN ", tgtLen)
	}

	var ar []string
	for k := range mp {
		ar = append(ar, k)
	}
	log.Println("values with length ", tgtLen, " : ", ar)
	return ar
}

// find first word with the desired length
func findDigitWithGivenLength(entry Entry, tgtLen int) string {
	return findDigitsWithGivenLength(entry, tgtLen)[0]
}

func GetOutput(entry Entry, mp map[string]int) int {
	n := 0
	mult := 1
	for i := 0; i < 4; i++ {
		n += mp[entry.Output[4-i-1]] * mult
		mult *= 10
	}
	fmt.Println(entry.Output, "  -->  ", n)
	return n
}

func GetTotalOutput(entries []Entry) int {
	n := 0
	for _, e := range entries {
		mp := DeduceMapping(e)
		n += GetOutput(e, mp)
	}
	return n
}
