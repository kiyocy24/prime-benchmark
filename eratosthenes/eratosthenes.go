package eratosthenes

import "math"

type Number struct {
	value   int64
	isPrime bool
}

func Genarate(max int64) []int64 {
	numbers := make([]*Number, max)
	for n := int64(0); n < max; n++ {
		num := &Number{
			value:   n,
			isPrime: true,
		}
		numbers[n] = num
	}

	limit := int64(math.Sqrt(float64(max)))
	for _, num := range numbers {
		if num.value < 2 || limit < num.value {
			continue
		}

		idx := int64(num.value * 2)
		if num.isPrime {
			for i := int64(2); idx < int64(len(numbers)); i++ {
				numbers[idx].isPrime = false
				idx = int64(num.value) * i
			}
		}
	}

	var primes []int64
	for _, num := range numbers {
		if num.value < 2 {
			continue
		}
		if num.isPrime {
			primes = append(primes, num.value)
		}
	}

	return primes
}
