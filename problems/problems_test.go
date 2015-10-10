package problems_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tn5993/projecteuler/problems"
)

var _ = Describe("Src/Github.Com/Tn5993/Projecteuler/Problems/Problems", func() {
	Context("problem 28", func() {
		FIt("Should return the following expection", func() {
			p := problems.NewProblem28()
			Ω(p.Solve()).To(Equal(int64(10)))
		})
	})
})
