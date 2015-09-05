package dynamic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDynamic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dynamic Suite")
}
