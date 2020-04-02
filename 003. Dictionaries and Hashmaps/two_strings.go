package main

import (
	"fmt"
	"strings"
)

// Complete the twoStrings function below.
func twoStrings(s1 string, s2 string) string {

	wordCounter := map[string]int{}

	for _, v := range s1 {
		wordCounter[string(v)]++
	}

	for _, v := range s2 {
		if _, ok := wordCounter[string(v)]; ok {
			return "YES"
		}
	}

	return "NO"
}

//Alternative solution with go built-in "strings"
func twoStringsAlternative(s1 string, s2 string) string {
	if strings.ContainsAny(s1, s2) || strings.ContainsAny(s2, s1) {
		return "YES"
	}
	return "NO"
}

func main() {
	fmt.Println(twoStrings("hi", "world"))
	fmt.Println(twoStringsAlternative("hi", "world"))
}
