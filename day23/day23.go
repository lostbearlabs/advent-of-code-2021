package main

import (
	"fmt"
	"log"
)

// Usage:
//   go run .

// https://adventofcode.com/2021/day/11

// Part 1 input
// #############
// #...........#
// ###D#C#D#B###
//   #C#A#A#B#
//   #########

// Part 2 input
// #############
// #...........#
// ###D#C#D#B###
//   #D#C#B#A#
//   #D#B#A#C#
//   #C#A#A#B#
//   #########

const Infinity = 1_000_000_000_000

func main() {
	state := NewState(14, 16, 17, 18, 13, 12, 11, 15)
	cost := BestCost(state, "", map[string]int{}, 20)
	log.Println("best cost: ", cost)

	state2 := NewState2(D, D, D, C, C, C, B, A, D, B, A, A, B, A, C, B)
	cost2 := BestCost(state2, "", map[string]int{}, 45)
	log.Println("best cost 2: ", cost2)
}

// ******************
// State
// ******************

// 0  1  2  3  4  5  6  7  8  9  10
//       11    13    15    17
//       12    14    16    18

const (
	EMPTY = iota
	A
	B
	C
	D
)

type State struct {
	n        int
	roomSize int
	ar       []int
}

func NewState(a1, a2, b1, b2, c1, c2, d1, d2 int) State {
	roomSize := 2
	n := 11 + 4*roomSize
	st := State{
		n:        n,
		roomSize: roomSize,
		ar:       make([]int, n),
	}
	st.ar[a1] = A
	st.ar[a2] = A
	st.ar[b1] = B
	st.ar[b2] = B
	st.ar[c1] = C
	st.ar[c2] = C
	st.ar[d1] = D
	st.ar[d2] = D

	return st
}

func NewState2(start ...int) State {
	roomSize := 4
	if len(start) != 4*roomSize {
		log.Fatalf("bad start size")
	}

	n := 11 + 4*roomSize
	st := State{
		n:        n,
		roomSize: roomSize,
		ar:       make([]int, n),
	}

	copy(st.ar[11:], start)

	return st
}

func (state State) IsFinal() bool {
	for i := 11; i < state.n; i++ {
		expected := A + (i-11)/state.roomSize
		if state.ar[i] != expected {
			return false
		}
	}
	return true
}

func (state State) Clone() State {
	t := State{
		n:        state.n,
		roomSize: state.roomSize,
	}
	t.ar = make([]int, state.n)
	copy(t.ar, state.ar)
	return t
}

func (state State) NearestHallwayPos(a int) int {
	if a <= 10 {
		return a
	}

	room := (a - 11) / state.roomSize
	return (room + 1) * 2 // 2,4,6,8
}

func distanceFromHallway(state State, a int) (int, bool) {

	if a <= 10 {
		return 0, false
	}

	room := (a - 11) / state.roomSize
	doorPos := 11 + room*state.roomSize
	distance := a - doorPos + 1

	// blocked if the hallway outside the door is blocked
	p := state.NearestHallwayPos(a)
	blocked := state.ar[p] != EMPTY

	// blocked if one of the room spaces between the hallway and the
	// target space is blocked
	for i := doorPos; i < a; i++ {
		if state.ar[i] != EMPTY {
			blocked = true
		}
	}

	return distance, blocked
}

func distanceAlongHallway(state State, a int, b int) (int, bool) {
	if a == b {
		return 0, false
	}

	p := state.NearestHallwayPos(a)
	q := state.NearestHallwayPos(b)
	distance := 0
	delta := 1
	blocked := false

	if p > q {
		delta = -1
	}
	for x := p + delta; x != q+delta; x += delta {
		distance++
		if state.ar[x] != EMPTY {
			blocked = true
		}
	}

	return distance, blocked
}

func CanMove(state State, a int, b int) (cost int, ok bool) {
	var roomPos = []int{0, 11, 11 + state.roomSize, 11 + 2*state.roomSize, 11 + 3*state.roomSize}

	mover := state.ar[a]
	moverEndDoor := roomPos[mover]
	moverEndWindow := moverEndDoor + state.roomSize - 1

	// have to actually be moving
	if a == b {
		return 0, false
	}

	// have to have something to move
	if mover == EMPTY {
		return 0, false
	}

	// have to have empty space to move to
	if state.ar[b] != EMPTY {
		return 0, false
	}

	// can't move within the hallway
	if a <= 10 && b <= 10 {
		return 0, false
	}

	// can't stop on space outside door
	if b == 2 || b == 4 || b == 6 || b == 8 {
		return 0, false
	}

	if b > 10 {
		// can only move into appropriate room
		if b < moverEndDoor || b > moverEndWindow {
			return 0, false
		}

		// don't move into pos if blocked
		for i := moverEndDoor; i < b; i++ {
			if state.ar[i] != EMPTY {
				return 0, false
			}
		}

		// don't move into near pos when far pos is not already occupied by target type
		for i := b + 1; i <= moverEndWindow; i++ {
			if state.ar[i] != mover {
				return 0, false
			}
		}
	}

	// if already in the right room, only move if there's something further
	// back in the room that still needs to move
	if a >= moverEndDoor && a <= moverEndWindow {
		allAreMover := true
		for i := a + 1; i <= moverEndWindow; i++ {
			if state.ar[i] != mover {
				allAreMover = false
			}
		}
		if allAreMover {
			return 0, false
		}
	}

	d1, b1 := distanceFromHallway(state, a)
	d2, b2 := distanceAlongHallway(state, a, b)
	d3, b3 := distanceFromHallway(state, b)

	if b1 || b2 || b3 {
		return 0, false
	}

	multipliers := []int{0, 1, 10, 100, 1000}
	return multipliers[mover-EMPTY] * (d1 + d2 + d3), true
}

func NextStates(state State) ([]State, []int) {
	var ar []State
	var br []int
	for i := 0; i < state.n; i++ {
		for j := 0; j < state.n; j++ {
			cost, ok := CanMove(state, i, j)
			if ok {
				t := state.Clone()
				t.ar[i] = EMPTY
				t.ar[j] = state.ar[i]
				ar = append(ar, t)
				br = append(br, cost)
			}
		}
	}
	return ar, br
}

// ******************
// Search
// ******************

func ToKey(state State) string {
	return fmt.Sprintf("%v", state.ar)
}

func BestCost(state State, st string, memo map[string]int, maxDepth int) int {
	key := ToKey(state)
	cost, ok := memo[key]
	if ok {
		return cost
	}

	if maxDepth < 1 {
		return Infinity
	}

	if state.IsFinal() {
		return 0
	}

	next, costs := NextStates(state)
	if len(next) == 0 {
		return Infinity
	}

	min := 0
	for i, s := range next {
		cost := BestCost(s, st+" ", memo, maxDepth-1)
		//log.Printf("%s%d + %d = %d -> %v\n", st, cost, costs[i], cost+costs[i], s.ar)
		if i == 0 || cost+costs[i] < min {
			min = cost + costs[i]
		}
	}
	memo[key] = min
	return min
}
