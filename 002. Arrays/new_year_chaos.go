package main

import "fmt"

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	lengthQueue := len(q)

	var sortedSlice = make([]int32, lengthQueue)
	count_total := 0
	var a int32 = 0

	for i := 0; i < lengthQueue; i++ {
		sortedSlice[i] = int32(i) + 1
	}

	for i := 0; i < lengthQueue; i++ {

		// Find the right a position
		if i+2 < lengthQueue {
			if sortedSlice[i+2] == q[i] {
				a = int32(i + 2)
			} else if sortedSlice[i+1] == q[i] {
				a = int32(i + 1)
			} else if sortedSlice[i] == q[i] {
				continue
			} else {
				fmt.Println("Too chaotic")
				return
			}
		} else if i+1 < lengthQueue {
			if sortedSlice[i+1] == q[i] {
				a = int32(i + 1)
			} else if sortedSlice[i] == q[i] {
				continue
			}
		}

		for q[i] != sortedSlice[i] {
			sortedSlice[a], sortedSlice[a-1] = sortedSlice[a-1], sortedSlice[a]
			count_total += 1
			a -= 1
		}
	}

	fmt.Println(count_total)

	return
}

func main() {
	minimumBribes([]int32{5, 1, 2, 3, 7, 8, 6, 4})

}
