package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func extractSameP1(first string, second string) int {
	// 65 = A
	// 97 = a
	for _, v := range first {
		if strings.Contains(second, string(v)) {
			if int(v) > 96 {
				return int(v) - 96
			} else {
				return int(v) - 38
			}
		}
	}
	return 0
}

func extractSameP2(lines []string) int {

	// 65 = A
	// 97 = a
	for _, v := range lines[0] {
		if strings.Contains(lines[1], string(v)) {
			if strings.Contains(lines[2], string(v)) {
				if int(v) > 96 {
					return int(v) - 96
				} else {
					return int(v) - 38
				}
			}
		}
	}
	return 0
}

func partTwo(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalScore := 0
	line := 1
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if line%3 == 0 {
			priority := extractSameP2(lines)
			lines = lines[:0]
			totalScore = totalScore + priority
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Score: %d", totalScore)

}
func partOne(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalScore := 0
	for scanner.Scan() {
		x := len(scanner.Text())
		half := x / 2
		firstCompartment := scanner.Text()[0:half]
		secondCompartment := scanner.Text()[half:]
		fmt.Printf("%s = %s : %s\n", scanner.Text(), firstCompartment, secondCompartment)
		priority := extractSameP1(firstCompartment, secondCompartment)
		totalScore = totalScore + priority
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Score: %d", totalScore)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

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
