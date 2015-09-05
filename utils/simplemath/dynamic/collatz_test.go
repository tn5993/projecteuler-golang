package dynamic_test

import (
	"github.com/tn5993/projecteuler/utils/simplemath/dynamic"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Collatz", func() {
	Context("Find longest collatz sequence", func() {
		It("For the number 13, we have 10 terms (13->40->20->10->5->16->8->4->2->1)", func() {
			Ω(dynamic.FindLongestCollatzSequence(14)).To(Equal(int64(9)))
		})

		It("For the number 1, it should be 1", func() {
			Ω(dynamic.FindLongestCollatzSequence(1)).To(Equal(int64(1)))
		})

		It("For 1000000, it should be 837799", func() {
			Ω(dynamic.FindLongestCollatzSequence(1000000)).To(Equal(int64(837799)))
		})
	})
})
