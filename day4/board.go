package main

import "fmt"

type Board struct {
	size   int
	cells  []int
	marked []bool
}

func NewBoard(input []int, size int) *Board {
	if len(input) != size*size {
		fmt.Errorf("bad data, len(input)=%d, size=%d\n", len(input), size)
	}

	return &Board{
		size:   size,
		cells:  input,
		marked: make([]bool, size*size),
	}
}

func NewBoards(input []int, size int) []*Board {
	var boards []*Board

	if len(input)%(size*size) != 0 {
		fmt.Errorf("bad data, len(input)=%d, size=%d\n", len(input), size)
	}

	for i := 0; i < len(input); i += size * size {
		board := NewBoard(input[i:i+size*size], size)
		boards = append(boards, board)
	}

	return boards
}

func PlayNumber(board *Board, num int) {
	for i, n := range board.cells {
		if n == num {
			board.marked[i] = true
		}
	}
}

func IsWinner(board *Board) bool {
	for i := 0; i < board.size; i++ {
		row := true
		col := true
		for j := 0; j < board.size; j++ {
			if !board.marked[i*board.size+j] {
				row = false
			}
			if !board.marked[i+j*board.size] {
				col = false
			}
		}
		if row || col {
			return true
		}
	}
	return false
}

func PlayGame(boards []*Board, moves []int) (*Board, int) {
	for _, move := range moves {
		for _, board := range boards {
			PlayNumber(board, move)
			if IsWinner(board) {
				return board, move
			}
		}
	}

	return nil, 0
}

func SumUnmarked(board *Board) int {
	sum := 0
	fmt.Println(board.cells)
	fmt.Println(board.marked)
	for i, n := range board.cells {
		if !board.marked[i] {
			sum += n
		}
	}
	return sum
}

func DisplayBoard(board *Board) {
	for i := 0; i < board.size; i++ {
		for j := 0; j < board.size; j++ {
			n := i*board.size + j
			if board.marked[n] {
				fmt.Printf(" [%2d]", board.cells[n])
			} else {
				fmt.Printf("  %2d ", board.cells[n])
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func PlayGameToWinLast(boards []*Board, moves []int) (*Board, int) {
	var lastWinner *Board
	var lastMove int

	for _, move := range moves {
		for _, board := range boards {
			if !IsWinner(board) {
				PlayNumber(board, move)
				if IsWinner(board) {
					lastWinner = board
					lastMove = move
				}
			}
		}
	}

	return lastWinner, lastMove
}
