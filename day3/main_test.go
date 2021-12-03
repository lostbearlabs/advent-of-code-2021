package main

import (
	"fmt"
	"strings"
	"testing"
)

func assertEq(t *testing.T, name string, expected int, actual int) {
	if expected != actual {
		t.Errorf("%s: expected %d, got %d", name, expected, actual)
	}
}

func TestSimple(t *testing.T) {
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`

	rd := strings.NewReader(input)
	nums, bits := LinesToNumbers(rd)
	gamma1, epsilon1 := CalcGammaEpsilon(nums, bits)
	assertEq(t, "gamma1", 22, gamma1)
	assertEq(t, "epsilon1", 9, epsilon1)

	o2, co2 := ComputeO2AndCO2(nums, bits)
	fmt.Printf("o2,co2=%d,%d\n", o2, co2)
	assertEq(t, "o2", 23, o2)
	assertEq(t, "co2", 10, co2)

}
