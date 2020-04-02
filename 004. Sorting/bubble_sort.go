package main

import "fmt"

// Complete the countSwaps function below.
func countSwaps(a []int32) {
	var swapCount int

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				temp := a[j]
				a[j] = a[i]
				a[i] = temp

				swapCount++
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", swapCount)
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d\n", a[len(a)-1])

}

func main() {
	countSwaps([]rune{4, 5, 2, 6, 1})
}

//Array is sorted in 6 swaps.
//First Element: 1
//Last Element: 6
