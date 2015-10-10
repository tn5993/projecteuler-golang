package problems

import (
	fac "github.com/tn5993/projecteuler/utils/simplemath/factorization"
)

func NewProblem23(limit int64) IProblem {
	return Problem23{limit}
}

type Problem23 struct {
	limit int64
}

func (p Problem23) Solve() int64 {
	abundants := findAllAbundantsWithin(p.limit)
	abundantSums := make([]bool, p.limit+1)

	for _, a1 := range abundants {
		for _, a2 := range abundants {
			if sum := a1 + a2; sum <= p.limit {
				abundantSums[sum] = true
			}
		}
	}

	//Sum them up
	var result int64
	for i, isAbundantSum := range abundantSums {
		if !isAbundantSum {
			result += int64(i)
		}
	}

	return result
}

func findAllAbundantsWithin(limit int64) []int64 {
	var abundants []int64
	for n := int64(1); n <= limit; n++ {
		if isAbundant(n) {
			abundants = append(abundants, n)
		}
	}

	return abundants
}

func isAbundant(n int64) bool {
	return fac.DivisorSum(n) > n
}
