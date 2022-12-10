package main

import (
	"fmt"
	"strings"
)

func main() {

	// myStr := "1 2 3 4 5"
	myStr := "1    2     3             4                   5"

	splits := strings.Split(myStr, " ")
	splitsF := strings.Fields(myStr)

	for x := 0; x < 5; x++ {
		fmt.Printf("Split %d: %s\n", x, splits[x])
		fmt.Printf("Field %d: %s\n", x, splitsF[x])
	}

}
