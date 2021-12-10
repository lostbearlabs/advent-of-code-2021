package main

import (
	"aoc21/lib"
	"strings"
	"testing"
)

func TestSampleScore(t *testing.T) {
	text := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

	rd := strings.NewReader(text)
	lines := lib.ReadLines(rd)
	errorScore, _ := ScoreLines(lines)
	lib.AssertEq(t, "errorScore", 26397, errorScore)
}

func TestSampleDelim1(t *testing.T) {
	sample := `{([(<{}[<>[]}>{[]{[(<()>`
	sampleExpected, sampleFound := "]", "}"

	expected, found, _ := ParseLine(sample)
	lib.AssertEqStr(t, "expected", sampleExpected, expected)
	lib.AssertEqStr(t, "found", sampleFound, found)
}

func TestValid(t *testing.T) {
	sample := `{([<>])}`

	expected, found, _ := ParseLine(sample)
	lib.AssertEqStr(t, "expected", "", expected)
	lib.AssertEqStr(t, "found", "", found)
}

func TestCompletion(t *testing.T) {
	sample := `<{([{{}}[<[[[<>{}]]]>[]]`

	_, _, completion := ParseLine(sample)
	lib.AssertEqStr(t, "completion", "])}>", completion)
}

func TestScoreCompletion(t *testing.T) {
	sample := `<{([{{}}[<[[[<>{}]]]>[]]`

	_, completionScore := ScoreLine(sample)
	lib.AssertEq(t, "completionScore", 294, completionScore)
}
