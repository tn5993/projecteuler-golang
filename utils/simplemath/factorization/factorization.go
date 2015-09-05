package factorization

import (
	"github.com/tn5993/projecteuler/utils/simplemath/prime"
)

func DivisorsOf(n int64) int64 {
	primes := prime.GetPrimesUnder(n)
	var i, sum, product int64 = 0, 0, 1
	for n > 1 {
		if n%primes[i] == 0 {
			sum += 1
			n = n / primes[i]
		} else {
			product *= (sum + 1)
			i += 1
			sum = 0
		}
	}

	product *= (sum + 1)
	return product
}
