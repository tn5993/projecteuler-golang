package intmath_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIntmath(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Intmath Suite")
}
