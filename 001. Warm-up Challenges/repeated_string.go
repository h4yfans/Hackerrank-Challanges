package main

import "fmt"

func repeatedString(s string, n int64) int64 {
	var counter int64
	for _, v := range s {
		if string(v) == "a" {
			counter++
		}
	}

	division := n / int64(len(s))

	counter = counter * division

	mod := n % int64(len(s))

	for _, v := range s[:mod] {
		if string(v) == "a" {
			counter++
		}
	}

	return counter
}

func main() {
	fmt.Println(repeatedString("aba", 10))
}
