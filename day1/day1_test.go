package main

import (
	"strings"
	"testing"
)

func assertEq(t *testing.T, name string, expected int, actual int) {
	if expected != actual {
		t.Errorf("%s: expected %d, got %d", name, expected, actual)
	}
}

func TestSimple(t *testing.T) {
	input := `1
2
3
1
5
`

	rd := strings.NewReader(input)
	numIncrease, numWindowIncrease := Count(rd)
	assertEq(t, "numIncrease", 3, numIncrease)
	assertEq(t, "numWindowIncrease", 1, numWindowIncrease)
}
