package domain

// Count the number of 1's in an int.
func bitsSet(v uint) (n int) {
	for v > 0 {
		v &= v - 1
		n++
	}
	return
}

func bitSet(v uint) (n uint) {
	for v > 1 {
		v >>= 1
		n++
	}
	return
}
