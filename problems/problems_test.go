package problems_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tn5993/projecteuler/problems"
)

var _ = Describe("Src/Github.Com/Tn5993/Projecteuler/Problems/Problems", func() {
	Context("problem 28", func() {
		It("should return the following number for problem 28", func() {
			result := NewProblem28(1001).Solve()
			Ω(result).To(Equal(int64(669171001)))
		})
	})
})
