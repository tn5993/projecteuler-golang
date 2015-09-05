package problems

import (
	"io/ioutil"
	"strconv"
)

type Problem8 struct {
	filename string
}

func NewProblem8(filename string) Problem8 {
	return Problem8{filename}
}

func (p Problem8) Solve() int64 {
	bytes, _ := ioutil.ReadFile(p.filename)
	str := string(bytes)

	var n [1000]int64
	for i := 0; i < 1000; i++ {
		n[i], _ = strconv.ParseInt(string(str[i]), 10, 64)
	}

	var largest int64
	for i := 0; i < 988; i++ {
		var product int64 = 1
		for j := 0; j < 13; j++ {
			product *= n[i+j]
		}

		if largest < product {
			largest = product
		}
	}

	return largest
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
