package printer

import "fmt"

func PrintMatrix(matrix [][]int64) {
	for i, row := range matrix {
		fmt.Printf("Row %2d:  ", i+1)
		for _, column := range row {
			fmt.Printf("%2d ", column)
		}
		fmt.Println("")
	}
}
