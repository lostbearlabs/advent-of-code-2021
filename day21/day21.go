package main

import (
	"fmt"
	"log"
)

// Usage:
//   The input is short, so the sample input and my answers are encoded in the unit tests

// https://adventofcode.com/2021/day/21

func main() {

}

type Die struct {
	value int
}

func NewDie() Die {
	return Die{value: 0}
}

func (d *Die) roll() int {
	val := d.value + 1
	d.value = (d.value + 1) % 100
	return val
}

type Player struct {
	pos   int
	score int
}

func NewPlayer(pos int) Player {
	return Player{pos: pos}
}

func (p Player) PlayTurn(x, y, z int) Player {
	newPos := (p.pos-1+x+y+z)%10 + 1
	return Player{
		pos:   newPos,
		score: p.score + newPos}
}

func PlayGame(start1 int, start2 int, maxScore int) (losingPlayerScore int, numDieRolls int) {
	d := NewDie()
	p := NewPlayer(start1)
	q := NewPlayer(start2)
	rolls := 0

	for {
		p = p.PlayTurn(d.roll(), d.roll(), d.roll())
		rolls += 3

		if p.score >= maxScore {
			return q.score, rolls
		}

		p, q = q, p
	}

}

type Universe struct {
	p      Player // represented as separate variables, not an array, for ease of use as a map key
	q      Player
	winner int // 0 for nobody yet, 1, or 2
}

func Clone(u Universe) Universe {
	return Universe{
		p:      u.p,
		q:      u.q,
		winner: u.winner,
	}
}

// universe -> number of times we find ourselves in that universe
type Universes struct {
	mp map[Universe]int
}

func (uu Universes) countWinners() (win1 int, win2 int, pending int) {
	ar := []int{0, 0, 0}
	for k, v := range uu.mp {
		ar[k.winner] += v
	}
	return ar[1], ar[2], ar[0]
}

func (u Universe) PlayTurn(dieValue int, player int, maxScore int) Universe {
	z := Clone(u)
	if player == 1 {
		z.p = z.p.PlayTurn(dieValue, 0, 0)
		if z.p.score >= maxScore {
			z.winner = 1
		}
	} else {
		z.q = z.q.PlayTurn(dieValue, 0, 0)
		if z.q.score >= maxScore {
			z.winner = 2
		}
	}
	return z
}

func (uu Universes) PrintSummary() {
	wins := []int{0, 0, 0}
	positions := make(map[int]int)
	scores := make(map[int]int)
	maxScore := 0
	minScore := 0

	for p, _ := range uu.mp {
		wins[p.winner]++
		positions[p.p.pos]++
		positions[p.q.pos]++
		scores[p.p.score]++
		scores[p.q.score]++

		if maxScore == 0 || p.p.score > maxScore {
			maxScore = p.p.score
		}
		if minScore == 0 || p.p.score < minScore {
			minScore = p.p.score
		}
		if maxScore == 0 || p.q.score > maxScore {
			maxScore = p.q.score
		}
		if minScore == 0 || p.q.score < minScore {
			minScore = p.q.score
		}
	}

	fmt.Println("size: ", len(uu.mp))
	fmt.Println("wins: ", wins)
	fmt.Println("num pos: ", len(positions))
	fmt.Println("num score: ", len(scores))
	fmt.Printf("min/max score: %d/%d\n", minScore, maxScore)
	fmt.Println()
}

func (uu Universes) PlayTurn(player int, maxScore int) Universes {
	zz := Universes{mp: map[Universe]int{}}
	for u, n := range uu.mp {
		if u.winner > 0 {
			// universe already has a winner, no turn is played
			zz.mp[u] += n
		} else {
			// player rolls 3 times and advances by the combined value
			for i := 1; i <= 3; i++ {
				for j := 1; j <= 3; j++ {
					for k := 1; k <= 3; k++ {
						zz.mp[u.PlayTurn(i+j+k, player, maxScore)] += n
					}
				}
			}
		}
	}
	log.Printf("PlayTurn: %d -> %d\n", len(uu.mp), len(zz.mp))
	if len(zz.mp) < 30 {
		log.Println(zz)
	}
	zz.PrintSummary()
	return zz
}

func PlayQuantumGame(start1 int, start2 int, maxScore int) (wins1 int, wins2 int) {
	u := Universe{
		p:      NewPlayer(start1),
		q:      NewPlayer(start2),
		winner: 0,
	}

	uu := Universes{mp: map[Universe]int{u: 1}}

	for {
		uu = uu.PlayTurn(1, maxScore)
		uu = uu.PlayTurn(2, maxScore)

		win1, win2, pending := uu.countWinners()
		if pending == 0 {
			return win1, win2
		}
	}

}
