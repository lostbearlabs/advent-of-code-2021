package main

import (
	"aoc21/lib"
	"strings"
	"testing"
)

func sampleInput() string {
	return `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`
}

func TestSampleOneFold(t *testing.T) {
	rd := strings.NewReader(sampleInput())
	fnord := ReadFnord(rd)
	mp := DoFold(fnord.points, fnord.instructions[0])
	lib.AssertEq(t, "num points", 17, len(mp))
}

func TestSampleAllFolds(t *testing.T) {
	rd := strings.NewReader(sampleInput())
	fnord := ReadFnord(rd)

	Display(fnord.points)

	mp := DoFold(fnord.points, fnord.instructions[0])
	lib.AssertEq(t, "num points", 17, len(mp))
	Display(mp)

	mp = DoFold(mp, fnord.instructions[1])
	Display(mp)

	mp = DoFold(mp, fnord.instructions[1])
	Display(mp)

}
