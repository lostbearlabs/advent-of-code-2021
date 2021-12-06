package main

import (
	"aoc21/lib"
	"testing"
)

func givenTestInput() []int {
	input := []string{"3,4,3,1,2"}
	return lib.LinesToNumbersSep(input, 10, ",")
}

func TestInputToCounts(t *testing.T) {
	input := givenTestInput()
	counts := InputToCounts(input)
	expected := []int{0, 1, 1, 2, 1, 0, 0, 0, 0}
	lib.AssertEqAr(t, "counts", expected, counts)
}

func TestSumCounts(t *testing.T) {
	input := givenTestInput()
	counts := InputToCounts(input)
	sum := SumCounts(counts)
	lib.AssertEq(t, "count", 5, sum)
}

func TestNextDay(t *testing.T) {
	counts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	next := NextDay(counts)
	expected := []int{2, 3, 4, 5, 6, 7, 9, 9, 1}
	lib.AssertEqAr(t, "expected", expected, next)
}

func TestSampleInput18(t *testing.T) {
	input := givenTestInput()
	counts := InputToCounts(input)

	for i := 0; i < 18; i++ {
		counts = NextDay(counts)
	}
	sum := SumCounts(counts)
	lib.AssertEq(t, "sum after 18 days", 26, sum)
}

func TestSampleInput80(t *testing.T) {
	input := givenTestInput()
	counts := InputToCounts(input)

	for i := 0; i < 80; i++ {
		counts = NextDay(counts)
	}
	sum := SumCounts(counts)
	lib.AssertEq(t, "sum after 80 days", 5934, sum)
}

func TestSampleInput256(t *testing.T) {
	input := givenTestInput()
	counts := InputToCounts(input)

	for i := 0; i < 256; i++ {
		counts = NextDay(counts)
	}
	sum := SumCounts(counts)
	lib.AssertEq(t, "sum after 256 days", 26984457539, sum)
}
