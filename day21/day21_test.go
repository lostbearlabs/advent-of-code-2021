package main

import (
	"aoc21/lib"
	"testing"
)

func TestRollDie(t *testing.T) {
	d := NewDie()
	for i := 0; i < 1000; i++ {
		val := d.roll()
		lib.AssertEq(t, "", i%100+1, val)
	}
}

func TestMove(t *testing.T) {
	p := NewPlayer(4)

	p = p.PlayTurn(1, 2, 3)
	lib.AssertEq(t, "", 10, p.score)

	p = p.PlayTurn(7, 8, 9)
	lib.AssertEq(t, "", 14, p.score)

	p = p.PlayTurn(13, 14, 15)
	lib.AssertEq(t, "", 20, p.score)

	p = p.PlayTurn(19, 20, 21)
	lib.AssertEq(t, "", 26, p.score)
}

func TestPlayGame(t *testing.T) {
	losingScore, numRolls := PlayGame(4, 8, 1000)
	lib.AssertEq(t, "", 745, losingScore)
	lib.AssertEq(t, "", 993, numRolls)
}

func TestPlayGame2(t *testing.T) {
	losingScore, numRolls := PlayGame(10, 8, 1000)
	lib.AssertEq(t, "", 752247, losingScore*numRolls)
}

func TestPlayQuantumGame(t *testing.T) {
	w1, w2 := PlayQuantumGame(4, 8, 21)
	lib.AssertEq(t, "", 444356092776315, w1)
	lib.AssertEq(t, "", 341960390180808, w2)
}

func TestPlayQuantumGame2(t *testing.T) {
	w1, w2 := PlayQuantumGame(10, 8, 21)
	lib.AssertEq(t, "", 221109915584112, w1)
	lib.AssertEq(t, "", 117096403483545, w2)
}
