package main

import (
	"aoc21/lib"
	"strings"
	"testing"
)

func sample() string {
	return `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`

	//	return `...>...
	//.......
	//......>
	//v.....>
	//......>
	//.......
	//..vvv..`
}

func TestGrid_NumSteps(t *testing.T) {
	rd := strings.NewReader(sample())
	grid := ReadGrid(rd)
	n := grid.NumSteps()
	lib.AssertEq(t, "", 58, n)
}
