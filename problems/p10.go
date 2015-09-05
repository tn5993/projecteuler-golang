package problems

import ()

type Problem10 struct {
	limit int
}

func NewProblem10(limit int) IProblem {
	return Problem10{limit}
}

func (p Problem10) Solve() int64 {
	primeTable := make([]byte, p.limit+1)
	for i := 2; i*i <= p.limit; i++ {
		if primeTable[i] == 0 {
			for j := i * i; j <= p.limit; j += i {
				primeTable[j] = 1
			}
		}
	}

	sum := 0
	for i, each := range primeTable {
		if i > 1 && each == 0 {
			sum += i
		}
	}

	return int64(sum)
}
