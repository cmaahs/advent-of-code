package main_test

import (
	"aocday/solution"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day{{.Day}}", func() {

	Describe("__p1_test_1__", func() {
		Context("__p1_test_1__", func() {
			It("should be __VAL__", func() {
				Expect(solution.SolveP1("__PARAM__")).To(Equal(0))
			})
		})
	})

	Describe("__p2_test_1__", func() {
		Context("__p2_test_1__", func() {
			It("should be __VAL__", func() {
				Expect(solution.SolveP2("__PARAM__")).To(Equal(0))
			})
		})
	})

})
