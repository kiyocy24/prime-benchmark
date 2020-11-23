package eratosthenes

import (
	"math"
	"math/big"
)

type NumberBig struct {
	value   *big.Int
	isPrime bool
}

func GenarateBig(max int64) []*big.Int {
	numberBigs := make([]*NumberBig, max)
	for n := int64(0); n < max; n++ {
		ng := &NumberBig{
			value:   big.NewInt(n),
			isPrime: true,
		}
		numberBigs[n] = ng
	}

	two := big.NewInt(2)
	limit := big.NewInt(int64(math.Sqrt(float64(max + 1))))
	for _, ng := range numberBigs {
		if ng.value.Cmp(two) == -1 || ng.value.Cmp(limit) == 1 {
			continue
		}

		if ng.isPrime {
			idx := int64(ng.value.Int64() * 2)
			for i := int64(2); idx < int64(len(numberBigs)); i++ {
				numberBigs[idx].isPrime = false
				idx = ng.value.Int64() * i
			}
		}
	}

	var primes []*big.Int
	for _, ng := range numberBigs {
		if ng.value.Cmp(two) == -1 {
			continue
		}
		if ng.isPrime {
			primes = append(primes, ng.value)
		}
	}

	return primes
}
