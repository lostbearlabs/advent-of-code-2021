package main

import (
	"aoc21/lib"
	"strings"
	"testing"
)

func sampleInput() string {
	return `2199943210
3987894921
9856789892
8767896789
9899965678`
}

func TestSampleInputRisk(t *testing.T) {
	reader := strings.NewReader(sampleInput())
	ar := lib.ReadDigitsArray(reader)
	riskLevel := ComputeRiskLevel(ar)
	lib.AssertEq(t, "riskLevel", 15, riskLevel)
}

func TestCorners(t *testing.T) {
	text := `0110
1111
1111
0110`
	reader := strings.NewReader(text)
	ar := lib.ReadDigitsArray(reader)
	riskLevel := ComputeRiskLevel(ar)
	lib.AssertEq(t, "riskLevel", 4, riskLevel)
}

func TestSampleInputBasins(t *testing.T) {
	reader := strings.NewReader(sampleInput())
	ar := lib.ReadDigitsArray(reader)
	basinProduct := ComputeBasinProduct(ar)
	lib.AssertEq(t, "basinProduct", 1134, basinProduct)
}
