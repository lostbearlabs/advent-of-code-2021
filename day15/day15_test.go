package main

import (
	"aoc21/lib"
	"strings"
	"testing"
)

func sample() string {
	return `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`
}

func TestSample(t *testing.T) {
	rd := strings.NewReader(sample())
	grid := lib.ReadDigitsArray(rd)
	cost := ShortestCornerToCornerCost(grid)
	lib.AssertEq(t, "cost", 40, cost)
}

func TestExpandedSample(t *testing.T) {
	rd := strings.NewReader(sample())
	grid := lib.ReadDigitsArray(rd)
	grid = Expand(grid)
	cost := ShortestCornerToCornerCost(grid)
	lib.AssertEq(t, "cost", 315, cost)
}
