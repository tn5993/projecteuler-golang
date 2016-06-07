package problems

import (
	"github.com/tn5993/projecteuler/utils/simplemath/prime"
)

func NewProblem27() IProblem {
	return problem27{}
}

type problem27 struct{}

func (p problem27) Solve() int64 {
	primes := prime.GetPrimesUnder(999)
	var maxN, A, B int64 = 0, 0, 0

	for _, b := range primes {
		for a := int64(-999); a <= 999; a += 2 {
			n := int64(0)

			for prime.IsPrime(n*n + a*n + b) {
				n++
			}

			if maxN < n {
				maxN, A, B = n, a, b
			}
		}
	}

	return A * B
}

/*
 n=0 => b = prime and b is odd
 n=1 => 1 + a + b = prime. Prime and b is odd so a is odd as (1 + b) is even
 n=2 => 4 + 2a + b = prime.
*/
/*
Euler discovered the remarkable quadratic formula:

n² + n + 41

It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39. However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and certainly when n = 41, 41² + 41 + 41 is clearly divisible by 41.

The incredible formula  n² − 79n + 1601 was discovered, which produces 80 primes for the consecutive values n = 0 to 79. The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

n² + an + b, where |a| < 1000 and |b| < 1000

where |n| is the modulus/absolute value of n
e.g. |11| = 11 and |−4| = 4
Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n = 0.
*/
