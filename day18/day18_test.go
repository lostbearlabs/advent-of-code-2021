package main

import (
	"aoc21/lib"
	"strings"
	"testing"
)

func testParseRoundtrip(t *testing.T, st string) {
	num := Parse(st)
	disp := Render(num)
	lib.AssertEqStr(t, "", st, disp)
}

func TestParse(t *testing.T) {
	testParseRoundtrip(t, "[1,2]")
	testParseRoundtrip(t, "[[1,2],3]")
	testParseRoundtrip(t, "[9,[8,7]]")
	testParseRoundtrip(t, "[[1,9],[8,5]]")
	testParseRoundtrip(t, "[[[[1,2],[3,4]],[[5,6],[7,8]]],9]")
	testParseRoundtrip(t, "[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]")
	testParseRoundtrip(t, "[1,2]")
	testParseRoundtrip(t, "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]")
}

func TestMagnitude(t *testing.T) {
	lib.AssertEq(t, "ex1", 143, Magnitude(Parse("[[1,2],[[3,4],5]]")))
	lib.AssertEq(t, "ex2", 1384, Magnitude(Parse("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")))
	lib.AssertEq(t, "ex3", 445, Magnitude(Parse("[[[[1,1],[2,2]],[3,3]],[4,4]]")))
	lib.AssertEq(t, "ex4", 791, Magnitude(Parse("[[[[3,0],[5,3]],[4,4]],[5,5]]")))
	lib.AssertEq(t, "ex5", 1137, Magnitude(Parse("[[[[5,0],[7,4]],[5,5]],[6,6]]")))
	lib.AssertEq(t, "ex6", 3488, Magnitude(Parse("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")))
}

func TestAddList(t *testing.T) {
	a := Parse("[0,1]")
	b := Parse("[2,3]")
	c := Parse("[4,5]")
	tot := AddList([]*SnailfishNumber{a, b, c})
	lib.AssertEqStr(t, "", "[[[0,1],[2,3]],[4,5]]", Render(tot))
}

func testExplodeRoundTrip(t *testing.T, in string, out string) {
	num := Parse(in)
	_ = Explode(num)
	rendered := Render(num)
	lib.AssertEqStr(t, "", out, rendered)
}

func TestExplode(t *testing.T) {
	testExplodeRoundTrip(t, "[[1,2],[[3,4],5]]", "[[1,2],[[3,4],5]]") // no change
	testExplodeRoundTrip(t, "[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]")
	testExplodeRoundTrip(t, "[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]")
	testExplodeRoundTrip(t, "[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]")
	testExplodeRoundTrip(t, "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]")
	testExplodeRoundTrip(t, "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]")
}

func testSplitRoundTrip(t *testing.T, in int, out string) {
	num := &SnailfishNumber{
		leaf:  true,
		x:     in,
		left:  nil,
		right: nil,
	}
	_ = Split(num)
	rendered := Render(num)
	lib.AssertEqStr(t, "", out, rendered)
}

func TestSplit(t *testing.T) {
	testSplitRoundTrip(t, 6, "6") // no change
	testSplitRoundTrip(t, 10, "[5,5]")
	testSplitRoundTrip(t, 11, "[5,6]")
}

func TestSample(t *testing.T) {
	sample := `[[[[4,3],4],4],[7,[[8,4],9]]]
[1,1]`
	rd := strings.NewReader(sample)
	lines := lib.ReadLines(rd)

	ar := make([]*SnailfishNumber, len(lines))
	for i, line := range lines {
		ar[i] = Parse(line)
	}

	num := AddList(ar)

	lib.AssertEqStr(t, "", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", Render(num))
}

func TestSample2(t *testing.T) {
	sample := `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`
	rd := strings.NewReader(sample)
	lines := lib.ReadLines(rd)

	ar := make([]*SnailfishNumber, len(lines))
	for i, line := range lines {
		ar[i] = Parse(line)
	}

	num := AddList(ar)

	lib.AssertEqStr(t, "", "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]", Render(num))
	lib.AssertEq(t, "", 4140, Magnitude(num))

	lib.AssertEq(t, "", 3993, LargestPairwiseMagnitude(ar))
}
