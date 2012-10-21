package domain

// Count the number of 1's in an int.
func BitsSet(v uint) (n int) {
	for v > 0 {
		v &= v - 1
		n++
	}
	return
}

// Return the number of the actived bit (zero indexed).
// As a precondition, only one bit should be enabled.
func BitSet(v uint) (n uint) {
	for v > 1 {
		v >>= 1
		n++
	}
	return
}
