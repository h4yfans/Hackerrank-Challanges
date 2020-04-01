package main

import "fmt"

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	var bribe int32

	arrlen := int32(len(q))

	for i := arrlen - 1; i >= 0; i-- {
		if q[i]-i+1 > 2 {
			fmt.Println("Too chaotic")
			return
		}

		for j := i - 1; i > 0 && j >= 0; j-- {
			if q[j] > q[i] {
				bribe++
			}
		}
	}

	fmt.Println(bribe)
}

func main() {
	minimumBribes([]int32{5, 1, 2, 3, 7, 8, 6, 4})

}
