package problems

import (
	"bufio"
	"io"
	"math/big"
	"os"
	"strconv"
)

func NewProblem13(filename string) IProblem {
	in, err := os.Open(filename)
	check(err)
	return Problem13{in}
}

type Problem13 struct {
	in io.Reader
}

func (p Problem13) Solve() int64 {
	sumInt := big.NewInt(0)
	scanner := bufio.NewScanner(p.in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		tempInt := new(big.Int)
		tempInt.UnmarshalText(scanner.Bytes())
		sumInt.Add(sumInt, tempInt)
	}
	result, err := strconv.ParseInt(sumInt.String()[0:10], 10, 64)
	check(err)
	return result
}
