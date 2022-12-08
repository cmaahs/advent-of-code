package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func partTwo(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		currentFloor := 0
		fmt.Println("---")
		for i, v := range scanner.Bytes() {
			fmt.Printf("%#v : %s\n", i, string(v))
			switch string(v) {
			case "(":
				currentFloor++
			case ")":
				currentFloor--
			}
			fmt.Printf("FL: %d\n", currentFloor)
			if currentFloor < 0 {
				fmt.Printf("Position: %d\n", i+1)
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
func partOne(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		up := strings.Count(scanner.Text(), "(")
		down := strings.Count(scanner.Text(), ")")
		fmt.Printf("Current Floor: %d\n", up-down)
	}

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
