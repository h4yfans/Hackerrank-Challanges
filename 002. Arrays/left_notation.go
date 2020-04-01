package main

import "fmt"

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	len := len(a)
	mod := d % int32(len)

	firstPart := a[:mod]
	secondPart := a[mod:]

	var result []int32
	for _, v := range secondPart {
		result = append(result, v)
	}

	for _, v := range firstPart {
		result = append(result, v)
	}

	return result
}

func main() {
	fmt.Println(rotLeft([]int32{98, 67, 35, 1, 74, 79, 7, 26, 54, 63, 24, 17, 32, 81}, 5))
	//[79 7 26 54 63 24 17 32 81 98 67 35 1 74]
}
