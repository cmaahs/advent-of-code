package main

import (
	"aocday/solution"
	"bufio"
	"fmt"
	"log"
	"os"
)

func partTwo(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

    total := 0
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		x := solution.SolveP2(line)
		total = total + x
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total: %d\n", total)

}
func partOne(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		x := solution.SolveP1(line)
		total = total + x
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total: %d\n", total)

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
