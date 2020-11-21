package eratosthenes

import "math"

func Genarate(max int64) []int64 {
	isPrimes := make([]bool, max+1)
	for n := int64(0); n <= max; n++ {
		isPrimes[n] = true
	}

	limit := int64(math.Sqrt(float64(max)))
	for n, isPrime := range isPrimes {
		if n < 2 || limit < int64(n) {
			continue
		}

		idx := int64(n * 2)
		if isPrime {
			for i := int64(2); idx < int64(len(isPrimes)); i++ {
				isPrimes[idx] = false
				idx = int64(n) * i
			}
		}
	}

	var primes []int64
	for n, isPrime := range isPrimes {
		if n < 2 {
			continue
		}
		if isPrime {
			primes = append(primes, int64(n))
		}
	}

	return primes
}
