package problems

import (
	"github.com/tn5993/projecteuler/utils/simplemath/prime"
)

type Problem10 struct {
	limit int
}

func NewProblem10(limit int) IProblem {
	return Problem10{limit}
}

func (p Problem10) Solve() int64 {
	primeTable := prime.SieveOfE(int64(p.limit))

	sum := 0
	for i, each := range primeTable {
		if i > 1 && !each {
			sum += i
		}
	}

	return int64(sum)
}
