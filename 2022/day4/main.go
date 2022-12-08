package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func partTwo(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalScore := 0
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			splits := strings.Split(scanner.Text(), ",")
			firstSection := strings.Split(splits[0], "-")
			startFirstSection, _ := strconv.Atoi(firstSection[0])
			endFirstSection, _ := strconv.Atoi(firstSection[1])

			secondSection := strings.Split(splits[1], "-")
			startSecondSection, _ := strconv.Atoi(secondSection[0])
			endSecondSection, _ := strconv.Atoi(secondSection[1])

			overlaps := false
			if startFirstSection >= startSecondSection || endFirstSection >= startSecondSection {
				if startFirstSection <= endSecondSection || endFirstSection <= endSecondSection {
					overlaps = true
				}
			}
			if startSecondSection >= startFirstSection || endSecondSection >= startFirstSection {
				if startSecondSection <= endFirstSection || endSecondSection <= endFirstSection {
					overlaps = true
				}
			}
			if overlaps {
				totalScore++
			}
		}
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
		if len(scanner.Text()) > 0 {
			splits := strings.Split(scanner.Text(), ",")
			firstSection := strings.Split(splits[0], "-")
			startFirstSection, _ := strconv.Atoi(firstSection[0])
			endFirstSection, _ := strconv.Atoi(firstSection[1])

			secondSection := strings.Split(splits[1], "-")
			startSecondSection, _ := strconv.Atoi(secondSection[0])
			endSecondSection, _ := strconv.Atoi(secondSection[1])

			overlaps := false
			if startFirstSection <= startSecondSection {
				if endFirstSection >= endSecondSection {
					overlaps = true
				}
			}
			if startSecondSection <= startFirstSection {
				if endSecondSection >= endFirstSection {
					overlaps = true
				}
			}
			if overlaps {
				totalScore++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Score: %d", totalScore)

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
