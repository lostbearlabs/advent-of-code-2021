package main

import (
	"aoc21/lib"
	"testing"
)

func assertFalse(t *testing.T, b bool, msg string) {
	if b {
		t.Errorf("unexpected True, %s", msg)
	}
}

func assertTrue(t *testing.T, b bool, msg string) {
	if !b {
		t.Errorf("unexpected False, %s", msg)
	}
}

func TestRowWins(t *testing.T) {
	board := NewBoard([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)

	PlayNumber(board, 1)
	assertFalse(t, IsWinner(board), "played 1")

	PlayNumber(board, 2)
	assertFalse(t, IsWinner(board), "played 2")

	PlayNumber(board, 3)
	assertTrue(t, IsWinner(board), "played 3")
}

func TestColWins(t *testing.T) {
	board := NewBoard([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)

	PlayNumber(board, 3)
	assertFalse(t, IsWinner(board), "played 3")

	PlayNumber(board, 6)
	assertFalse(t, IsWinner(board), "played 6")

	PlayNumber(board, 9)
	assertTrue(t, IsWinner(board), "played 9")
}

func TestDiagFails(t *testing.T) {
	board := NewBoard([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3)

	PlayNumber(board, 1)
	assertFalse(t, IsWinner(board), "played 1")

	PlayNumber(board, 5)
	assertFalse(t, IsWinner(board), "played 5")

	PlayNumber(board, 9)
	assertFalse(t, IsWinner(board), "played 9")
}

func TestMakeBoards(t *testing.T) {
	boards := NewBoards([]int{1, 2, 3, 4, 5, 6, 7, 8}, 2)
	lib.AssertEq(t, "numBoards", len(boards), 2)
	lib.AssertEq(t, "boardSquared", len(boards[1].cells), 4)
}
