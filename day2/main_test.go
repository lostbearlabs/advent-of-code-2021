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
	input := `forward 5
down 5
forward 8
up 3
down 8
forward 2
`

	rd := strings.NewReader(input)
	horz1, depth1, horz2, depth2 := Count(rd)
	assertEq(t, "horz1", 15, horz1)
	assertEq(t, "depth1", 10, depth1)
	assertEq(t, "horz2", 15, horz2)
	assertEq(t, "depth2", 60, depth2)

}
