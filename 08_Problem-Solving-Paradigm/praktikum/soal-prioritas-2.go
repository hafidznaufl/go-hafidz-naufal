package main

import (
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	if n <= 1 {
		return 0
	}

	// Buat sebuah slice untuk menyimpan biaya minimum untuk mencapai setiap batu.
	minCost := make([]int, n)
	minCost[0] = 0
	minCost[1] = int(math.Abs(float64(jumps[1] - jumps[0])))

	// Hitung biaya minimum untuk setiap batu dari batu ketiga hingga batu terakhir.
	for i := 2; i < n; i++ {
		option1 := minCost[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))
		option2 := minCost[i-2] + int(math.Abs(float64(jumps[i]-jumps[i-2])))
		minCost[i] = int(math.Min(float64(option1), float64(option2)))
	}

	return minCost[n-1]
}