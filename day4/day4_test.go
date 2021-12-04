package main

import (
	"aoc21/lib"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

	// Part 1
	rd := strings.NewReader(input)
	lines := lib.ReadLines(rd)
	moves := lib.LinesToNumbersSep(lines[0:1], 10, ",")
	lib.AssertEq(t, "len(moves)", 27, len(moves))

	cells := lib.LinesToNumbersSep(lines[2:], 10, " ")
	lib.AssertEq(t, "len(cells)", 75, len(cells))
	boards := NewBoards(cells, 5)
	lib.AssertEq(t, "len(boards)", 3, len(boards))

	winner, last := PlayGame(boards, moves)
	lib.AssertEq(t, "winning move", 24, last)
	lib.AssertEq(t, "first square of winning board 3", 14, winner.cells[0])

	unmarked := SumUnmarked(winner)
	DisplayBoard(winner)
	lib.AssertEq(t, "sum unmarked", 188, unmarked)

	// Part 2
	boards = NewBoards(cells, 5)
	lib.AssertEq(t, "len(boards)", 3, len(boards))

	lastWinner, last := PlayGameToWinLast(boards, moves)
	lib.AssertEq(t, "losing move", 13, last)
	lib.AssertEq(t, "first square of last-winning board 2", 3, lastWinner.cells[0])

	unmarked = SumUnmarked(lastWinner)
	DisplayBoard(lastWinner)
	lib.AssertEq(t, "sum unmarked", 148, unmarked)
}
