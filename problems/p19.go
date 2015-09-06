package problems

import (
	"github.com/tn5993/projecteuler/utils/calendar"
)

func NewProblem19() IProblem {
	return Problem19{}
}

type Problem19 struct {
}

func (p Problem19) Solve() int64 {
	var counter int64 = 0
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			if calendar.GetDayOfWeek(m, 1, y) == 0 {
				counter += 1
			}
		}
	}

	return counter
}
