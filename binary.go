package binary

import "math"

func PaddedBinaryArray(source []int, length int) []int {
	if len(source) == length {
		return source
	}

	arr := make([]int, length, length)

	for i := 0; i < len(source); i++ {
		arr[len(arr)-1-i] = source[len(source)-1-i]
	}

	return arr
}

func IntToBinaryArray(x int) []int {
	arrLen := int(math.Ceil(math.Log2(float64(x + 1))))
	arr := make([]int, arrLen, arrLen)

	if x == 0 {
		return arr
	}

	for i := 0; i < arrLen; i++ {
		if x >= int(math.Pow(2, float64(arrLen-1-i))) {
			arr[i] = 1
			x = x - int(math.Pow(2, float64(arrLen-1-i)))
		} else if x == 0 {
			return arr
		}
	}

	return arr
}

func BinaryArrayToInt(arr *[]int) int {
	sum := 0

	iterArr := reverseArray(*arr)

	for i := 0; i < len(iterArr); i++ {
		if iterArr[i] == 1 {
			sum += int(math.Pow(2, float64(i)))
		}
	}

	return sum
}

func reverseArray(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	revArr := make([]int, len(arr), len(arr))

	for i := 0; i < len(arr); i++ {
		revArr[i] = arr[len(arr)-1-i]
	}

	return revArr
}
