package radix

import (
	"math"
)

func getBits(a int) int {
	i := 0
	for 0 != a {
		i++
		a /= 10
	}
	return i
}

func getHexBits(a int) int {
	i := 0
	for 0 != a {
		i++
		a >>= 4
	}
	return i
}

func sum(a []int) {
	length := len(a)
	for i := 1; i < length; i++ {
		a[i] += a[i-1]
	}
}

func radixSort(original []int) {
	if 0 == len(original) {
		return
	}

	// the bits of the digit
	bits := getBits(original[0])

	for i := 1; i <= bits; i++ {
		// the length of the bucket is 10 because of the bit range is 0~9
		bucket := make([][]int, 10)
		divisor := int(math.Pow10(i - 1))

		for j := range original {
			// get the number of the i'th bit
			n := original[j] / divisor % 10

			bucket[n] = append(bucket[n], original[j])
		}

		position := 0

		for _, b := range bucket {
			bLen := len(b)
			if bLen == 0 {
				continue
			}

			copy(original[position:], b)
			position += bLen
		}
	}
}

func radixSort1(original []int) {
	length := len(original)

	if 0 == length {
		return
	}

	// the bits of the digit
	bits := getHexBits(original[0])
	r := make([]int, length)
	bitRestore := make([]int, length)

	for i := 1; i <= bits; i++ {
		// the length of the bucket is 16 because of the bit range is 0~15
		bucket := make([]int, 16)
		rightShift := 4 * (i - 1)

		for i := length - 1; i >= 0; i-- {
			// get the number of the i'th bit(0x3039 >> (4*(1-1)) & 0xf = 9)
			n := original[i] >> rightShift & 0xf
			bucket[n]++
			bitRestore[i] = n
		}

		sum(bucket)

		for i := length - 1; i >= 0; i-- {
			element := bitRestore[i]
			position := bucket[element] - 1
			r[position] = original[i]
			bucket[element]--
		}

		copy(original, r)
	}
}

func radixSort2(original []int) {
	length := len(original)

	if 0 == length {
		return
	}

	// the bits of the digit
	bits := getBits(original[0])
	r := make([]int, length)
	bitRestore := make([]int, length)

	for i := 1; i <= bits; i++ {
		// the length of the bucket is 10 because of the bit range is 0~9
		bucket := make([]int, 10)
		divisor := int(math.Pow10(i - 1))

		for i := length - 1; i >= 0; i-- {
			// get the number of the i'th bit
			n := original[i] / divisor % 10
			bucket[n]++
			bitRestore[i] = n
		}

		sum(bucket)

		for i := length - 1; i >= 0; i-- {
			element := bitRestore[i]
			position := bucket[element] - 1
			r[position] = original[i]
			bucket[element]--
		}

		copy(original, r)
	}
}
