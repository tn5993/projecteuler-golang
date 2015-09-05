package combinatorics_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCombinatorics(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Combinatorics Suite")
}
