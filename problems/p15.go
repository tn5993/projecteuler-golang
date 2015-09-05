package problems

import (
	"github.com/tn5993/projecteuler/utils/simplemath/combinatorics"
)

func NewProblem15(gridSize int64) IProblem {
	return Problem15{gridSize}
}

type Problem15 struct {
	gridSize int64
}

func (p Problem15) Solve() int64 {
	return combinatorics.CentralBinomialCoefficient(p.gridSize)
}
