package prime_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPrime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prime Suite")
}
