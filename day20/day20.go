package main

import (
	"aoc21/lib"
	"fmt"
	"log"
	"os"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/20

func main() {
	lines := lib.ReadLines(os.Stdin)
	num := 50
	problem := ParseProblem(lines, num+5)

	image := problem.image
	for i := 0; i < num; i++ {
		image = EnhanceImage(image, problem.algorithm)

		if i == 0 || i == 1 || i == 49 {
			log.Println("count image ", i+1, " ==> ", CountPixels(image))
		}
	}

}

func StringToBits(st string) []byte {
	ar := make([]byte, len(st))
	for i, c := range st {
		if c != '.' {
			ar[i] = 1
		}
	}
	return ar
}

// *************
// Algorithm
// *************

type Algorithm struct {
	mapping []byte
}

func NewAlgorithm(mapping []byte) Algorithm {
	if len(mapping) != 512 {
		log.Fatal("bad algorithm len ", len(mapping))
	}
	return Algorithm{mapping: mapping}
}

func ParseAlgorithm(line string) Algorithm {
	ar := StringToBits(line)
	return NewAlgorithm(ar)
}

// *************
// Image
// *************

type Image struct {
	grid [][]byte
}

func ParseImage(lines []string) Image {
	var grid [][]byte
	for _, st := range lines {
		if st == "" {
			continue
		}
		grid = append(grid, StringToBits(st))
	}
	return Image{grid: grid}
}

func PadImage(image Image, pad int) Image {

	var grid [][]byte
	for i := 0; i < pad; i++ {
		grid = append(grid, make([]byte, len(image.grid[0])+2*pad))
	}

	for _, row := range image.grid {
		ar := make([]byte, len(row)+2*pad)
		copy(ar[pad:], row)
		grid = append(grid, ar)
	}

	for i := 0; i < pad; i++ {
		grid = append(grid, make([]byte, len(image.grid[0])+2*pad))
	}

	return Image{grid: grid}
}

func PrintImage(image Image) {
	fmt.Printf("\n")
	for _, row := range image.grid {
		for _, b := range row {
			if b > 0 {
				fmt.Printf("*")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func CountPixels(image Image) int {
	n := 0
	for i := 0; i < len(image.grid); i++ {
		for j := 0; j < len(image.grid[i]); j++ {
			if image.grid[i][j] > 0 {
				n++
			}
		}
	}
	return n
}

func Clone(image Image) Image {
	return PadImage(image, 0)
}

// *************
// Problem
// *************

type Problem struct {
	algorithm Algorithm
	image     Image
}

func ParseProblem(lines []string, pad int) Problem {
	return Problem{
		algorithm: ParseAlgorithm(lines[0]),
		image:     PadImage(ParseImage(lines[1:]), pad),
	}
}

// *************
// Enhance
// *************

func lookup(i int, j int, image Image, algorithm Algorithm) int {

	// special case for the infinite beyond: cells outside the padding are either
	// all going to stay dark or all swap between light and dark
	if i < 1 || j < 1 || i >= len(image.grid)-1 || j >= len(image.grid[i])-1 {
		i = 1
		j = 1
	}

	idx := 0
	idx = idx<<1 + int(image.grid[i-1][j-1])
	idx = idx<<1 + int(image.grid[i-1][j])
	idx = idx<<1 + int(image.grid[i-1][j+1])
	idx = idx<<1 + int(image.grid[i][j-1])
	idx = idx<<1 + int(image.grid[i][j])
	idx = idx<<1 + int(image.grid[i][j+1])
	idx = idx<<1 + int(image.grid[i+1][j-1])
	idx = idx<<1 + int(image.grid[i+1][j])
	idx = idx<<1 + int(image.grid[i+1][j+1])

	return int(algorithm.mapping[idx])
}

func EnhanceImage(image Image, algorithm Algorithm) Image {
	next := Clone(image)

	for i := 0; i < len(image.grid); i++ {
		for j := 0; j < len(image.grid[i]); j++ {
			next.grid[i][j] = byte(lookup(i, j, image, algorithm))
		}
	}

	return next
}
