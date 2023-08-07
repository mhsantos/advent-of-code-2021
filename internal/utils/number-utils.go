package utils

import "math"

// InvertBits takes an array of 0s and 1s and returns an array with the same number of elements
// but with the values flipped. For example, [0,0,1,0,0,1] returns [1,1,0,1,1,0]
func InvertBits(bits []int) []int {
	var inverted []int
	for _, val := range bits {
		if val == 0 {
			inverted = append(inverted, 1)
		} else {
			inverted = append(inverted, 0)
		}
	}
	return inverted
}

// BitsToBase10 converts an int array containing 0's and 1's to a base10 representation of that.
// for example [0,1,0,0] returns 4 and [1,0,0,1] returns 9
func BitsToBase10(bits []int) int {
	base10 := 0
	N := len(bits)
	for i := 0; i < N; i++ {
		if bits[N-i-1] == 1 {
			base10 += int(math.Pow(2, float64(i)))
		}
	}
	return base10
}
