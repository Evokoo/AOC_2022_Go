package day04_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test04(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "04 Suite")
}
