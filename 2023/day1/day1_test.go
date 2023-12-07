package main_test

import (
	"aocday/solution"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Day1", func() {
	Describe("line1", func() {
		Context("first line", func() {
			It("should be 12", func() {
				Expect(solution.LineScore("1abc2")).To(Equal(12))
			})
		})
	})

	Describe("line2", func() {
		Context("second line", func() {
			It("should be 38", func() {
				Expect(solution.LineScore("pqr3stu8vwx")).To(Equal(38))
			})
		})
	})

	Describe("line3", func() {
		Context("third line", func() {
			It("should be 25", func() {
				Expect(solution.LineScore("a1b2c3d4e5f")).To(Equal(15))
			})
		})
	})

	Describe("line4", func() {
		Context("third line", func() {
			It("should be 7", func() {
				Expect(solution.LineScore("treb7uchet")).To(Equal(77))
			})
		})
	})

	Describe("word1", func() {
		Context("first word line", func() {
			It("should be 29", func() {
				Expect(solution.LineScoreP2("two1nine")).To(Equal(29))
			})
		})
	})

	Describe("word2", func() {
		Context("second word line", func() {
			It("should be 83", func() {
				Expect(solution.LineScoreP2("eightwothree")).To(Equal(83))
			})
		})
	})

	Describe("word3", func() {
		Context("third word line", func() {
			It("should be 13", func() {
				Expect(solution.LineScoreP2("abcone2threexyz")).To(Equal(13))
			})
		})
	})
	Describe("word4", func() {
		Context("forth word line", func() {
			It("should be 24", func() {
				Expect(solution.LineScoreP2("xtwone3four")).To(Equal(24))
			})
		})
	})
	Describe("word5", func() {
		Context("fifth word line", func() {
			It("should be 42", func() {
				Expect(solution.LineScoreP2("4nineeightseven2")).To(Equal(42))
			})
		})
	})
	Describe("word6", func() {
		Context("sixth word line", func() {
			It("should be 14", func() {
				Expect(solution.LineScoreP2("zoneight234")).To(Equal(14))
			})
		})
	})
	Describe("word7", func() {
		Context("seventh word line", func() {
			It("should be 76", func() {
				Expect(solution.LineScoreP2("7pqrstsixteen")).To(Equal(76))
			})
		})
	})

})
