package main

import (
	"aoc21/lib"
	"strings"
	"testing"
)

func sample() string {
	return `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`
}

func TestSampleBruteForce(t *testing.T) {
	rd := strings.NewReader(sample())
	fnord := ReadFnord(rd)

	st := UpdatePolymer(fnord.Polymer, fnord.Rules)
	lib.AssertEqStr(t, "step 1", "NCNBCHB", st)

	st = UpdatePolymer(st, fnord.Rules)
	lib.AssertEqStr(t, "step 2", "NBCCNBBBCBHCB", st)

	st = UpdatePolymer(st, fnord.Rules)
	lib.AssertEqStr(t, "step 3", "NBBBCNCCNBBNBNBBCHBHHBCHB", st)

	st = UpdatePolymer(st, fnord.Rules)
	lib.AssertEqStr(t, "step 4", "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", st)

	for i := 5; i <= 10; i++ {
		st = UpdatePolymer(st, fnord.Rules)
	}

	ar := GetCharCountsFromString(st)

	lib.AssertEqStr(t, "least common name", "H", ar[0].Name)
	lib.AssertEq(t, "least common count", 161, ar[0].Count)

	lib.AssertEqStr(t, "most common name", "B", ar[len(ar)-1].Name)
	lib.AssertEq(t, "most common count", 1749, ar[len(ar)-1].Count)
}

func TestGetCharCounts(t *testing.T) {
	st := "ABCABA"
	ar := GetCharCountsFromString(st)

	lib.AssertEqStr(t, "", ar[0].Name, "C")
	lib.AssertEq(t, "", ar[0].Count, 1)

	lib.AssertEqStr(t, "", ar[1].Name, "B")
	lib.AssertEq(t, "", ar[1].Count, 2)

	lib.AssertEqStr(t, "", ar[2].Name, "A")
	lib.AssertEq(t, "", ar[2].Count, 3)
}

func TestSampleWithPairMap(t *testing.T) {
	rd := strings.NewReader(sample())
	fnord := ReadFnord(rd)

	mp := ToPairMap(fnord.Polymer)
	for i := 0; i < 10; i++ {
		mp = UpdatePairMap(mp, fnord.Rules)
	}

	ar := GetCharCountsFromPairMap(mp, fnord.Polymer)

	lib.AssertEqStr(t, "least common name", "H", ar[0].Name)
	lib.AssertEq(t, "least common count", 161, ar[0].Count)

	lib.AssertEqStr(t, "most common name", "B", ar[len(ar)-1].Name)
	lib.AssertEq(t, "most common count", 1749, ar[len(ar)-1].Count)
}
