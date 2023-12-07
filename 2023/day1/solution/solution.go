package solution

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type (
	words struct {
		Name   string
		Number string
	}
)

func replaceWords(in string) string {
	out := in
	list := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	convert := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	order := map[int]words{}
	for _, v := range list {
		// x := strings.Index(in, v)
		// fmt.Printf("%s found at loc: %d\n", v, x)
		remaining := in
		indexer := 0
		for {
			if i := strings.Index(remaining, v); i > -1 {
				// fmt.Printf("Adding - %s at %d\n", v, i)
				order[i+indexer] = words{
					Name:   v,
					Number: convert[v],
				}
				remaining = strings.Replace(remaining, v, convert[v], 1)
				indexer = indexer + i
			} else {
				break
			}
		}
	}

	for k, v := range order {
		fmt.Printf("key: %d, value: %#v\n", k, v)
	}
	// Extract keys from map
	keys := make([]int, 0, len(order))
	for k := range order {
		keys = append(keys, k)
	}

	// Sort keys
	sort.Ints(keys)

	// if len(keys) > 0 {
	for i := range keys {
		fmt.Printf("Replacing - %d - %#v\n", keys[i], order[keys[i]])
		out = strings.Replace(out, order[keys[i]].Name, order[keys[i]].Number, 1)
		// lastIndex := len(keys) - 1
		// firstIndex := 0
		// fmt.Printf("Replacing - %d - %#v\n", keys[firstIndex], order[keys[firstIndex]])
		// out = strings.Replace(out, order[keys[firstIndex]].Name, order[keys[firstIndex]].Number, 1)
		// fmt.Printf("Replacing - %d - %#v\n", keys[lastIndex], order[keys[lastIndex]])
		// out = strings.ReplaceAll(out, order[keys[lastIndex]].Name, order[keys[lastIndex]].Number)
	}
	// }

	return out

}

func LineScoreP2(in string) int {

	fmt.Printf("sample in %#v\n", in)
	workstring := replaceWords(in)
	fmt.Printf("workstring %#v\n", workstring)
	// Convert string to a slice of runes
	runes := []rune(workstring)
	gotFirst := false
	first := '0'
	last := ' '
	// Loop over individual characters
	for i := 0; i < len(runes); i++ {
		if unicode.IsNumber(runes[i]) {
			if gotFirst {
				last = runes[i]
			} else {
				first = runes[i]
				last = runes[i]
				gotFirst = true
			}

		}

	}
	num := fmt.Sprintf("%c%c", first, last)
	i, err := strconv.Atoi(strings.TrimSpace(num))
	if err != nil {
		panic(err)
	}

	fmt.Printf("line score: %d\n\n", i)
	return i
}
func LineScore(in string) int {

	// Convert string to a slice of runes
	runes := []rune(in)

	gotFirst := false
	first := '0'
	last := ' '
	// Loop over individual characters
	for i := 0; i < len(runes); i++ {
		if unicode.IsNumber(runes[i]) {
			if gotFirst {
				last = runes[i]
			} else {
				first = runes[i]
				last = runes[i]
				gotFirst = true
			}

		}

	}
	num := fmt.Sprintf("%c%c", first, last)
	i, err := strconv.Atoi(strings.TrimSpace(num))
	if err != nil {
		panic(err)
	}

	return i
}
