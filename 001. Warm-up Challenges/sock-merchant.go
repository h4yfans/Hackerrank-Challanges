package main

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	m := make(map[int32]int)

	for _, number := range ar {
		_, found := m[number]
		if !found {
			m[number] = 1
		} else {
			m[number] += 1
		}
	}

	var result int
	for _, value := range m {
		result += value / 2
	}

	return int32(result)
}
