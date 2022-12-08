package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type ElfWeight struct {
	Elf    int
	Weight int
}

func partTwo(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elf := 1
	var highElves []ElfWeight
	currentValue := 0
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			fmt.Printf("Elf %d carries %d\n", elf, currentValue)
			newElf := ElfWeight{
				Elf:    elf,
				Weight: currentValue,
			}
			highElves = append(highElves, newElf)
			currentValue = 0
			elf = elf + 1
		} else {
			intVar, ierr := strconv.Atoi(scanner.Text())
			if ierr != nil {
				fmt.Println("ERROR: bad int conversion")
				os.Exit(1)
			}
			currentValue = currentValue + intVar
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(highElves, func(i, j int) bool {
		return highElves[i].Weight < highElves[j].Weight
	})

	totalWeight := 0
	totalWeight = totalWeight + highElves[len(highElves)-1].Weight
	totalWeight = totalWeight + highElves[len(highElves)-2].Weight
	totalWeight = totalWeight + highElves[len(highElves)-3].Weight

	fmt.Printf("Total Weight for 3 Top Elves: %d", totalWeight)

}

func partOne(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elf := 1
	highestValue := 0
	highestElf := 0
	currentValue := 0
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			if currentValue > highestValue {
				highestElf = elf
				highestValue = currentValue
			}
			currentValue = 0
			elf = elf + 1
		} else {
			intVar, ierr := strconv.Atoi(scanner.Text())
			if ierr != nil {
				fmt.Println("ERROR: bad int conversion")
				os.Exit(1)
			}
			currentValue = currentValue + intVar
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Elf: %d, Carries: %d\n", highestElf, highestValue)

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
