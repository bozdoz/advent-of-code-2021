package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"bozdoz.com/aoc-2021/utils"
)

type LengthTypeId int

const (
	// the next 15 bits are a number that represents
	// the total length in bits of the sub-packets
	LENGTH_BITS LengthTypeId = iota
	// the next 11 bits are a number that represents
	// the number of sub-packets immediately contained
	LENGTH_PACKETS
)

// contains a literal value
type LiteralValuePacket struct {
	value int
}

// contains one or more packets
type OperatorPacket struct {
	lengthTypeId LengthTypeId
	packets      []Packet
}

type TypeId int

const (
	TYPE_SUM TypeId = iota
	TYPE_PRODUCT
	TYPE_MIN
	TYPE_MAX
	TYPE_LITERAL
	TYPE_GT
	TYPE_LT
	TYPE_EQ
)

type Packet struct {
	version int
	typeId  TypeId
	LiteralValuePacket
	OperatorPacket
}

type Binary string

func (binary Binary) splitAt(index int) (Binary, Binary) {
	return binary[:index], binary[index:]
}

func hexToBinary(hexStr string) (binary Binary, err error) {
	log.Println("hex", hexStr)

	decoded, err := hex.DecodeString(hexStr)

	if err != nil {
		return "", err
	}

	for _, bit := range decoded {
		// the binary number is padded with leading zeroes
		// until its length is a multiple of four bits
		// TODO: is %08b always right?
		binary += Binary(fmt.Sprintf("%08b", int(bit)))
	}

	log.Println("binary", binary)

	return binary, err
}

func newPacket(binary Binary) (packet *Packet, tail Binary, err error) {
	// first 3 bits = version
	head, tail := binary.splitAt(3)
	version, err := utils.BinaryToInt(head)

	// next 3 bits = type id
	head, tail = tail.splitAt(3)
	typeId, err := utils.BinaryToInt(head)

	packet = &Packet{
		version: version,
		typeId:  TypeId(typeId),
	}

	log.Println("packet", packet)

	// typeId 4 is literal value;
	if packet.typeId == TYPE_LITERAL {
		tail, err = packet.parseLiteralValue(tail)
	} else {
		// every other typeId is an "operator"
		tail, err = packet.parseOperator(tail)
	}

	return packet, tail, err
}

func (packet *Packet) parseLiteralValue(binary Binary) (tail Binary, err error) {
	// "groups of five bits"
	chunkSize := 5

	// get literal value from remaining bits
	var valueBin Binary
	head, tail := binary.splitAt(chunkSize)

	for len(head) == 5 {
		// strip off the leading bit
		first, rest := head.splitAt(1)

		valueBin += rest

		// 0 means it is the last group
		if first == "0" {
			break
		}

		head, tail = tail.splitAt(5)
	}

	value, err := utils.BinaryToInt(valueBin)

	if err != nil {
		return tail, err
	}

	log.Println("value", value)

	packet.value = value

	return tail, nil
}

func (packet *Packet) parseOperator(binary Binary) (tail Binary, err error) {
	// get length type id
	head, tail := binary.splitAt(1)

	if head == "0" {
		packet.lengthTypeId = LENGTH_BITS
		tail, err = packet.parseOperatorBitCount(tail)
	} else {
		packet.lengthTypeId = LENGTH_PACKETS
		tail, err = packet.parseOperatorPacketCount(tail)
	}

	return tail, err
}

// the next 15 bits are a number that represents
// the total length in bits of the sub-packets
func (packet *Packet) parseOperatorBitCount(binary Binary) (tail Binary, err error) {
	if packet.lengthTypeId != LENGTH_BITS {
		return binary, errors.New("can't parse bit count from packet count")
	}

	// 15 bits
	head, tail := binary.splitAt(15)

	bitCount, err := utils.BinaryToInt(head)

	log.Println("bitcount", bitCount)

	if err != nil {
		return binary, err
	}

	subpackets, tail := tail.splitAt(bitCount)

	log.Println("subpackets", subpackets)
	log.Println("tail", tail)

	// could be adjacent or nested packets
	// TODO: bit of guessing here...
	// ? discard any tail values that are exactly 0
	for subpackets != Binary(strings.Repeat("0", len(subpackets))) {
		subpacket, nextPackets, err := newPacket(subpackets)

		if err != nil {
			return tail, nil
		}

		packet.packets = append(packet.packets, *subpacket)

		subpackets = nextPackets
	}

	return tail, nil
}

// the next 11 bits are a number that represents
// the number of sub-packets immediately contained
func (packet *Packet) parseOperatorPacketCount(binary Binary) (tail Binary, err error) {
	if packet.lengthTypeId != LENGTH_PACKETS {
		return binary, errors.New("can't parse packet count from bit count")
	}

	// 11 bits
	head, tail := binary.splitAt(11)

	packetCount, err := utils.BinaryToInt(head)

	log.Println("packetCount", packetCount)

	if err != nil {
		return binary, err
	}

	for packetCount > 0 {
		packetCount--

		subpacket, nextTail, err := newPacket(tail)

		if err != nil {
			return tail, nil
		}

		packet.packets = append(packet.packets, *subpacket)

		tail = nextTail
	}

	return tail, err
}

func (packet *Packet) versionSum() (sum int) {
	sum += packet.version

	// recursively check nested packets
	for _, subpacket := range packet.packets {
		sum += subpacket.versionSum()
	}

	return
}

// Part Two introduces operator logic
func (packet *Packet) evaluateExpression() int {
	if packet.typeId == TYPE_LITERAL {
		return packet.value
	}

	values := []int{}

	// recursively check nested packets
	for _, subpacket := range packet.packets {
		values = append(values, subpacket.evaluateExpression())
	}

	switch packet.typeId {
	case TYPE_SUM:
		return utils.Sum(values...)
	case TYPE_PRODUCT:
		product := values[0]

		for i := 1; i < len(values); i++ {
			product *= values[i]
		}

		return product
	case TYPE_MIN:
		return utils.MinInt(values...)
	case TYPE_MAX:
		return utils.MaxInt(values...)
	case TYPE_GT:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case TYPE_LT:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case TYPE_EQ:
		if values[0] == values[1] {
			return 1
		}
		return 0
	}

	return 0
}
