package main

import (
	"aoc21/lib"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/22

func main() {
	// part 1
	//cmds := ParseInstructions(lib.ReadLines(os.Stdin), 50)

	// part 2
	cmds := ParseInstructions(lib.ReadLines(os.Stdin), 100000000)

	log.Printf("len(cmds)=%d\n", len(cmds))
	cubes := ApplyInstructions(cmds)
	v := cubes.CountVolume()
	log.Printf("volume:%d\n", v)
}

type Region struct {
	x1, x2, y1, y2, z1, z2 int
	on                     bool
}

// *****************
// RegionSet
// *****************

type RegionSet struct {
	ar map[Region]bool
}

func NewRegionSet() RegionSet {
	return RegionSet{ar: map[Region]bool{}}
}

func (regions *RegionSet) CountVolume() int {
	v := 0
	for cube := range regions.ar {
		if cube.on {
			v += volume(cube)
		}
	}
	return v
}

func (regions *RegionSet) Add(region Region) {
	regions.ar[region] = true
}

func (regions *RegionSet) Remove(region Region) {
	delete(regions.ar, region)
}

func (regions *RegionSet) Size() int {
	return len(regions.ar)
}

func (regions *RegionSet) Clone() RegionSet {
	br := NewRegionSet()
	for region := range regions.ar {
		br.ar[region] = true
	}
	return br
}

// *****************
// Parsing
// *****************

func ParseInstruction(line string) Region {
	line = strings.Replace(line, "..", " ", -1)
	line = strings.Replace(line, "x=", " ", -1)
	line = strings.Replace(line, "y=", " ", -1)
	line = strings.Replace(line, "z=", " ", -1)
	line = strings.Replace(line, ",", " ", -1)
	ar := strings.Fields(line)

	cmd := Region{}
	cmd.on = ar[0] == "on"
	cmd.x1, _ = strconv.Atoi(ar[1])
	cmd.x2, _ = strconv.Atoi(ar[2])
	cmd.y1, _ = strconv.Atoi(ar[3])
	cmd.y2, _ = strconv.Atoi(ar[4])
	cmd.z1, _ = strconv.Atoi(ar[5])
	cmd.z2, _ = strconv.Atoi(ar[6])

	// the input says (10,10) is a cube, but internally we'll represent it at (10,11)
	cmd.x2++
	cmd.y2++
	cmd.z2++

	if cmd.x1 > cmd.x2 {
		cmd.x2, cmd.x1 = cmd.x1, cmd.x2
	}
	if cmd.y1 > cmd.y2 {
		cmd.y2, cmd.y1 = cmd.y1, cmd.y2
	}
	if cmd.z1 > cmd.z2 {
		cmd.z2, cmd.z1 = cmd.z1, cmd.z2
	}

	return cmd
}

func restrict(x int, b int) int {
	if x < -b {
		return -b
	}
	if x > b+1 {
		return b + 1
	}
	return x
}

func ParseInstructions(lines []string, boundary int) []Region {
	var ar []Region
	ar = append(ar, Region{
		x1: -boundary,
		x2: boundary + 1,
		y1: -boundary,
		y2: boundary + 1,
		z1: -boundary,
		z2: boundary + 1,
		on: false,
	})
	for _, line := range lines {
		cmd := ParseInstruction(line)

		if cmd.x2 < -boundary ||
			cmd.x1 > boundary ||
			cmd.y2 < -boundary ||
			cmd.y1 > boundary ||
			cmd.z2 < -boundary ||
			cmd.z1 > boundary {
			continue
		}

		cmd.x1 = restrict(cmd.x1, boundary)
		cmd.x2 = restrict(cmd.x2, boundary)
		cmd.y1 = restrict(cmd.y1, boundary)
		cmd.y2 = restrict(cmd.y2, boundary)
		cmd.z1 = restrict(cmd.z1, boundary)
		cmd.z2 = restrict(cmd.z2, boundary)

		ar = append(ar, cmd)
	}
	return ar
}

// *****************
// Helpers
// *****************

func uniqueValues(ar []Region, fn func(Region) (int, int)) []int {
	var br []int

	for _, a := range ar {
		n1, n2 := fn(a)
		br = append(br, n1)
		br = append(br, n2)
	}

	sort.Ints(br)

	var cr []int
	bLast := 0
	for i, b := range br {
		if i == 0 || b != bLast {
			cr = append(cr, b)
			bLast = b
		}
	}

	return cr
}

func AllSubRegions(ar []Region) RegionSet {
	rs := NewRegionSet()
	xr := uniqueValues(ar, func(cmd Region) (int, int) { return cmd.x1, cmd.x2 })
	yr := uniqueValues(ar, func(cmd Region) (int, int) { return cmd.y1, cmd.y2 })
	zr := uniqueValues(ar, func(cmd Region) (int, int) { return cmd.z1, cmd.z2 })

	log.Printf("len(xr)=%d - %v\n", len(xr), xr)
	log.Printf("len(yr)=%d - %v\n", len(yr), yr)
	log.Printf("len(zr)=%d - %v\n", len(zr), zr)

	for i := 1; i < len(xr); i++ {
		for j := 1; j < len(yr); j++ {
			for k := 1; k < len(zr); k++ {
				c := Region{
					x1: xr[i-1],
					x2: xr[i],
					y1: yr[j-1],
					y2: yr[j],
					z1: zr[k-1],
					z2: zr[k],
				}

				rs.Add(c)
			}
		}
	}

	log.Printf("len(cubes)=%d\n", len(rs.ar))
	return rs
}

// *****************
// Algorithm
// *****************

//func regionInside(region Region, cmd Region) bool {
//	xm := region.x1 >= cmd.x1 && region.x2 <= cmd.x2
//	ym := region.y1 >= cmd.y1 && region.y2 <= cmd.y2
//	zm := region.z1 >= cmd.z1 && region.z2 <= cmd.z2
//
//	return xm && ym && zm
//}

//func ApplyInstruction(regions RegionSet, cmd Region) RegionSet {
//	br := NewRegionSet()
//
//	vol := 0
//	nChanged := 0
//	nInside := 0
//	for region := range regions.ar {
//		if regionInside(region, cmd) {
//			nInside++
//			if region.on != cmd.on {
//				region.on = cmd.on
//				vol += volume(region)
//				nChanged++
//			}
//		}
//		br.Add(region)
//	}
//
//	log.Printf("cmd %v, hit %d regions, set %d regions to %t, total volume %d\n", cmd, nInside, nChanged, cmd.on, vol)
//	return br
//}

func ApplyInstruction(regions *RegionSet, cmd Region, xr []int, yr []int, zr []int, cmdNum int) {
	nHit := 0
	nChanged := 0

	for i := 1; i < len(xr); i++ {
		if xr[i] > cmd.x2 || xr[i-1] < cmd.x1 {
			continue
		}

		for j := 1; j < len(yr); j++ {
			if yr[j] > cmd.y2 || yr[j-1] < cmd.y1 {
				continue
			}

			for k := 1; k < len(zr); k++ {
				if zr[k] > cmd.z2 || zr[k-1] < cmd.z1 {
					continue
				}

				nHit++

				c := Region{
					x1: xr[i-1],
					x2: xr[i],
					y1: yr[j-1],
					y2: yr[j],
					z1: zr[k-1],
					z2: zr[k],
					on: true,
				}

				if cmd.on {
					if !regions.ar[c] {
						regions.Add(c)
						nChanged++
					}
				} else {
					if regions.ar[c] {
						regions.Remove(c)
						nChanged++
					}
				}
			}
		}
	}

	log.Printf("cmd %d %v, nHit=%d, nChanged=%d, nTot=%d\n", cmdNum, cmd, nHit, nChanged, regions.Size())
}

func ApplyInstructions(cmds []Region) RegionSet {

	xr := uniqueValues(cmds, func(cmd Region) (int, int) { return cmd.x1, cmd.x2 })
	yr := uniqueValues(cmds, func(cmd Region) (int, int) { return cmd.y1, cmd.y2 })
	zr := uniqueValues(cmds, func(cmd Region) (int, int) { return cmd.z1, cmd.z2 })

	//regions := AllSubRegions(cmds)
	regions := NewRegionSet()

	for i, cmd := range cmds {
		ApplyInstruction(&regions, cmd, xr, yr, zr, i)
	}

	return regions
}

func volume(cube Region) int {
	return (cube.x2 - cube.x1) * (cube.y2 - cube.y1) * (cube.z2 - cube.z1)
}
