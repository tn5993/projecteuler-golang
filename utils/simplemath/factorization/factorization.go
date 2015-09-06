package factorization

import (
	"github.com/tn5993/projecteuler/utils/simplemath/prime"
	"math/big"
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

/*
http://mathschallenge.net/library/number/sum_of_divisors
*/
func DivisorSum(n int64) int64 {
	if n == 1 {
		return 1
	}

	original := n
	primes := prime.GetPrimesUnder(n)
	var i, count, product int64 = 0, 1, 1
	expInt := new(big.Int)
	for n > 1 {
		if n%primes[i] == 0 {
			count += 1
			n = n / primes[i]
		} else {
			if count > 1 {
				upper := expInt.Exp(big.NewInt(primes[i]), big.NewInt(count), nil).Int64() - 1
				product *= (upper / (primes[i] - 1))
			}
			count = 1
			i += 1
		}
	}

	if count > 1 {
		upper := expInt.Exp(big.NewInt(primes[i]), big.NewInt(count), nil).Int64() - 1
		product *= upper / (primes[i] - 1)
	}

	product -= original

	return product
}
