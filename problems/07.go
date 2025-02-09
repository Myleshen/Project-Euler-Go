package problems

import (
	"math"
)

func Problem7(required_prime_count int) int {
	prime_count := 0
	for i := 2; ; i++ {
		if IsPrime(i) {
			prime_count++
			if prime_count == required_prime_count {
				return i
			}
		}
	}
}

func IsPrime(dividend int) bool {
	max_limit := int(math.Sqrt(float64(dividend))) + 1
	for i := 2; i < max_limit; i++ {
		if dividend%i == 0 {
			return false
		}
	}
	return true
}
