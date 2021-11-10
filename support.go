package CKK

import "math"

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func getSum(arr *[]int) (s int) {
	for _, v := range *arr {
		s += v
	}
	return
}
