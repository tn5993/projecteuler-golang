package reader_test

import (
	"github.com/tn5993/projecteuler/utils/reader"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reader", func() {
	Context("Read in test file", func() {
		It("Convert to matrix appropriately", func() {
			filename := "files/test_reader_matrix.txt"
			result := reader.ReadFileToMatrix(filename)
			Ω(result[0]).To(Equal([]int64{3}))
			Ω(result[1]).To(Equal([]int64{7, 4}))
			Ω(result[2]).To(Equal([]int64{2, 4, 6}))
			Ω(result[3]).To(Equal([]int64{8, 5, 9, 3}))
			Ω(len(result)).To(Equal(4))
		})
	})
})
