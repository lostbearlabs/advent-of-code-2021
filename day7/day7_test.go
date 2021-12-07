package main

import (
	"aoc21/lib"
	"testing"
)

var testInput = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func TestExamplePart1(t *testing.T) {
	n := BestPos(testInput, Cost1)
	lib.AssertEq(t, "best pos", 2, n)

	lib.AssertEq(t, "cost at 1", 41, Cost1(testInput, 1))
	lib.AssertEq(t, "cost at 3", 39, Cost1(testInput, 3))
	lib.AssertEq(t, "cost at 10", 71, Cost1(testInput, 10))
}

func TestExamplePart2(t *testing.T) {
	n := BestPos(testInput, Cost2)
	lib.AssertEq(t, "best pos", 5, n)

	lib.AssertEq(t, "cost at 5", 168, Cost2(testInput, 5))
	lib.AssertEq(t, "cost at 2", 206, Cost2(testInput, 2))
}
