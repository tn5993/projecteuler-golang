package problems

import (
	"math"
)

func NewProblem25() IProblem {
	return Problem25{}
}

type Problem25 struct{}

func (p Problem25) Solve() int64 {
	result := (math.Log(10)*999 + math.Log(5)/2) / math.Log(math.Phi)
	result = math.Ceil(result)
	return int64(result)
}
