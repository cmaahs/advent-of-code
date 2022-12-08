package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// A = rock, B = paper, C = scissors
// X = rock, Y = paper, Z = scissors
// +1        +2         +3
// Win + 6, Draw +3, Loss +0
// X = loss, Y = draw, Z = win

type Strat struct {
	Opponent   string
	Throw      string
	RoundScore int
}

func partTwo(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scoreTable := make(map[string]int, 0)
	scoreTable["A X"] = 3 // draw 4 / loss 3
	scoreTable["A Y"] = 4 // win  8 / draw 4
	scoreTable["A Z"] = 8 // loss 3 / win  8
	scoreTable["B X"] = 1 // loss 1 / loss 1
	scoreTable["B Y"] = 5 // draw 5 / draw 5
	scoreTable["B Z"] = 9 // win  9 / win  9
	scoreTable["C X"] = 2 // win  7 / loss 2
	scoreTable["C Y"] = 6 // loss 2 / draw 6
	scoreTable["C Z"] = 7 // draw 6 / win  7

	totalScore := 0
	for scanner.Scan() {
		totalScore = totalScore + scoreTable[scanner.Text()]
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

	scoreTable := make(map[string]int, 0)
	scoreTable["A X"] = 4 // draw 4 / loss 3
	scoreTable["A Y"] = 8 // win  8 / draw 4
	scoreTable["A Z"] = 3 // loss 3 / win  8
	scoreTable["B X"] = 1 // loss 1 / loss 1
	scoreTable["B Y"] = 5 // draw 5 / draw 5
	scoreTable["B Z"] = 9 // win  9 / win  9
	scoreTable["C X"] = 7 // win  7 / loss 2
	scoreTable["C Y"] = 2 // loss 2 / draw 6
	scoreTable["C Z"] = 6 // draw 6 / win  7

	totalScore := 0
	for scanner.Scan() {
		totalScore = totalScore + scoreTable[scanner.Text()]
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
