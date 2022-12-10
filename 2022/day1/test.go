package main

import (
	"fmt"
	"strings"
)

func test() {

	myStr := "1 2 3 4 5"

	splits := strings.Split(myStr, " ")

	fmt.Printf("second: %s\n", splits[1])

}
