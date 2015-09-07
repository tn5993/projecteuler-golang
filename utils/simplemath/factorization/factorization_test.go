package factorization_test

import (
	. "github.com/tn5993/projecteuler/utils/simplemath/factorization"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Factorization", func() {
	Describe("Test Factization Package", func() {
		Context("Number of divisors of 28", func() {
			It("should be 6 (1,2,4,7,14,28)", func() {
				Ω(DivisorsCount(28)).Should(Equal(int64(6)))
			})
		})

		Context("Number of divisors of 21, 15, 10, 6", func() {
			It("should be the following", func() {
				Ω(DivisorsCount(21)).To(Equal(int64(4)))
				Ω(DivisorsCount(15)).To(Equal(int64(4)))
				Ω(DivisorsCount(10)).To(Equal(int64(4)))
				Ω(DivisorsCount(6)).To(Equal(int64(4)))
				Ω(DivisorsCount(3)).To(Equal(int64(2)))
				Ω(DivisorsCount(1)).To(Equal(int64(1)))
			})
		})

		Context("Divisor Sum of 1, 3, 6, 10, 15, 21, 28", func() {
			It("should be the following", func() {
				Ω(DivisorSum(1)).To(Equal(int64(1)))
				Ω(DivisorSum(3)).To(Equal(int64(1)))
				Ω(DivisorSum(6)).To(Equal(int64(6)))
				Ω(DivisorSum(10)).To(Equal(int64(8)))
				Ω(DivisorSum(15)).To(Equal(int64(9)))
				Ω(DivisorSum(21)).To(Equal(int64(11)))
				Ω(DivisorSum(28)).To(Equal(int64(28)))
				Ω(DivisorSum(120)).To(Equal(int64(240)))
				Ω(DivisorSum(284)).To(Equal(int64(220)))
				Ω(DivisorSum(220)).To(Equal(int64(284)))
			})
		})
	})
})
