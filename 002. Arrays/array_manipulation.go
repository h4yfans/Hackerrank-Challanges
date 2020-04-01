package main

import "fmt"

func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int32, n+1)

	for _, q := range queries {
		a := q[0]
		b := q[1]
		k := q[2]

		arr[a-1] += k
		arr[b] -= k
	}

	var max int64
	var sum int64
	for _, i := range arr {
		sum += int64(i)
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	var input = [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}
	fmt.Println(arrayManipulation(10, input))
}
