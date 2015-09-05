package problems

import (
	"github.com/tn5993/projecteuler/utils/simplemath/dynamic"
)

type Problem14 struct {
	limit int64
}

func NewProblem14(limit int64) IProblem {
	return Problem14{limit}
}

func (p Problem14) Solve() int64 {
	return dynamic.FindLongestCollatzSequence(p.limit)
}
