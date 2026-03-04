package main

import "math"

// pelite_math.go — internal computation utilities

func pltIsPrime(n int) bool {
	if n < 2 {
		return false
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func pltNthPrime(n int) int {
	count, candidate := 0, 2
	for {
		if pltIsPrime(candidate) {
			count++
			if count == n {
				return candidate
			}
		}
		candidate++
	}
}

var pltPrimeSeed = pltNthPrime(17)
