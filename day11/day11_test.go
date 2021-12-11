package main

import (
	"aoc21/lib"
	"strings"
	"testing"
)

func TestSmallExample(t *testing.T) {
	input := `11111
19991
19191
19991
11111`
	reader := strings.NewReader(input)
	ar := lib.ReadDigitsArray(reader)
	lib.AssertEq(t, "after 1", 9, CountFlashes(ar, 1))
}

func TestBigExample10(t *testing.T) {
	input := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	reader := strings.NewReader(input)
	ar := lib.ReadDigitsArray(reader)
	lib.AssertEq(t, "after 10", 204, CountFlashes(ar, 10))
	//lib.AssertEq(t, "after 100", 1656, CountFlashes(ar, 100))
}

func TestBigExample100(t *testing.T) {
	input := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	reader := strings.NewReader(input)
	ar := lib.ReadDigitsArray(reader)
	lib.AssertEq(t, "after 100", 1656, CountFlashes(ar, 100))
}

func TestBigExampleAll(t *testing.T) {
	input := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	reader := strings.NewReader(input)
	ar := lib.ReadDigitsArray(reader)
	lib.AssertEq(t, "after 100", 195, AllFlash(ar))
}
