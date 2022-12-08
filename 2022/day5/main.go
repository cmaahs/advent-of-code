package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getStackCount(fileName string) int {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), " 1 ") {
			stacks := strings.Fields(scanner.Text())
			fmt.Printf("Stacks: %d\n", len(stacks))
			return len(stacks)
		}
	}
	return 0

}

func removeP2(slice []string, s int) []string {
	// return append(slice[:s], slice[s+1:]...)
	var newArray []string

	for x := 0; x < len(slice)-s; x++ {
		newArray = append(newArray, slice[x])
	}
	return newArray
}

func injectP2(slice []string, new string) []string {

	var newArray []string
	newArray = append(newArray, new)
	for _, v := range slice {
		newArray = append(newArray, v)
	}
	return newArray
}
func removeP1(slice []string, s int) []string {
	// return append(slice[:s], slice[s+1:]...)
	var newArray []string

	for x := 0; x < len(slice)-s; x++ {
		newArray = append(newArray, slice[x])
	}
	return newArray
}

func injectP1(slice []string, new string) []string {

	var newArray []string
	newArray = append(newArray, new)
	for _, v := range slice {
		newArray = append(newArray, v)
	}
	return newArray
}
func partTwo(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	stackNum := getStackCount(fileName)
	stacks := make(map[int][]string, 0)

	re := regexp.MustCompile("[0-9]+")
	// totalScore := 0
	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "[") {
			tmp := []byte(scanner.Text())
			for i := 0; i < stackNum; i++ {
				item := strings.TrimSpace(string(tmp[(4 * i):((4 * i) + 3)]))
				fmt.Println(item)
				if strings.Contains(item, "[") {
					stacks[i] = injectP2(stacks[i], item)
				}
			}
		} else {
			if strings.HasPrefix(scanner.Text(), "move") {
				matches := re.FindAllString(scanner.Text(), -1)
				numMoves, _ := strconv.Atoi(matches[0])
				fromStack, _ := strconv.Atoi(matches[1])
				toStack, _ := strconv.Atoi(matches[2])
				fromStack--
				toStack--
				fmt.Printf("Num %d, From %d To %d\n", numMoves, fromStack, toStack)
				fmt.Printf("Before: %s\n", stacks[fromStack])
				fmt.Printf("To Before: %s\n", stacks[toStack])
				fromLoc := len(stacks[fromStack]) - numMoves
				fmt.Printf("fromLoc:numMoves %d:%d\n", fromLoc, numMoves)
				var moveMe []string
				for x := fromLoc; x < numMoves+fromLoc; x++ {
					moveMe = append(moveMe, stacks[fromStack][x])
				}
				fmt.Printf("MoveMe: %s\n", moveMe)
				for x := 1; x <= numMoves; x++ {
					stacks[fromStack] = removeP2(stacks[fromStack], 1)
				}
				stacks[toStack] = append(stacks[toStack], moveMe...)
				fmt.Printf("After: %s\n", stacks[fromStack])
				fmt.Printf("To After: %s\n", stacks[toStack])

			}
		}
		line++
		// scanner.Text()
	}

	for i := 0; i < 9; i++ {
		fmt.Printf("Stack %d: %s\n", i+1, stacks[i])
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

	stackNum := getStackCount(fileName)
	stacks := make(map[int][]string, 0)

	re := regexp.MustCompile("[0-9]+")
	// totalScore := 0
	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "[") {
			tmp := []byte(scanner.Text())
			for i := 0; i < stackNum; i++ {
				item := strings.TrimSpace(string(tmp[(4 * i):((4 * i) + 3)]))
				fmt.Println(item)
				if strings.Contains(item, "[") {
					stacks[i] = injectP1(stacks[i], item)
				}
			}
		} else {
			if strings.HasPrefix(scanner.Text(), "move") {
				matches := re.FindAllString(scanner.Text(), -1)
				numMoves, _ := strconv.Atoi(matches[0])
				fromStack, _ := strconv.Atoi(matches[1])
				toStack, _ := strconv.Atoi(matches[2])
				fromStack--
				toStack--
				fmt.Printf("Num %d, From %d To %d\n", numMoves, fromStack, toStack)
				fmt.Printf("Before: %s\n", stacks[fromStack])
				fmt.Printf("To Before: %s\n", stacks[toStack])
				for x := 1; x <= numMoves; x++ {
					moveMe := stacks[fromStack][len(stacks[fromStack])-1]
					stacks[fromStack] = removeP1(stacks[fromStack], 1)
					stacks[toStack] = append(stacks[toStack], moveMe)
				}
				fmt.Printf("After: %s\n", stacks[fromStack])
				fmt.Printf("To After: %s\n", stacks[toStack])

			}
		}
		line++
		// scanner.Text()
	}

	for i := 0; i < 9; i++ {
		fmt.Printf("Stack %d: %s\n", i+1, stacks[i])
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
