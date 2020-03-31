package main

import "fmt"

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	var count int32
	length := len(c)
	for i := 0; i < length-1; {
		if i+2 < length && c[i+2] != 1 {
			count++
			i += 2
		} else {
			count++
			i++
		}
	}
	return count
}

func main() {
	fmt.Println(jumpingOnClouds([]int32{0, 0, 0, 0, 1, 0}))
}
