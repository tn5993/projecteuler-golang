package problems_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProjectEuler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Problems Suite")
}
