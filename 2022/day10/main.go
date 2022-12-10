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

	options := make(map[string]int)
	options["noop"] = 1
	options["addx"] = 2
	options["start"] = 0
	options["freq"] = 40

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cycleCount := 0
	runningTotal := 1
	overallTotal := 0
	crtLine := []byte(strings.ReplaceAll(fmt.Sprintf("'%-40s'", " "), " ", "."))
	scanner := bufio.NewScanner(file)
	// fmt.Println("'##..##..##..##..##..##..##..##..##..##..'")
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			fields := strings.Fields(scanner.Text())
			// fmt.Printf("%#v\n", fields)
			opValue := 0
			if len(fields) > 1 {
				opValue, _ = strconv.Atoi(fields[1])
			}
			for x := 0; x < options[fields[0]]; x++ {
				cycleCount++
				resetCycles := captureTotalP2(options, &cycleCount, &runningTotal, &overallTotal, &opValue, &crtLine)
				// fmt.Printf("%d/%d-%d\n", cycleCount, runningTotal, runningTotal+2)
				if cycleCount >= runningTotal && cycleCount <= runningTotal+2 {
					crtLine[cycleCount] = '#'
				}
				if resetCycles {
					fmt.Printf("%s\n", crtLine)
					crtLine = []byte(strings.ReplaceAll(fmt.Sprintf("'%-40s'", " "), " ", "."))
					cycleCount = 0
				}
			}
			// fmt.Printf("\tCALC: rt/val: %d / %d\n", runningTotal, opValue)
			// sprite := []byte(strings.ReplaceAll(fmt.Sprintf("'%-40s'", " "), " ", "."))
			// for x := runningTotal; x < runningTotal+3; x++ {
			// 	sprite[x] = '#'
			// }
			// fmt.Println("      : '...............###......................'")
			// fmt.Printf("SPRITE: %s\n", sprite)
			runningTotal += opValue
		}
	}

	// fmt.Printf("ANSWER: %d", overallTotal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// noop = 1 cycle
// addx = 2 cycles
// cycle start = 20
// cycle freq = 40

func countedCycle(options map[string]int, cc int) bool {

	testVal := cc - options["start"]

	if testVal%options["freq"] == 0 {
		return true
	}
	return false

}

func captureTotalP2(options map[string]int, cycleCount *int, runningTotal *int, overallTotal *int, opValue *int, crtLine *[]byte) bool {

	reset := false
	if countedCycle(options, *cycleCount) {
		// fmt.Printf("CYCLE/TOTAL: %d/%d", cycleCount, *runningTotal)
		*overallTotal += *runningTotal * *cycleCount
		// fmt.Printf("%s\n", *crtLine)
		// *crtLine = []byte(strings.ReplaceAll(fmt.Sprintf("'%-40s'", " "), " ", "."))
		reset = true
		// *runningTotal = 1
		// } else {
		// fmt.Printf("\tCALC: rt/val: %d / %d\n", *runningTotal, *opValue)
		// *runningTotal += *opValue
	}
	// sprite := []byte(strings.ReplaceAll(fmt.Sprintf("'%-40s'", " "), " ", "."))
	// for x := *runningTotal; x < *runningTotal+3; x++ {
	// 	sprite[x] = '#'
	// }
	// fmt.Printf("SPRITE: %s\n", sprite)
	// if *runningTotal >= cycleCount && *runningTotal <= cycleCount+3 {
	// 	temp := *crtLine
	// 	temp[*runningTotal] = '#'
	// 	*crtLine = temp
	// }
	return reset
}

func captureTotal(options map[string]int, cycleCount int, runningTotal *int, overallTotal *int, opValue *int) {

	if countedCycle(options, cycleCount) {
		fmt.Printf("CYCLE/TOTAL: %d/%d", cycleCount, *runningTotal)
		*overallTotal += *runningTotal * *&cycleCount
		// *runningTotal = 1
		// } else {
		// fmt.Printf("\tCALC: rt/val: %d / %d\n", *runningTotal, *opValue)
		// *runningTotal += *opValue
	}
}

func partOne(fileName string) {

	options := make(map[string]int)
	options["noop"] = 1
	options["addx"] = 2
	options["start"] = 20
	options["freq"] = 40

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cycleCount := 0
	runningTotal := 1
	overallTotal := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			fields := strings.Fields(scanner.Text())
			fmt.Printf("%#v\n", fields)
			opValue := 0
			if len(fields) > 1 {
				opValue, _ = strconv.Atoi(fields[1])
			}
			for x := 0; x < options[fields[0]]; x++ {
				cycleCount++
				captureTotal(options, cycleCount, &runningTotal, &overallTotal, &opValue)
			}
			fmt.Printf("\tCALC: rt/val: %d / %d\n", runningTotal, opValue)
			runningTotal += opValue
		}
	}

	fmt.Printf("ANSWER: %d", overallTotal)

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
