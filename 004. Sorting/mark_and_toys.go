package main

import (
	"fmt"
	"sort"
)

// Complete the maximumToys function below.
func maximumToys(prices []int32, k int32) int32 {
	var spendCounter int32

	sort.Slice(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})

	fmt.Println(prices)
	for j := 0; j < len(prices); j++ {
		if k-prices[j] > 0 {
			k -= prices[j]
			spendCounter++
		}
	}

	return spendCounter
}

func main() {
	fmt.Println(maximumToys([]int32{1, 12, 5, 111, 200, 1000, 10}, 50))
}
