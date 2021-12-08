package main

import (
	"aoc21/lib"
	"strings"
	"testing"
)

func sample1() string {
	return `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`
}

func sample2() string {
	return `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`
}

func TestExampleEasyDigits(t *testing.T) {
	reader := strings.NewReader(sample2())
	lines := lib.ReadLines(reader)
	entries := ParseEntries(lines)

	easyCount := EasyCount(entries)
	lib.AssertEq(t, "easyCount", 26, easyCount)
}

func TestAsInt(t *testing.T) {
	lib.AssertEq(t, "a", 1, asInt("a"))
	lib.AssertEq(t, "a", 3, asInt("ab"))
}

func TestAsStr(t *testing.T) {
	lib.AssertEqStr(t, "afg", "afg", asStr(0b1100001))
}

func TestDiff(t *testing.T) {
	lib.AssertEqStr(t, "", "df", diff("abcdefg", "abceg"))
}

func TestUnion(t *testing.T) {
	lib.AssertEqStr(t, "", "adeg", union("ag", "d", "ae"))
}

func TestNot(t *testing.T) {
	lib.AssertEqStr(t, "", "acd", not("befg"))
}

func TestCanonical(t *testing.T) {
	lib.AssertEqStr(t, "", "abcdefg", canonicalString("gdfegacdcab"))
}

func TestIntersect(t *testing.T) {
	lib.AssertEqStr(t, "", "cd", intersect("cdefg", "bcd"))
}

func TestDeduceDefaultMapping(t *testing.T) {
	line := `abcefg cf acdeg acdfg bcdf abdfg abdefg acf abcdefg abcdfg | cf cf cf cf`

	entry := ParseEntry(line)
	mp := DeduceMapping(entry)

	lib.AssertEqMap(t, map[string]int{
		canonicalString("abcefg"):  0,
		canonicalString("cf"):      1,
		canonicalString("acdeg"):   2,
		canonicalString("acdfg"):   3,
		canonicalString("bcdf"):    4,
		canonicalString("abdfg"):   5,
		canonicalString("abdefg"):  6,
		canonicalString("acf"):     7,
		canonicalString("abcdefg"): 8,
		canonicalString("abcdfg"):  9,
	}, mp)

}

func TestDeduceMapping1(t *testing.T) {
	reader := strings.NewReader(sample1())
	lines := lib.ReadLines(reader)
	entry := ParseEntry(lines[0])
	mp := DeduceMapping(entry)

	lib.AssertEqMap(t, map[string]int{
		canonicalString("acedgfb"): 8,
		canonicalString("cdfbe"):   5,
		canonicalString("gcdfa"):   2,
		canonicalString("fbcad"):   3,
		canonicalString("dab"):     7,
		canonicalString("cefabd"):  9,
		canonicalString("cdfgeb"):  6,
		canonicalString("eafb"):    4,
		canonicalString("cagedb"):  0,
		canonicalString("ab"):      1,
	}, mp)
}

func TestGetTotalOutput1(t *testing.T) {
	reader := strings.NewReader(sample1())
	lines := lib.ReadLines(reader)
	entries := ParseEntries(lines)
	n := GetTotalOutput(entries)
	lib.AssertEq(t, "n", 5353, n)
}

func TestGetTotalOutput2(t *testing.T) {
	reader := strings.NewReader(sample2())
	lines := lib.ReadLines(reader)
	entries := ParseEntries(lines)
	n := GetTotalOutput(entries)
	lib.AssertEq(t, "n", 61229, n)
}
