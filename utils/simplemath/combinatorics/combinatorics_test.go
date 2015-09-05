package combinatorics_test

import (
	"github.com/tn5993/projecteuler/utils/simplemath/combinatorics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Combinatorics", func() {
	Context("Test Central Binomial Coefficient", func() {
		It("should return the follow numbers for central binomial coefficient", func() {
			Ω(combinatorics.CentralBinomialCoefficient(1)).To(Equal(int64(2)))
			Ω(combinatorics.CentralBinomialCoefficient(2)).To(Equal(int64(6)))
			Ω(combinatorics.CentralBinomialCoefficient(3)).To(Equal(int64(20)))
			Ω(combinatorics.CentralBinomialCoefficient(20)).To(Equal(int64(137846528820)))
		})
	})
})
