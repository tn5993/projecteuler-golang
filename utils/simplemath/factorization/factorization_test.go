package factorization_test

import (
	factor "github.com/tn5993/projecteuler/utils/simplemath/factorization"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factorization", func() {
	Describe("Test Factization Package", func() {
		Context("Number of divisors of 28", func() {
			It("should be 6 (1,2,4,7,14,28)", func() {
				Expect(factor.DivisorsOf(28)).Should(Equal(int64(6)))
			})
		})

		Context("Divisors of 21, 15, 10, 6", func() {
			It("should be 4", func() {
				Ω(factor.DivisorsOf(21)).To(Equal(int64(4)))
				Ω(factor.DivisorsOf(15)).To(Equal(int64(4)))
				Ω(factor.DivisorsOf(10)).To(Equal(int64(4)))
				Ω(factor.DivisorsOf(6)).To(Equal(int64(4)))
				Ω(factor.DivisorsOf(3)).To(Equal(int64(2)))
				Ω(factor.DivisorsOf(1)).To(Equal(int64(1)))

			})
		})
	})
})
