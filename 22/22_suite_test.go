package day22_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test22(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "22 Suite")
}