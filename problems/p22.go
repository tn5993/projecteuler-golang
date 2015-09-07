package problems

import (
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func NewProblem22(filename string) IProblem {
	in, err := os.Open(filename)
	check(err)
	return Problem22{in}
}

type Problem22 struct {
	in io.Reader
}

func (p Problem22) Solve() int64 {
	text, err := ioutil.ReadAll(p.in)
	check(err)
	names := strings.Split(strings.Replace(string(text), "\"", "", -1), ",")
	sort.Strings(names)

	var sum int64
	for i, name := range names {
		sum += alphabetValue(name) * int64(i+1)
	}

	return sum
}

func alphabetValue(s string) int64 {
	var sum int64
	for _, v := range s {
		sum += int64(v) - 64
	}
	return sum
}
