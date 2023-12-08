package main_test

import (
	"aocday/solution"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

//  only 12 red cubes, 13 green cubes, and 14 blue cubes

var _ = Describe("Day2", func() {

	Describe("__p1_test_1__", func() {
		Context("__p1_test_1__", func() {
			It("should be true", func() {
				Expect(solution.SolveP1("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 12, 13, 14)).To(BeTrue())
			})
		})
	})

	Describe("__p1_test_2__", func() {
		Context("__p1_test_2__", func() {
			It("should be true", func() {
				Expect(solution.SolveP1("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12, 13, 14)).To(BeTrue())
			})
		})
	})

	Describe("__p1_test_3__", func() {
		Context("__p1_test_3__", func() {
			It("should be false", func() {
				Expect(solution.SolveP1("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 12, 13, 14)).To(BeFalse())
			})
		})
	})

	Describe("__p1_test_4__", func() {
		Context("__p1_test_4__", func() {
			It("should be false", func() {
				Expect(solution.SolveP1("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 12, 13, 14)).To(BeFalse())
			})
		})
	})

	Describe("__p1_test_5__", func() {
		Context("__p1_test_5__", func() {
			It("should be true", func() {
				Expect(solution.SolveP1("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 12, 13, 14)).To(BeTrue())
			})
		})
	})

	Describe("__p2_test_1__", func() {
		Context("__p2_test_1__", func() {
			It("should be 48", func() {
				Expect(solution.SolveP2("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green")).To(Equal(48))
			})
		})
	})

	Describe("__p2_test_2__", func() {
		Context("__p2_test_2__", func() {
			It("should be 12", func() {
				Expect(solution.SolveP2("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")).To(Equal(12))
			})
		})
	})

	Describe("__p2_test_3__", func() {
		Context("__p2_test_3__", func() {
			It("should be 1560", func() {
				Expect(solution.SolveP2("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red")).To(Equal(1560))
			})
		})
	})

	Describe("__p2_test_4__", func() {
		Context("__p2_test_4__", func() {
			It("should be 630", func() {
				Expect(solution.SolveP2("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red")).To(Equal(630))
			})
		})
	})

	Describe("__p2_test_5__", func() {
		Context("__p2_test_5__", func() {
			It("should be 36", func() {
				Expect(solution.SolveP2("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green")).To(Equal(36))
			})
		})
	})

})
