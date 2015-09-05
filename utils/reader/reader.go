package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFileToMatrix(filename string) [][]int64 {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	var matrix [][]int64
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		strs := strings.Split(line, " ")
		ints := make([]int64, len(strs))
		for i := range strs {
			ints[i], err = strconv.ParseInt(strs[i], 10, 64)
			if err != nil {
				fmt.Println("Error converting from string to int")
			}
		}
		matrix = append(matrix, ints)
	}

	return matrix
}
