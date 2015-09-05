package problems

import (
	"github.com/tn5993/projecteuler/utils/printer"
	pereader "github.com/tn5993/projecteuler/utils/reader"
	"github.com/tn5993/projecteuler/utils/simplemath/intmath"
)

type Problem11 struct {
	filename string
}

func NewProblem11(filename string) IProblem {
	return Problem11{filename}
}

func (p Problem11) Solve() int64 {
	matrix := pereader.ReadFileToMatrix(p.filename)
	var largest int64 = 0
	for i := range matrix {
		for j := range matrix[i] {
			var product int64
			//Horizontaly -
			if j < 17 {
				product = matrix[i][j] * matrix[i][j+1] * matrix[i][j+2] * matrix[i][j+3]
			}
			largest = intmath.Max64(largest, product)
			//Vertically |
			if i < 17 {
				product = matrix[i][j] * matrix[i+1][j] * matrix[i+2][j] * matrix[i+3][j]
			}
			largest = intmath.Max64(largest, product)
			//Diagonal \
			if i < 17 && j < 17 {
				product = matrix[i][j] * matrix[i+1][j+1] * matrix[i+2][j+2] * matrix[i+3][j+3]
			}
			largest = intmath.Max64(largest, product)
			//Diagonal /
			if i < 17 && j > 2 {
				product = matrix[i][j] * matrix[i+1][j-1] * matrix[i+2][j-2] * matrix[i+3][j-3]
			}
			largest = intmath.Max64(largest, product)
		}
	}
	printer.PrintMatrix(matrix)
	return largest
}
