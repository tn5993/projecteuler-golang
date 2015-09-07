package problems

import (
	"fmt"
	fac "github.com/tn5993/projecteuler/utils/simplemath/factorization"
)

func NewProblem21(n int64) IProblem {
	return Problem21{n}
}

type Problem21 struct {
	n int64
}

func (p Problem21) Solve() int64 {
	amicableMap := make(map[int64]int64)
	var sum int64 = 0
	for i := int64(0); i <= p.n; i++ {
		_, hasKey := amicableMap[i]
		if !hasKey {
			hasPair, m := hasAmicable(i)
			if hasPair {
				sum += (i + m)
				amicableMap[m] = i
			}
		}
	}

	return sum
}

func hasAmicable(n int64) (bool, int64) {
	m := fac.DivisorSum(n)
	if m != n && n == fac.DivisorSum(m) {
		fmt.Printf("Found amicable pairs (%d, %d)\n", m, n)
		return true, m
	} else {
		return false, -1
	}
}
