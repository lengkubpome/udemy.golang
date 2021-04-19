package benchmark

import "math"

// 7867
func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func isPrime2(x int) bool {
	for i := 2; i < int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func isPrime3(x int) bool {
	upper := int(math.Sqrt(float64(x)))
	for i := 2; i < upper; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
