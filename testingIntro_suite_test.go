package testingIntro_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestTestingIntro(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestingIntro Suite")
}
