package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func hasDupsP2(check []byte, num int) bool {

	mapping := make(map[byte]byte, 0)
	for _, v := range check {
		mapping[v] = v
	}

	return len(mapping) < num
}

func hasDupsP1(check []byte) bool {

	mapping := make(map[byte]byte, 0)
	for _, v := range check {
		mapping[v] = v
	}

	return len(mapping) < 4
}

func scanTextP1(text []byte) int {

	size := len(text)
	lastByte := 0

	for x := 3; x < size; x++ {
		check := text[x-3 : x+1]
		fmt.Printf("Check: %s\n", check)
		if !hasDupsP1(check) {
			return x + 1
		}
	}

	return lastByte
}

func scanTextP2(text []byte) int {

	size := len(text)
	lastByte := 0

	for x := 13; x < size; x++ {
		check := text[x-13 : x+1]
		fmt.Printf("Check: %s\n", check)
		if !hasDupsP2(check, 14) {
			return x + 1
		}
	}

	return lastByte
}

func partTwo(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// totalScore := 0
	for scanner.Scan() {
		processText := scanner.Bytes()
		fmt.Printf("Processing: %s\n", processText)

		x := scanTextP2(processText)
		fmt.Printf("Start = %d\n", x)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Total Score: %d", totalScore)

}
func partOne(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// totalScore := 0
	for scanner.Scan() {
		processText := scanner.Bytes()
		fmt.Printf("Processing: %s\n", processText)

		x := scanTextP1(processText)
		fmt.Printf("Start = %d\n", x)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Total Score: %d", totalScore)

}

func main() {

	dataFile := "sample1.txt"

	if len(os.Args) > 2 {
		dataFile = os.Args[2]
	}

	switch os.Args[1] {
	case "p1":
		partOne(dataFile)
	case "p2":
		partTwo(dataFile)
	}
}
