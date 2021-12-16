package main

import (
	"aoc21/lib"
	"log"
	"os"
)

// Usage:
//   go run . < input

// https://adventofcode.com/2021/day/11

func main() {
	lines := lib.ReadLines(os.Stdin)
	bits := ParseHex(lines[0])
	packet := ParsePacket(bits, 0)
	sum := SumVersions(packet)
	log.Println("sum of versions: ", sum)

	val := ComputeExpression(packet)
	log.Println("result of message: ", val)
}

func ParseHex(st string) []uint8 {
	ar := make([]uint8, 4*len(st))

	for i, c := range st {
		var d uint8
		if '0' <= c && c <= '9' {
			d = uint8(c - '0')
		} else {
			d = uint8(10 + c - 'A')
		}

		for j := 0; j < 4; j++ {
			ar[i*4+3-j] = d >> j & 1
		}
	}

	return ar
}

func BinaryToText(bits []uint8) string {
	st := ""
	for _, b := range bits {
		if b > 0 {
			st = st + "1"
		} else {
			st = st + "0"
		}
	}
	return st
}

const (
	LITERAL = iota
)

type Packet struct {
	packetVersion int
	packetType    int
	packetLiteral int
	packetLength  int
	subPackets    []Packet
}

func ParsePacket(bits []uint8, offset int) Packet {
	subPackets := []Packet{}

	pos := offset
	packetVersion := ReadInt(bits, pos, 3)
	pos += 3
	packetType := ReadInt(bits, pos, 3)
	pos += 3

	packetLiteral := 0

	if packetType == 4 {
		lastPacket := false
		for !lastPacket {
			if bits[pos] == 0 {
				lastPacket = true
			}
			pos += 1
			packetLiteral = packetLiteral<<4 + ReadInt(bits, pos, 4)
			log.Println("pos", pos, "literal", packetLiteral)
			pos += 4
		}
	} else {
		lengthTypeId := bits[pos]
		pos += 1

		if lengthTypeId == 0 {
			totalSubPacketLength := ReadInt(bits, pos, 15)
			pos += 15
			log.Println("totalSubPacketLength", totalSubPacketLength)
			for n := 0; n < totalSubPacketLength; {
				sub := ParsePacket(bits, pos)
				pos += sub.packetLength
				n += sub.packetLength
				log.Println(" read subpacket, n=", n)
				subPackets = append(subPackets, sub)
			}
		} else {
			totalNumSubPackets := ReadInt(bits, pos, 11)
			pos += 11
			for i := 0; i < totalNumSubPackets; i++ {
				sub := ParsePacket(bits, pos)
				pos += sub.packetLength
				subPackets = append(subPackets, sub)
			}
		}
	}

	//// increment to next hex digit boundary
	//for pos%4 > 0 {
	//	pos++
	//}

	return Packet{
		packetVersion: packetVersion,
		packetType:    packetType,
		packetLiteral: packetLiteral,
		packetLength:  pos - offset,
		subPackets:    subPackets,
	}

}

func ReadInt(bits []uint8, offset int, numBits int) int {
	val := 0
	for i := 0; i < numBits; i++ {
		val = val<<1 + int(bits[offset+i])
	}
	return val
}

func SumVersions(packet Packet) int {
	sum := packet.packetVersion
	for _, q := range packet.subPackets {
		sum += SumVersions(q)
	}
	return sum
}

func ComputeExpression(packet Packet) int {

	ar := make([]int, len(packet.subPackets))
	for i, q := range packet.subPackets {
		ar[i] = ComputeExpression(q)
	}

	switch packet.packetType {
	case 0: // sum
		return fold(0, ar, func(i, j int) int { return i + j })
	case 1: // product
		return fold(1, ar, func(i, j int) int { return i * j })
	case 2: // minimum
		return fold(1_000_000_000, ar, func(i, j int) int {
			if i < j {
				return i
			} else {
				return j
			}
		})
	case 3: // maximum
		return fold(-1_000_000_000, ar, func(i, j int) int {
			if i > j {
				return i
			} else {
				return j
			}
		})
	case 4: // literal
		return packet.packetLiteral
	case 5: // gt
		if ar[0] > ar[1] {
			return 1
		} else {
			return 0
		}
	case 6: // lt
		if ar[0] < ar[1] {
			return 1
		} else {
			return 0
		}
	case 7: // eq
		if ar[0] == ar[1] {
			return 1
		} else {
			return 0
		}
	}
	log.Fatal("bad packetType ", packet.packetType)
	return 0
}

func fold(init int, ar []int, fn func(int, int) int) int {
	x := init
	for _, n := range ar {
		x = fn(x, n)
	}
	return x
}
