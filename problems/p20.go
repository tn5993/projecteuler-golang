package problems

import (
	"github.com/tn5993/projecteuler/utils/simplemath/intmath"
)

func NewProblem20(factorial int) IProblem {
	return Problem20{factorial}
}

type Problem20 struct {
	factorial int
}

func (p Problem20) Solve() int64 {
	fac, err := intmath.BigFactorial(p.factorial)
	check(err)
	return intmath.SumDigit(fac).Int64()
}
