package main

import (
	"fmt"
	"sort"
	"strings"
)

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	var count int32
	strCounts := map[string]int32{}

	for i := range s {
		for j := i + 1; j < len(s)+1; j++ {
			substr := strings.Split(s[i:j], "")
			sort.Strings(substr)
			strCounts[strings.Join(substr, "")]++
		}
	}

	for _, v := range strCounts {
		count += v * (v - 1) / 2
	}

	return count
}

func main() {
	fmt.Println(sherlockAndAnagrams("ifailuhkqq"))
}
