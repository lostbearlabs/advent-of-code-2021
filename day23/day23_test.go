package main

import (
	"aoc21/lib"
	"log"
	"testing"
)

func TestState_NearestHallwayPos(t *testing.T) {
	state := NewState(11, 12, 13, 14, 15, 16, 17, 18)
	lib.AssertEq(t, "", 0, state.NearestHallwayPos(0))
	lib.AssertEq(t, "", 10, state.NearestHallwayPos(10))

	lib.AssertEq(t, "", 2, state.NearestHallwayPos(11))
	lib.AssertEq(t, "", 2, state.NearestHallwayPos(11))

	lib.AssertEq(t, "", 8, state.NearestHallwayPos(17))
	lib.AssertEq(t, "", 8, state.NearestHallwayPos(18))
}

func TestState_IsFinal(t *testing.T) {
	state := NewState(11, 12, 13, 14, 15, 16, 17, 18)
	lib.AssertTrue(t, "", state.IsFinal())

	state = NewState2(A, A, A, A, B, B, B, B, C, C, C, C, D, D, D, D)
	lib.AssertTrue(t, "", state.IsFinal())
}

// 0  1  2  3  4  5  6  7  8  9  10
//       11    13    15    17
//       12    14    16    18

func TestDistanceAlongHallway(t *testing.T) {
	// forward
	state := NewState(0, 0, 0, 0, 0, 0, 0, 0)
	d, b := distanceAlongHallway(state, 2, 8)
	lib.AssertEq(t, "t1", 6, d)
	lib.AssertFalse(t, "t1", b)

	// backward
	state = NewState(0, 0, 0, 0, 0, 0, 0, 0)
	d, b = distanceAlongHallway(state, 8, 2)
	lib.AssertEq(t, "t2", 6, d)
	lib.AssertFalse(t, "t2", b)

	// not blocked by self at start
	state = NewState(8, 0, 0, 0, 0, 0, 0, 0)
	d, b = distanceAlongHallway(state, 8, 2)
	lib.AssertEq(t, "t3", 6, d)
	lib.AssertFalse(t, "t3", b)

	// end occupied
	state = NewState(2, 0, 0, 0, 0, 0, 0, 0)
	d, b = distanceAlongHallway(state, 8, 2)
	lib.AssertEq(t, "t4", 6, d)
	lib.AssertTrue(t, "t4", b)

	// middle occupied
	state = NewState(5, 0, 0, 0, 0, 0, 0, 0)
	d, b = distanceAlongHallway(state, 8, 2)
	lib.AssertEq(t, "t5", 6, d)
	lib.AssertTrue(t, "t5", b)

}

func TestDistanceFromHallway1(t *testing.T) {

	// next to door
	state := NewState(0, 0, 0, 0, 0, 0, 0, 0)
	d, b := distanceFromHallway(state, 13)
	lib.AssertEq(t, "t1", 1, d)
	lib.AssertFalse(t, "t1", b)

	// next to window
	state = NewState(0, 0, 0, 0, 0, 0, 0, 0)
	d, b = distanceFromHallway(state, 14)
	lib.AssertEq(t, "t2", 2, d)
	lib.AssertFalse(t, "t2", b)

	// next to door, hallway occupied
	state = NewState(4, 0, 0, 0, 0, 0, 0, 0)
	d, b = distanceFromHallway(state, 14)
	lib.AssertEq(t, "t3", 2, d)
	lib.AssertTrue(t, "t3", b)

	// next to window, door occupied
	state = NewState(13, 0, 0, 0, 0, 0, 0, 0)
	d, b = distanceFromHallway(state, 14)
	lib.AssertEq(t, "t4", 2, d)
	lib.AssertTrue(t, "t4", b)

	// next to window, occupied by self is not failure
	state = NewState(14, 0, 0, 0, 0, 0, 0, 0)
	d, b = distanceFromHallway(state, 14)
	lib.AssertEq(t, "t5", 2, d)
	lib.AssertFalse(t, "t5", b)
}

func TestCanMove(t *testing.T) {
	// One Bronze amphipod moves into the hallway, taking 4 steps and using 40 energy:
	state := NewState(12, 18, 11, 15, 13, 16, 14, 17)
	cost, ok := CanMove(state, 15, 3)
	lib.AssertTrue(t, "m1", ok)
	lib.AssertEq(t, "m1", 40, cost)

	// The only Copper amphipod not in its side room moves there, taking 4 steps and using 400 energy
	state = NewState(12, 18, 11, 3, 13, 16, 14, 17)
	cost, ok = CanMove(state, 13, 15)
	lib.AssertTrue(t, "m2", ok)
	lib.AssertEq(t, "m2", 400, cost)

	// A Desert amphipod moves out of the way, taking 3 steps and using 3000 energy, and then the Bronze amphipod takes its place, taking 3 steps and using 30 energy:
	state = NewState(12, 18, 11, 3, 15, 16, 14, 17)
	cost, ok = CanMove(state, 14, 5)
	lib.AssertTrue(t, "m3", ok)
	lib.AssertEq(t, "m3", 3000, cost)

	state = NewState(12, 18, 11, 3, 15, 16, 5, 17)
	cost, ok = CanMove(state, 3, 14)
	lib.AssertTrue(t, "m4", ok)
	lib.AssertEq(t, "m4", 30, cost)

	// The leftmost Bronze amphipod moves to its room using 40 energy:
	state = NewState(12, 18, 11, 14, 15, 16, 5, 17)
	cost, ok = CanMove(state, 11, 13)
	lib.AssertTrue(t, "m5", ok)
	lib.AssertEq(t, "m5", 40, cost)

	// Both amphipods in the rightmost room move into the hallway, using 2003 energy in total:
	state = NewState(12, 18, 13, 14, 15, 16, 5, 17)
	cost, ok = CanMove(state, 17, 7)
	lib.AssertTrue(t, "m6", ok)
	lib.AssertEq(t, "m6", 2000, cost)

	state = NewState(12, 18, 13, 14, 15, 16, 5, 7)
	cost, ok = CanMove(state, 18, 9)
	lib.AssertTrue(t, "m7", ok)
	lib.AssertEq(t, "m7", 3, cost)

	// Both Desert amphipods move into the rightmost room using 7000 energy:
	state = NewState(12, 9, 13, 14, 15, 16, 5, 7)
	cost, ok = CanMove(state, 7, 18)
	lib.AssertTrue(t, "m8", ok)
	lib.AssertEq(t, "m8", 3000, cost)

	state = NewState(12, 9, 13, 14, 15, 16, 5, 18)
	cost, ok = CanMove(state, 5, 17)
	lib.AssertTrue(t, "m9", ok)
	lib.AssertEq(t, "m9", 4000, cost)

	state = NewState(12, 9, 13, 14, 15, 16, 17, 18)
	cost, ok = CanMove(state, 9, 11)
	lib.AssertTrue(t, "m10", ok)
	lib.AssertEq(t, "m10", 8, cost)

	state = NewState(12, 11, 13, 14, 15, 16, 17, 18)
	lib.AssertTrue(t, "", state.IsFinal())
}

func TestBestCost(t *testing.T) {
	var state State
	var cost int

	log.Printf("final")
	state = NewState(12, 11, 13, 14, 15, 16, 17, 18)
	cost = BestCost(state, "", map[string]int{}, 10)
	lib.AssertEq(t, "", 0, cost)

	log.Printf("m10")
	state = NewState(12, 9, 13, 14, 15, 16, 17, 18)
	cost = BestCost(state, "m10 ", map[string]int{}, 10)
	lib.AssertEq(t, "", 8, cost)

	log.Printf("m9")
	state = NewState(12, 9, 13, 14, 15, 16, 5, 18)
	cost = BestCost(state, "m9 ", map[string]int{}, 10)
	lib.AssertEq(t, "", 8+4000, cost)

	log.Printf("m8")
	state = NewState(12, 9, 13, 14, 15, 16, 5, 7)
	cost = BestCost(state, "m8 ", map[string]int{}, 10)
	lib.AssertEq(t, "", 8+4000+3000, cost)

	log.Printf("m7")
	state = NewState(12, 18, 13, 14, 15, 16, 5, 7)
	cost = BestCost(state, "m7 ", map[string]int{}, 10)
	lib.AssertEq(t, "", 8+4000+3000+3, cost)

	log.Printf("m6")
	state = NewState(12, 18, 13, 14, 15, 16, 5, 17)
	cost = BestCost(state, "m6 ", map[string]int{}, 10)
	lib.AssertEq(t, "", 8+4000+3000+3+2000, cost)

	log.Printf("m5")
	state = NewState(12, 18, 11, 14, 15, 16, 5, 17)
	cost = BestCost(state, "m5 ", map[string]int{}, 10)
	lib.AssertEq(t, "", 8+4000+3000+3+2000+40, cost)

	log.Printf("m4")
	state = NewState(12, 18, 11, 3, 15, 16, 5, 17)
	cost = BestCost(state, "m4 ", map[string]int{}, 10)
	lib.AssertEq(t, "", 8+4000+3000+3+2000+40+30, cost)

	log.Printf("m3")
	state = NewState(12, 18, 11, 3, 15, 16, 14, 17)
	cost = BestCost(state, "m4 ", map[string]int{}, 10)
	lib.AssertEq(t, "", 8+4000+3000+3+2000+40+30+3000, cost)

	log.Printf("m2")
	state = NewState(12, 18, 11, 3, 13, 16, 14, 17)
	cost = BestCost(state, "m4 ", map[string]int{}, 12)
	lib.AssertEq(t, "", 8+4000+3000+3+2000+40+30+3000+400, cost)

}

func tryMove(t *testing.T, state State, a int, b int, name string, mover int) {
	if state.ar[a] != mover {
		t.Fatalf("moved wrong piece %s", name)
	}
	_, ok := CanMove(state, a, b)
	lib.AssertTrue(t, name, ok)

	state.ar[b] = state.ar[a]
	state.ar[a] = EMPTY
}

func TestBestCost2(t *testing.T) {
	state := NewState2(B, D, D, A, C, C, B, D, B, B, A, C, D, A, C, A)
	tryMove(t, state, 23, 10, "m1", D)
	tryMove(t, state, 24, 0, "m2", A)
	tryMove(t, state, 19, 9, "m3", B)
	tryMove(t, state, 20, 7, "m4", B)
	tryMove(t, state, 21, 1, "m5", A)
	tryMove(t, state, 15, 21, "m6", C)
	tryMove(t, state, 16, 20, "m7", C)
	tryMove(t, state, 17, 5, "m8", B)
	tryMove(t, state, 18, 3, "m9", D)
	tryMove(t, state, 5, 18, "m10", B)
	tryMove(t, state, 7, 17, "m11", B)
	tryMove(t, state, 9, 16, "m12", B)
	tryMove(t, state, 25, 19, "m13", C)
	tryMove(t, state, 26, 9, "m14", A)
	tryMove(t, state, 3, 26, "m15", D)
	tryMove(t, state, 11, 15, "m16 ", B)
	tryMove(t, state, 12, 25, "m17 ", D)
	tryMove(t, state, 13, 3, "m18", D)
	tryMove(t, state, 1, 13, "m19", A)
	tryMove(t, state, 0, 12, "m20", A)
	tryMove(t, state, 3, 24, "m21", D)
	tryMove(t, state, 9, 11, "m22", A)
	tryMove(t, state, 10, 23, "m23", D)
	lib.AssertTrue(t, "", state.IsFinal())
}

// 0  1  2  3  4  5  6  7  8  9  10
//       11    15    19    23
//       12    16    20    24
//       13    17    21    25
//       14    18    22    26
//

// 0  1  2  3  4  5  6  7  8  9  10
//       11    13    15    17
//       12    14    16    18

// ############
// #...........#
// ###B#C#B#D###
//   #A#D#C#A#
// #########

func TestSample(t *testing.T) {
	state := NewState(12, 18, 11, 15, 13, 16, 14, 17)
	d := BestCost(state, "", map[string]int{}, 15)

	lib.AssertEq(t, "", 12521, d)
}

func TestSample2(t *testing.T) {

	state := NewState2(B, D, D, A, C, C, B, D, B, B, A, C, D, A, C, A)

	d := BestCost(state, "", map[string]int{}, 35)

	lib.AssertEq(t, "", 44169, d)
}
