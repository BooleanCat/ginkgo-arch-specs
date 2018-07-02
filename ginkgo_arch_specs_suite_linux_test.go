package ginkgo_arch_specs_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgoArchSpecs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoArchSpecs Suite")
}
