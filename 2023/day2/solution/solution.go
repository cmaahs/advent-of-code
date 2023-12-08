package solution

import (
	"strconv"
	"strings"
	// "sort"
	// "strconv"
	// "strings"
	// "unicode"
)

type (
	words struct {
		Name   string
		Number string
	}
)

func SolveP2(in string) int {

	gameLine := strings.TrimSpace(strings.SplitN(in, ":", 2)[1])
	adjusted := strings.ReplaceAll(gameLine, ";", ",")
	splits := strings.Split(adjusted, ",")

	highR := 0
	highG := 0
	highB := 0
	for _, v := range splits {
		game := strings.Split(strings.TrimSpace(v), " ")
		switch game[1] {
		case "red":
			if toI(game[0]) > highR {
				highR = toI(game[0])
			}
		case "green":
			if toI(game[0]) > highG {
				highG = toI(game[0])
			}
		case "blue":
			if toI(game[0]) > highB {
				highB = toI(game[0])
			}
		}
	}

	return highR * highG * highB

}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
//
//	only 12 red cubes, 13 green cubes, and 14 blue cubes
func SolveP1(in string, r, g, b int) bool {

	gameLine := strings.TrimSpace(strings.SplitN(in, ":", 2)[1])
	adjusted := strings.ReplaceAll(gameLine, ";", ",")
	splits := strings.Split(adjusted, ",")

	for _, v := range splits {
		game := strings.Split(strings.TrimSpace(v), " ")
		switch game[1] {
		case "red":
			if toI(game[0]) > r {
				return false
			}
		case "green":
			if toI(game[0]) > g {
				return false
			}
		case "blue":
			if toI(game[0]) > b {
				return false
			}
		}
	}

	return true
}

func toI(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return i
}
