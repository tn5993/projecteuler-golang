package problems

import ()

type problem30 struct {
}

func NewProblem30() IProblem {
	return problem30{}
}

func (p problem30) Solve() int64 {
	var result int64

	for i := int64(0); i <= 354294; i++ {
		var sum int64
		n := i

		for n > 0 {
			d := n % 10
			sum += (d * d * d * d * d)
			n = n / 10
		}

		if sum == i {
			result += i
		}
	}

	return result
}
