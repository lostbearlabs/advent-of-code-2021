package main

import (
	"aoc21/lib"
	"strings"
	"testing"
)

func TestExample1(t *testing.T) {
	input := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

	rd := strings.NewReader(input)
	lines := lib.ReadLines(rd)
	graph := ParseGraph(lines)

	count := CountPaths(graph)
	lib.AssertEq(t, "count", 10, count)

	count2 := CountPaths2(graph)
	lib.AssertEq(t, "count2", 36, count2)
}

func TestExample2(t *testing.T) {
	input := `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

	rd := strings.NewReader(input)
	lines := lib.ReadLines(rd)
	graph := ParseGraph(lines)
	count := CountPaths(graph)
	lib.AssertEq(t, "count", 19, count)

	count2 := CountPaths2(graph)
	lib.AssertEq(t, "count2", 103, count2)
}

func TestExample3(t *testing.T) {
	input := `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

	rd := strings.NewReader(input)
	lines := lib.ReadLines(rd)
	graph := ParseGraph(lines)
	count := CountPaths(graph)
	lib.AssertEq(t, "count", 226, count)

	count2 := CountPaths2(graph)
	lib.AssertEq(t, "count2", 3509, count2)
}
