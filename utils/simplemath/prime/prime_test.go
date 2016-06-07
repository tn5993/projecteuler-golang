package prime_test

import (
	"github.com/tn5993/projecteuler/utils/simplemath/prime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Prime", func() {
	Describe("Test Sieve Of E for finding prime numbers less than or equal n", func() {
		Context("If n <= 1", func() {
			It("should return nil", func() {
				Expect(prime.GetPrimesUnder(1)).To(Equal([]int64{}))
			})
		})

		Context("If n == 20", func() {
			It("should return 2,3,5,7,11,13,17,19", func() {
				Expected := []int64{2, 3, 5, 7, 11, 13, 17, 19}
				Expect(prime.GetPrimesUnder(20)).To(Equal(Expected))
			})
		})

		Context("If n == 7 (a prime number at boundary)", func() {
			It("should include 7 in result ", func() {
				Expected := []int64{2, 3, 5, 7}
				Ω(prime.GetPrimesUnder(7)).To(Equal(Expected))
			})
		})
	})

	Describe("IsPrime", func() {
		Context("For n is prime or not prime", func() {
			It("correctly point out", func() {
				Ω(prime.IsPrime(2)).To(BeTrue())
				Ω(prime.IsPrime(3)).To(BeTrue())
				Ω(prime.IsPrime(324292)).To(BeFalse())
				Ω(prime.IsPrime(433)).To(BeTrue())
				Ω(prime.IsPrime(429)).To(BeFalse())
				Ω(prime.IsPrime(-4)).To(BeFalse())
				Ω(prime.IsPrime(0)).To(BeFalse())
				Ω(prime.IsPrime(1)).To(BeFalse())
			})
		})
	})
})
