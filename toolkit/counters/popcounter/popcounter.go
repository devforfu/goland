package popcounter


var pc[256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i & 1)
	}
}

const bitsPerByte = 8

/// Calculates number of non-zero bits in uint64 number.
func PopCount(x uint64) int {
	var total byte
	var i uint
	for i = 0; i < 8; i++ {
		total += pc[byte(x >> (i * bitsPerByte))]
	}
	return int(total)
}