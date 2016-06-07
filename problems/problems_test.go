package problems_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tn5993/projecteuler/problems"
)

var _ = Describe("Src/Github.Com/Tn5993/Projecteuler/Problems/Problems", func() {
	Context("problem 28", func() {
		It("should return the following number for problem 28", func() {
			result, err := SolveProblem28(1001)
			Ω(err).ToNot(HaveOccurred())
			Ω(result).To(Equal(669171001))
		})
	})
})
