package problems

import (
	"github.com/tn5993/projecteuler/utils/simplemath/intmath"
	"math/big"
)

/*
215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 21000?
*/
func NewProblem16() IProblem {
	return Problem16{}
}

type Problem16 struct {
}

func (p Problem16) Solve() int64 {
	power2 := new(big.Int)
	power2.Exp(big.NewInt(2), big.NewInt(1000), nil)

	return intmath.SumDigit(power2).Int64()
}
