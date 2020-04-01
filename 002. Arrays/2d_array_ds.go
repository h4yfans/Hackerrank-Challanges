package main

import (
	"fmt"
	"math"
)

func main() {
	var arr [][]int32
	a1 := []int32{0, -4, -6, 0, -7, -6}
	a2 := []int32{-1, -2, -6, -8, -3, -1}
	a3 := []int32{-8, -4, -2, -8, -8, -6}
	a4 := []int32{-3, -1, -2, -5, -7, -4}
	a5 := []int32{-3, -5, -3, -6, -6, -6}
	a6 := []int32{-3, -6, 0, -8, -6, -7}

	arr = append(arr, a1, a2, a3, a4, a5, a6)

	fmt.Println(hourglassSum(arr))
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var maxSum int32 = -math.MaxInt32

	for i := 0; i < len(arr)-2; i++ {
		for k := 0; k < len(arr)-2; k++ {
			a, b, c, d, e, f, g := arr[i][k], arr[i][k+1], arr[i][k+2], arr[i+1][k+1], arr[i+2][k], arr[i+2][k+1], arr[i+2][k+2]
			total := a + b + c + d + e + f + g
			//fmt.Println(a, b, c, d, e, f, g)
			//fmt.Println("total", total)
			if maxSum < total {
				maxSum = total
			}
		}
	}
	return maxSum
}
