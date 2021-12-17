package main

import (
	"os"
	"testing"
)

// show log output for tests only
func init() {
	log.SetOutput(os.Stdout)
}

func TestPartOne1(t *testing.T) {
	expected := 2021
	binary, err := hexToBinary("D2FE28")

	packet, tail, err := newPacket(binary)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	val := packet.value

	if val != expected {
		t.Logf("Answer should be %d, but got %d", expected, val)
		t.Fail()
	}

	if tail != "000" {
		t.Log("tail should be 000 but got: ", tail)
		t.Fail()
	}
}

func TestPartOne2(t *testing.T) {
	binary, err := hexToBinary("38006F45291200")

	packet, tail, err := newPacket(binary)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if packet.version != 1 {
		t.Log("packet version should be 1", packet.version)
		t.Fail()
	}

	if packet.typeId != 6 {
		t.Log("packet typeId should be 6", packet.typeId)
		t.Fail()
	}

	lengthTypeId := packet.OperatorPacket.lengthTypeId

	if lengthTypeId != LENGTH_BITS {
		t.Logf("packet typeId should be %d, got %d", LENGTH_BITS, lengthTypeId)
		t.Fail()
	}

	packetCount := len(packet.packets)

	if packetCount != 2 {
		t.Log("packet should have (2) packets:", packetCount)
		t.Fail()
	}

	if tail != "0000000" {
		t.Log("tail should be 0000000 but got: ", tail)
		t.Fail()
	}
}

func TestPartOne3(t *testing.T) {
	binary, err := hexToBinary("EE00D40C823060")

	packet, tail, err := newPacket(binary)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if packet.version != 7 {
		t.Log("packet version should be 7", packet.version)
		t.Fail()
	}

	if packet.typeId != 3 {
		t.Log("packet typeId should be 3", packet.typeId)
		t.Fail()
	}

	lengthTypeId := packet.lengthTypeId

	if lengthTypeId != LENGTH_PACKETS {
		t.Logf("packet typeId should be %d, got %d", LENGTH_BITS, lengthTypeId)
		t.Fail()
	}

	packetCount := len(packet.packets)

	if packetCount != 3 {
		t.Log("packet should have (3) packets:", packetCount)
		t.Fail()
	}

	if tail != "00000" {
		t.Log("tail should be 00000 but got: ", tail)
		t.Fail()
	}
}

func TestPartOne_VersionSum1(t *testing.T) {
	binary, err := hexToBinary("8A004A801A8002F478")

	packet, _, err := newPacket(binary)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if packet.versionSum() != 16 {
		t.Log("packet version sum should be 16, got:", packet.versionSum())
		t.Fail()
	}
}

func TestPartOne_VersionSum2(t *testing.T) {
	binary, err := hexToBinary("620080001611562C8802118E34")

	packet, _, err := newPacket(binary)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if packet.versionSum() != 12 {
		t.Log("packet version sum should be 12, got:", packet.versionSum())
		t.Fail()
	}
}

func TestPartOne_VersionSum3(t *testing.T) {
	binary, err := hexToBinary("C0015000016115A2E0802F182340")

	packet, _, err := newPacket(binary)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if packet.versionSum() != 23 {
		t.Log("packet version sum should be 23, got:", packet.versionSum())
		t.Fail()
	}
}

func TestPartOne_VersionSum4(t *testing.T) {
	binary, err := hexToBinary("A0016C880162017C3686B18A3D4780")

	packet, _, err := newPacket(binary)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if packet.versionSum() != 31 {
		t.Log("packet version sum should be 31, got:", packet.versionSum())
		t.Fail()
	}
}

func TestPartOne(t *testing.T) {
	// the real data
	content := FileLoader("input.txt")
	binary, err := hexToBinary(content)

	packet, _, err := newPacket(binary)

	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	if packet.versionSum() != 981 {
		t.Log("packet version sum should be 981, got:", packet.versionSum())
		t.Fail()
	}
}

var hexes = map[string]int{
	"C200B40A82":                 3,
	"04005AC33890":               54,
	"880086C3E88112":             7,
	"CE00C43D881120":             9,
	"D8005AC2A8F0":               1,
	"F600BC2D8F":                 0,
	"9C005AC2F8F0":               0,
	"9C0141080250320F1802104A08": 1,
}

func TestPartTwo(t *testing.T) {
	for hex, expected := range hexes {
		binary, err := hexToBinary(hex)

		packet, _, err := newPacket(binary)

		if err != nil {
			t.Log("error should be nil", err)
			t.Fail()
		}

		expr := packet.evaluateExpression()

		if expr != expected {
			t.Logf("packet version sum should be %d, got: %d", expected, expr)
			t.Fail()
		}
	}
}
