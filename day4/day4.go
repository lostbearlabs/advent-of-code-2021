package main

import (
	"aoc21/lib"
	"fmt"
	"os"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/4

func main() {
	lines := lib.ReadLines(os.Stdin)
	moves := lib.LinesToNumbersSep(lines[0:1], 10, ",")
	fmt.Printf("read %d moves, first=%d, last=%d\n", len(moves), moves[0], moves[len(moves)-1])

	// Part 1
	cells := lib.LinesToNumbersSep(lines[2:], 10, " ")
	boards := NewBoards(cells, 5)
	fmt.Printf("read %d boards\n", len(boards))

	winner, last := PlayGame(boards, moves)
	unmarked := SumUnmarked(winner)

	DisplayBoard(winner)
	fmt.Printf("First winner: last=%d, sum unmarked=%d, product=%d\n", last, unmarked, last*unmarked)

	// Part 2
	// Done in a separate pass over the boards;  could be adjusted so both results are done in a single sweep
	fmt.Println()
	boards = NewBoards(cells, 5)

	lastWinner, last := PlayGameToWinLast(boards, moves)
	unmarked = SumUnmarked(lastWinner)

	DisplayBoard(lastWinner)
	fmt.Printf("Last winner: last=%d, sum unmarked=%d, product=%d\n", last, unmarked, last*unmarked)
}
