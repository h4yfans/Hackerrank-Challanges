package main

import "fmt"

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) {

	magazineWordMap := map[string]int{}
	flag := true
	for _, m := range magazine {
		if _, ok := magazineWordMap[m]; !ok {
			magazineWordMap[m] = 1
		} else {
			magazineWordMap[m] += 1
		}
	}

	for _, n := range note {
		if _, ok := magazineWordMap[n]; !ok || magazineWordMap[n] == 0 {
			flag = false
		} else {
			magazineWordMap[n] -= 1
		}
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func main() {
	magazine := []string{"two", "times", "three", "is", "not", "four"}
	notes := []string{"two", "times", "is", "four"}

	checkMagazine(magazine, notes)
}
