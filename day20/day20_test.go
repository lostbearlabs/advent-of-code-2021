package main

import (
	"aoc21/lib"
	"strings"
	"testing"
)

func sample() string {
	return `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`
}

func TestParseProblem(t *testing.T) {
	lines := lib.ReadLines(strings.NewReader(sample()))
	problem := ParseProblem(lines, 3)
	PrintImage(problem.image)
	lib.AssertEq(t, "", 11, len(problem.image.grid))
	lib.AssertEq(t, "", 11, len(problem.image.grid[0]))
	lib.AssertEq(t, "", 10, CountPixels(problem.image))
}

func TestEnhanceImage(t *testing.T) {
	lines := lib.ReadLines(strings.NewReader(sample()))
	problem := ParseProblem(lines, 5)
	PrintImage(problem.image)

	image1 := EnhanceImage(problem.image, problem.algorithm)
	PrintImage(image1)

	image2 := EnhanceImage(image1, problem.algorithm)
	PrintImage(image2)

	lib.AssertEq(t, "image2 len ", 35, CountPixels(image2))
}

func TestEnhanceImage50(t *testing.T) {
	lines := lib.ReadLines(strings.NewReader(sample()))
	problem := ParseProblem(lines, 55)

	image := problem.image
	for i := 0; i < 50; i++ {
		image = EnhanceImage(image, problem.algorithm)
	}

	lib.AssertEq(t, "image50 len ", 3351, CountPixels(image))
}
