package problems

import (
	"math"
)

type Problem9 struct {
}

func NewProblem9() Problem9 {
	return Problem9{}
}

func (p Problem9) Solve() int64 {
	b := 1
	for b < 500 && ((500000-1000*b)%(1000-b) != 0) {
		b += 1
	}

	a := int((500000 - 1000*b) / (1000 - b))
	c := int(math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2)))

	return int64(a * b * c)
}
