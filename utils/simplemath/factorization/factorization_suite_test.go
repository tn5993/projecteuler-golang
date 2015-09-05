package factorization_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFactorization(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Factorization Suite")
}
