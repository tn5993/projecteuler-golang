package main

import (
	"flag"
	"fmt"
	probs "github.com/tn5993/projecteuler/problems"
	"time"
)

func main() {
	var probNum int
	flag.IntVar(&probNum, "problem", 9, "a problem number")
	flag.Parse()
	m := map[int]probs.IProblem{
		8:  probs.NewProblem8("files/p8.txt"),
		9:  probs.NewProblem9(),
		10: probs.NewProblem10(2000000),
		11: probs.NewProblem11("files/p11.txt"),
		12: probs.NewProblem12(),
		13: probs.NewProblem13("files/p13.txt"),
		14: probs.NewProblem14(1000000),
		15: probs.NewProblem15(20),
		16: probs.NewProblem16(),
		17: probs.NewProblem17(),
		18: probs.NewProblem18("files/p18.txt"),
		19: probs.NewProblem19(),
		20: probs.NewProblem20(100),
		21: probs.NewProblem21(10000),
	}
	fmt.Printf("=============\n| Problem %d |\n=============\n", probNum)
	solveProblem(m[probNum])
}

func solveProblem(problem probs.IProblem) {
	start := time.Now()
	fmt.Printf("The answer is %d\n", problem.Solve())
	elapsed := time.Since(start)
	fmt.Printf("Solution took %s\n", elapsed)
}
