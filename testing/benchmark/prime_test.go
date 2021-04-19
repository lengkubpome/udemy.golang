package benchmark

import "testing"

func TestPrime(t *testing.T) {
	if !isPrime(7867) {
		t.Error("7867 is not prime")

	}

	if !isPrime2(7867) {
		t.Error("7867 is not prime")

	}

	if isPrime(10) {
		t.Error("10 is prime")
	}

	if isPrime2(10) {
		t.Error("10 is prime")
	}

}

// Benchmark[Name] (* testing.B)
func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(7867)
	}
}

func BenchmarkIsPrime2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime2(7867)
	}
}

func BenchmarkIsPrime3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime3(7867)
	}
}

// go test -v -bench .
// go test -v -bench . -benchmem    //เช็คการจองเมมโมรี่

// go test -bench BenchmarkIsPrime$ //เลือกรันเฉพาะ BenchmarkIsPrime
