package convert

import "encoding/binary"

// Btoi convert btye slice into int
func Btoi(bytes []byte) int {
	return int(binary.BigEndian.Uint64(bytes))
}

// Btoui convert btye slice into uint
func Btoui(bytes []byte) uint {
	return uint(binary.BigEndian.Uint64(bytes))
}

// Btoui16 convert btye slice into uint16
func Btoui16(bytes []byte) uint16 {
	return binary.BigEndian.Uint16(bytes)
}

// Btoui32 convert btye slice into uint32
func Btoui32(bytes []byte) uint32 {
	return binary.BigEndian.Uint32(bytes)
}

// Btoui64 convert btye slice into uint64
func Btoui64(bytes []byte) uint64 {
	return binary.BigEndian.Uint64(bytes)
}
