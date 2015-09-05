package intmath_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tn5993/projecteuler/utils/simplemath/intmath"
	"math/big"
)

var _ = Describe("Intmath", func() {
	Context("Test Factorial", func() {
		It("Should throw an error if negative number is given", func() {
			result, err := Factorial(-1)
			Ω(result).To(Equal(int64(-1)))
			Ω(err).To(HaveOccurred())
		})

		It("Should return 1 for 0", func() {
			Ω(Factorial(0)).To(Equal(int64(1)))
		})

		It("Should return appropriate values for the following numbers", func() {
			Ω(Factorial(1)).To(Equal(int64(1)))
			Ω(Factorial(10)).To(Equal(int64(3628800)))
			Ω(Factorial(20)).To(Equal(int64(2432902008176640000)))
		})
	})

	Context("Test absolute int value", func() {
		It("Should return correct absolute value", func() {
			Ω(Abs(-5)).To(Equal(5))
			Ω(Abs(-1)).To(Equal(1))
			Ω(Abs(0)).To(Equal(0))
			Ω(Abs(5)).To(Equal(5))
			Ω(Abs(1)).To(Equal(1))
		})
	})

	Context("Text big factorial", func() {
		It("Should throw an error for value less than 0", func() {
			res, err := BigFactorial(-3)
			Ω(err).To(HaveOccurred())
			Ω(res.Cmp(big.NewInt(-1))).To(Equal(0))
		})

		It("Should return 1 for 0", func() {
			res, err := BigFactorial(0)
			Ω(err).ToNot(HaveOccurred())
			Ω(res.Cmp(big.NewInt(1))).To(Equal(0))
		})

		It("Should return appropriate value for the following numbers", func() {
			res, err := BigFactorial(1)
			Ω(res.Cmp(big.NewInt(1))).To(Equal(0))
			Ω(err).ToNot(HaveOccurred())

			res, err = BigFactorial(10)
			Ω(res.Cmp(big.NewInt(3628800))).To(Equal(0))
			Ω(err).ToNot(HaveOccurred())

			res, err = BigFactorial(20)
			Ω(res.Cmp(big.NewInt(2432902008176640000))).To(Equal(0))
			Ω(err).ToNot(HaveOccurred())
		})
	})

	Context("Test SumDigit", func() {
		It("Should return the sum of all digit in a number", func() {
			Ω(SumDigit(big.NewInt(999999)).Int64()).To(Equal(int64(54)))
			Ω(SumDigit(big.NewInt(42566)).Int64()).To(Equal(int64(23)))
		})
	})

})
