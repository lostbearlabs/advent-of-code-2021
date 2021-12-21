package main

import (
	"aoc21/lib"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/14

func main() {
	fnord := ReadFnord(os.Stdin)

	st := fnord.Polymer
	for i := 1; i <= 10; i++ {
		st = UpdatePolymer(st, fnord.Rules)

		if i == 10 {
			counts := GetCharCountsFromString(st)
			log.Printf("*** i=%d ***\n", i)
			log.Println("smallest count", counts[0])
			log.Println("largest count", counts[len(counts)-1])
			log.Println("delta", counts[len(counts)-1].Count-counts[0].Count)
		}
	}

	mp := ToPairMap(fnord.Polymer)
	for i := 1; i <= 40; i++ {
		mp = UpdatePairMap(mp, fnord.Rules)

		if i == 10 || i == 40 {
			counts := GetCharCountsFromPairMap(mp, fnord.Polymer)
			log.Printf("*** i=%d ***\n", i)
			log.Println("smallest count", counts[0])
			log.Println("largest count", counts[len(counts)-1])
			log.Println("delta", counts[len(counts)-1].Count-counts[0].Count)
		}
	}
}

// ******************
// Fnord

type Fnord struct {
	Polymer string
	Rules   map[string]string
}

func ReadFnord(rd io.Reader) Fnord {
	fnord := Fnord{}

	lines := lib.ReadLines(rd)
	fnord.Polymer = lines[0]
	fnord.Rules = make(map[string]string)

	for i := 2; i < len(lines); i++ {
		ar := strings.Split(lines[i], " -> ")
		fnord.Rules[ar[0]] = ar[1]
	}

	return fnord
}

// ******************
// CharCount

type CharCount struct {
	Name  string
	Count int
}

func GetCharCountsFromString(st string) []CharCount {
	countMap := make(map[string]int)
	for _, c := range st {
		countMap[string(c)] += 1
	}

	var ar []CharCount
	for name, count := range countMap {
		ar = append(ar, CharCount{Name: name, Count: count})
	}

	sort.Slice(ar, func(i, j int) bool {
		return ar[i].Count < ar[j].Count
	})

	return ar
}

// GetCharCountsFromPairMap
// Converts pair map into char counts.
// Needs original string to account for the fact that every character is double-counted,
// appearing in the pairs that begin with it and also the pairs that end with it, except
// for the very first and last characters of the string.
func GetCharCountsFromPairMap(mp map[string]int, origString string) []CharCount {

	countMap := make(map[string]int)
	for pair, count := range mp {
		countMap[string(pair[0])] += count
		countMap[string(pair[1])] += count
	}

	// account for first and last char not being double-counted yet
	countMap[string(origString[0])] += 1
	countMap[string(origString[len(origString)-1])] += 1

	var ar []CharCount
	for name, count := range countMap {
		// /2 here since every character is double counted
		ar = append(ar, CharCount{Name: name, Count: count / 2})
	}

	sort.Slice(ar, func(i, j int) bool {
		return ar[i].Count < ar[j].Count
	})

	return ar
}

// ******************
// PlayTurn

// UpdatePolymer
// Brute-force solution which just creates a new string from the old one.
// Does not scale to part 2.
func UpdatePolymer(st string, rules map[string]string) string {

	prev := ""
	tt := ""

	for _, c := range st {
		cc := string(c)
		key := prev + cc
		tt = tt + rules[key] + cc
		prev = cc
	}

	return tt
}

// ******************
// Pair Stuff

func ToPairMap(st string) map[string]int {
	mp := make(map[string]int)

	prev := ""

	for _, c := range st {
		cc := string(c)
		key := prev + cc
		mp[key]++
		prev = cc
	}

	return mp
}

// UpdatePairMap
// Efficient solution that just keeps track of how many times each pair occurs
// in the string.
func UpdatePairMap(mp map[string]int, rules map[string]string) map[string]int {
	nt := make(map[string]int)

	for pair, count := range mp {
		between := rules[pair] // character to insert in middle of pair
		if between != "" {
			first := string(pair[0]) + between  // first replacement pair
			second := between + string(pair[1]) // second replacement pair

			nt[first] += count
			nt[second] += count
		}
	}

	return nt
}
