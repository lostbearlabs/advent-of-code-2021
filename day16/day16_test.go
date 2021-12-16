package main

import (
	"aoc21/lib"
	"testing"
)

func TestParseHex(t *testing.T) {
	lib.AssertEqStr(t, "5", "0101", BinaryToText(ParseHex("5")))
	lib.AssertEqStr(t, "A", "1010", BinaryToText(ParseHex("A")))
	lib.AssertEqStr(t, "F", "1111", BinaryToText(ParseHex("F")))

	text := "38006F45291200"
	expectedBinaryText := "00111000000000000110111101000101001010010001001000000000"
	binary := ParseHex(text)
	binaryText := BinaryToText(binary)
	lib.AssertEqStr(t, "", expectedBinaryText, binaryText)
}

func TestReadInt(t *testing.T) {
	bits := []uint8{1, 0, 1}
	lib.AssertEq(t, "3 bits", 5, ReadInt(bits, 0, 3))
}

func TestParsePacket(t *testing.T) {
	bits := ParseHex("D2FE28")
	packet := ParsePacket(bits, 0)
	lib.AssertEq(t, "packetVersion", 6, packet.packetVersion)
	lib.AssertEq(t, "packetType", 4, packet.packetType)
	lib.AssertEq(t, "packetLiteral", 2021, packet.packetLiteral)
}

func TestParseSubPacketType0(t *testing.T) {
	bits := ParseHex("38006F45291200")
	packet := ParsePacket(bits, 0)
	lib.AssertEq(t, "packetVersion", 1, packet.packetVersion)
	lib.AssertEq(t, "packetType", 6, packet.packetType)
	lib.AssertEq(t, "packetLiteral", 0, packet.packetLiteral)
	lib.AssertEq(t, "len(sub)", 2, len(packet.subPackets))
	lib.AssertEq(t, "sub[0].packetLiteral", 10, packet.subPackets[0].packetLiteral)
	lib.AssertEq(t, "sub[1].packetLiteral", 20, packet.subPackets[1].packetLiteral)
}

func TestParseSubPacketType1(t *testing.T) {
	bits := ParseHex("EE00D40C823060")
	packet := ParsePacket(bits, 0)
	lib.AssertEq(t, "packetVersion", 7, packet.packetVersion)
	lib.AssertEq(t, "packetType", 3, packet.packetType)
	lib.AssertEq(t, "packetLiteral", 0, packet.packetLiteral)
	lib.AssertEq(t, "len(sub)", 3, len(packet.subPackets))
	lib.AssertEq(t, "sub[0].packetLiteral", 1, packet.subPackets[0].packetLiteral)
	lib.AssertEq(t, "sub[1].packetLiteral", 2, packet.subPackets[1].packetLiteral)
	lib.AssertEq(t, "sub[2].packetLiteral", 3, packet.subPackets[2].packetLiteral)
}

func TestSumVersions(t *testing.T) {
	lib.AssertEq(t, "packetVersion1", 16, SumVersions(ParsePacket(ParseHex("8A004A801A8002F478"), 0)))
	lib.AssertEq(t, "packetVersion2", 12, SumVersions(ParsePacket(ParseHex("620080001611562C8802118E34"), 0)))
	lib.AssertEq(t, "packetVersion3", 23, SumVersions(ParsePacket(ParseHex("C0015000016115A2E0802F182340"), 0)))
	lib.AssertEq(t, "packetVersion4", 31, SumVersions(ParsePacket(ParseHex("A0016C880162017C3686B18A3D4780"), 0)))
}

func TestComputeExpression(t *testing.T) {
	lib.AssertEq(t, "1", 3, ComputeExpression(ParsePacket(ParseHex("C200B40A82"), 0)))
	lib.AssertEq(t, "2", 54, ComputeExpression(ParsePacket(ParseHex("04005AC33890"), 0)))
	lib.AssertEq(t, "3", 7, ComputeExpression(ParsePacket(ParseHex("880086C3E88112"), 0)))
	lib.AssertEq(t, "4", 9, ComputeExpression(ParsePacket(ParseHex("CE00C43D881120"), 0)))
	lib.AssertEq(t, "5", 1, ComputeExpression(ParsePacket(ParseHex("D8005AC2A8F0"), 0)))
	lib.AssertEq(t, "6", 0, ComputeExpression(ParsePacket(ParseHex("F600BC2D8F"), 0)))
	lib.AssertEq(t, "7", 0, ComputeExpression(ParsePacket(ParseHex("9C005AC2F8F0"), 0)))
	lib.AssertEq(t, "8", 1, ComputeExpression(ParsePacket(ParseHex("9C0141080250320F1802104A08"), 0)))
}
