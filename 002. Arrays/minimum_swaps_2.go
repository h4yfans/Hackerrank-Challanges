package main

import "fmt"

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	len := int32(len(arr))
	var swaps int32
	for i := int32(0); i < len; {
		if arr[i] != i+1 {
			temp := arr[i]
			arr[i] = arr[temp-1]
			arr[temp-1] = temp
			swaps++
		} else {
			i++
		}
	}
	return swaps
}

func main() {
	fmt.Println(minimumSwaps([]int32{2, 3, 4, 1, 5})) // 5

}
