package day24_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test24(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "24 Suite")
}