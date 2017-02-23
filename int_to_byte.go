package convert

import "encoding/binary"

// Itob convert Integer to uint64 BigEndian byte sequence
func Itob8(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// Itob16 convert Integer to 8 byte length uint64 BigEndian byte sequence
func Itob16(v int) []byte {
	b := make([]byte, 16)
	binary.BigEndian.PutUint16(b, uint16(v))
	return b
}

// Itob32 convert Integer to 8 byte length uint64 BigEndian byte sequence
func Itob32(v int) []byte {
	b := make([]byte, 32)
	binary.BigEndian.PutUint32(b, uint32(v))
	return b
}

// Itob32 convert Integer to 8 byte length uint64 BigEndian byte sequence
func Itob64(v int) []byte {
	b := make([]byte, 64)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
