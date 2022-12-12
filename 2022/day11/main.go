package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type MonkeyList map[int]Monkey

type Monkey struct {
	StartingItems     map[int]uint64
	Operation         string
	OperationModifier uint64
	DivisibleBy       uint64
	TrueMonkey        int
	FalseMonkey       int
	Inspections       int
}

func clearList(list map[int]uint64) {
	for k := range list {
		delete(list, k)
	}
}

func partTwo(fileName string) {

	monkeyList := MonkeyList{}

	// monkeyItems := make(map[int]map[int]int)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	newMonkey := Monkey{
		StartingItems:     make(map[int]uint64),
		Operation:         "",
		OperationModifier: 0,
		DivisibleBy:       0,
		TrueMonkey:        -1,
		FalseMonkey:       -1,
		Inspections:       0,
	}

	num := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			fmt.Printf("Adding Monkey: %d, %#v\n", num, newMonkey)
			monkeyList.setupMonkey(num, newMonkey)
			num++
			newMonkey = Monkey{
				StartingItems:     make(map[int]uint64),
				Operation:         "",
				OperationModifier: 0,
				DivisibleBy:       0,
				TrueMonkey:        -1,
				FalseMonkey:       -1,
				Inspections:       0,
			}
		}
		if strings.HasPrefix(scanner.Text(), "  Starting items:") {
			splits := strings.Split(scanner.Text(), ":")
			items := strings.Split(splits[1], ",")
			for x := 0; x < len(items); x++ {
				intVal, _ := strconv.Atoi(strings.TrimSpace(items[x]))
				fmt.Printf("Adding Item: %d/%d=%d\n", num, x, intVal)
				newMonkey.StartingItems[x] = uint64(intVal)
			}
		}
		if strings.HasPrefix(scanner.Text(), "  Operation:") {
			splits := strings.Split(scanner.Text(), ":")
			items := strings.Split(splits[1], "=")
			ops := strings.Fields(items[1])
			newMonkey.Operation = ops[1]
			if ops[2] == "old" {
				newMonkey.OperationModifier = 0
			} else {
				intVal, _ := strconv.Atoi(ops[2])
				newMonkey.OperationModifier = uint64(intVal)
			}
		}
		if strings.HasPrefix(scanner.Text(), "  Test:") {
			splits := strings.Split(scanner.Text(), ":")
			items := strings.Fields(splits[1])
			intVal, _ := strconv.Atoi(items[2])
			newMonkey.DivisibleBy = uint64(intVal)
		}
		if strings.HasPrefix(scanner.Text(), "    If true:") {
			splits := strings.Split(scanner.Text(), ":")
			items := strings.Fields(splits[1])
			intVal, _ := strconv.Atoi(items[3])
			newMonkey.TrueMonkey = intVal
		}
		if strings.HasPrefix(scanner.Text(), "    If false:") {
			splits := strings.Split(scanner.Text(), ":")
			items := strings.Fields(splits[1])
			intVal, _ := strconv.Atoi(items[3])
			newMonkey.FalseMonkey = intVal
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	productDivisor := uint64(1)

	for _, v := range monkeyList {
		productDivisor *= v.DivisibleBy
	}
	fmt.Printf("ProductDivisor: %d", productDivisor)
	for r := 0; r < 10000; r++ {
		fmt.Printf("---- ROUND %d ----\n", r)
		for x := 0; x < len(monkeyList); x++ {
			hadInspections := monkeyList[x].Inspections
			monkeyList.processMonkey(x, productDivisor, false)
			fmt.Printf("Inspections: %d: %d -> %d\n", x, hadInspections, monkeyList[x].Inspections)
		}
		// fmt.Printf("%#v\n", monkeyList)
	}

	inspections := []int{}
	for i, v := range monkeyList {
		inspections = append(inspections, v.Inspections)
		fmt.Printf("Monkey: %d, Inspections: %d\n", i, v.Inspections)
	}

	sort.Ints(inspections)

	fmt.Printf("ANSWER: %d", inspections[len(inspections)-1]*inspections[len(inspections)-2])
}

func (m MonkeyList) setupMonkey(num int, monkey Monkey) {
	m[num] = monkey
}

func (m MonkeyList) processMonkey(num int, productDivisor uint64, modify bool) {

	inspections := 0
	for x := 0; x < len(m[num].StartingItems); x++ {
		logLine := ""
		inspections++
		// fmt.Printf("Item: %d=%d\n", x, m[num].StartingItems[x])
		operationMod := m[num].OperationModifier
		if operationMod == 0 {
			operationMod = m[num].StartingItems[x]
		}
		worryLevel := uint64(0)
		switch m[num].Operation {
		case "*":
			worryLevel = m[num].StartingItems[x] * operationMod
			logLine = fmt.Sprintf("%s WL: %d * %d -> %d", logLine, m[num].StartingItems[x], operationMod, worryLevel)
		case "+":
			worryLevel = m[num].StartingItems[x] + operationMod
			logLine = fmt.Sprintf("%s WL: %d * %d -> %d", logLine, m[num].StartingItems[x], operationMod, worryLevel)
		}
		modifiedWorryLevel := worryLevel % productDivisor
		if modify {
			modifiedWorryLevel = uint64(math.Round(float64(worryLevel / 3)))
		}
		remainder := modifiedWorryLevel % m[num].DivisibleBy
		logLine = fmt.Sprintf("%s R: %d", logLine, remainder)
		if remainder == 0 {
			// to truemonkey
			targetMonkey := m[num].TrueMonkey
			fmt.Printf("\tTT: %d,%s, I: %d=%d\n", targetMonkey, logLine, x, worryLevel)
			targetMonkeyNext := len(m[targetMonkey].StartingItems)
			m[targetMonkey].StartingItems[targetMonkeyNext] = modifiedWorryLevel
		} else {
			// to falsemonkey
			targetMonkey := m[num].FalseMonkey
			fmt.Printf("\tTF: %d,%s I: %d=%d\n", targetMonkey, logLine, x, worryLevel)
			targetMonkeyNext := len(m[targetMonkey].StartingItems)
			m[targetMonkey].StartingItems[targetMonkeyNext] = modifiedWorryLevel
		}
	}
	clearList(m[num].StartingItems)
	entry := m[num]
	entry.Inspections += inspections
	m[num] = entry
}

func partOne(fileName string) {

	monkeyList := MonkeyList{}

	// monkeyItems := make(map[int]map[int]int)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	newMonkey := Monkey{
		StartingItems:     make(map[int]uint64),
		Operation:         "",
		OperationModifier: 0,
		DivisibleBy:       0,
		TrueMonkey:        -1,
		FalseMonkey:       -1,
		Inspections:       0,
	}

	num := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			fmt.Printf("Adding Monkey: %d, %#v\n", num, newMonkey)
			monkeyList.setupMonkey(num, newMonkey)
			num++
			newMonkey = Monkey{
				StartingItems:     make(map[int]uint64),
				Operation:         "",
				OperationModifier: 0,
				DivisibleBy:       0,
				TrueMonkey:        -1,
				FalseMonkey:       -1,
				Inspections:       0,
			}
		}
		if strings.HasPrefix(scanner.Text(), "  Starting items:") {
			splits := strings.Split(scanner.Text(), ":")
			items := strings.Split(splits[1], ",")
			for x := 0; x < len(items); x++ {
				intVal, _ := strconv.Atoi(strings.TrimSpace(items[x]))
				fmt.Printf("Adding Item: %d/%d=%d\n", num, x, intVal)
				newMonkey.StartingItems[x] = uint64(intVal)
			}
		}
		if strings.HasPrefix(scanner.Text(), "  Operation:") {
			splits := strings.Split(scanner.Text(), ":")
			items := strings.Split(splits[1], "=")
			ops := strings.Fields(items[1])
			newMonkey.Operation = ops[1]
			if ops[2] == "old" {
				newMonkey.OperationModifier = 0
			} else {
				intVal, _ := strconv.Atoi(ops[2])
				newMonkey.OperationModifier = uint64(intVal)
			}
		}
		if strings.HasPrefix(scanner.Text(), "  Test:") {
			splits := strings.Split(scanner.Text(), ":")
			items := strings.Fields(splits[1])
			intVal, _ := strconv.Atoi(items[2])
			newMonkey.DivisibleBy = uint64(intVal)
		}
		if strings.HasPrefix(scanner.Text(), "    If true:") {
			splits := strings.Split(scanner.Text(), ":")
			items := strings.Fields(splits[1])
			intVal, _ := strconv.Atoi(items[3])
			newMonkey.TrueMonkey = intVal
		}
		if strings.HasPrefix(scanner.Text(), "    If false:") {
			splits := strings.Split(scanner.Text(), ":")
			items := strings.Fields(splits[1])
			intVal, _ := strconv.Atoi(items[3])
			newMonkey.FalseMonkey = intVal
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for r := 0; r < 20; r++ {
		for x := 0; x < len(monkeyList); x++ {
			monkeyList.processMonkey(x, 3, true)
		}
		fmt.Printf("%#v\n", monkeyList)
	}

	inspections := []int{}
	for _, v := range monkeyList {
		inspections = append(inspections, v.Inspections)
	}

	sort.Ints(inspections)

	fmt.Printf("ANSWER: %d", inspections[len(inspections)-1]*inspections[len(inspections)-2])

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
