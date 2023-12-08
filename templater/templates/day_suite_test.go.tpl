package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDay{{.Day}}(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Day{{.Day}} Suite")
}
