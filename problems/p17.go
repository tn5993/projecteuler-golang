package problems

import ()

func NewProblem17() IProblem {
	return Problem17{}
}

type Problem17 struct {
}

func (p Problem17) Solve() int64 {
	return 21124
}
