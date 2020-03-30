package main

import "fmt"

func countingValleys(n int, s string) int32 {
	var level, valleyCounter int32

	for _, ch := range s {
		if ch == 'U' {
			level++
		}
		if ch == 'D' {
			level--
		}

		if level == 0 && ch == 'U' {
			valleyCounter++
		}
	}

	return valleyCounter
}

func main() {
	fmt.Println(countingValleys(8, "UDDDUDUU"))
}
