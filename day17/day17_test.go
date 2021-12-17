package main

import (
	"aoc21/lib"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := "target area: x=20..30, y=-10..-5"
	target := ParseInput(input)
	lib.AssertEq(t, "x0", 20, target.x0)
	lib.AssertEq(t, "y0", -10, target.y0)
	lib.AssertEq(t, "x1", 30, target.x1)
	lib.AssertEq(t, "y1", -5, target.y1)
}

func TestSample(t *testing.T) {
	input := "target area: x=20..30, y=-10..-5"
	target := ParseInput(input)
	ymax, found := FindHighestHit(target)
	lib.AssertEq(t, "ymax", 45, ymax)
	lib.AssertEq(t, "len(found)", 112, len(found))
}
