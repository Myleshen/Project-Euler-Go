package problems

import (
	"math"
)

// Get the largest prime factor for 600851475143

func Problem3(large_number int) []int {
	var factors []int

	// Divide by 2 as many times as there are even numbers / factors like 4, 6, 8, etc...
	if large_number%2 == 0 {
		factors = append(factors, 2)
	}

	for {
		if large_number%2 == 0 {
			large_number = large_number / 2
		} else {
			break
		}
	}

	for i := 3; i < int(math.Sqrt(float64(large_number)))+1; i += 2 {
		if large_number%i == 0 {
			large_number = large_number / i
			factors = append(factors, i)
		}
	}

	if large_number > 0 {
		factors = append(factors, large_number)
	}

	return factors
}
