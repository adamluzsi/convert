package convert

// Copy will copy a btye slice, this function is useful,
// when you want to use values from boltdb in outside of the transaction view
func Copy(bytes []byte) []byte {
	newBytes := make([]byte, 0, cap(bytes))
	return append(newBytes, bytes...)
}
