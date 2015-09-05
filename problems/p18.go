package problems

import (
	"github.com/tn5993/projecteuler/utils/printer"
	pereader "github.com/tn5993/projecteuler/utils/reader"
	peintmath "github.com/tn5993/projecteuler/utils/simplemath/intmath"
)

func NewProblem18(filename string) IProblem {
	return Problem18{filename}
}

type Problem18 struct {
	filename string
}

func (p Problem18) Solve() int64 {
	var matrix [][]int64 = pereader.ReadFileToMatrix(p.filename)
	printer.PrintMatrix(matrix)
	rLength := len(matrix)
	for row := rLength - 2; row >= 0; row-- {
		cLength := len(matrix[row])
		for col := 0; col < cLength; col++ {
			sum1 := matrix[row][col] + matrix[row+1][col]
			sum2 := matrix[row][col] + matrix[row+1][col+1]
			matrix[row][col] = peintmath.Max64(sum1, sum2)
		}
	}
	return matrix[0][0]
}
