package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	Row int
	Col int
}

func displayGridP2(positions map[int]Coord) {

	displayGrid := false
	if displayGrid {
		colLeft := 0
		colRight := 10
		colLength := 10
		rowBottom := 0
		rowTop := 10
		rowLength := 10
		// Seriously, what was I thinking
		for x := 0; x < 10; x++ {
			if positions[x].Col < colLeft {
				colLeft = positions[x].Col
			}
			if positions[x].Col > colRight {
				colRight = positions[x].Col
			}
			if positions[x].Row < rowBottom {
				rowBottom = positions[x].Row
			}
			if positions[x].Row > rowTop {
				rowTop = positions[x].Row
			}
		}
		// fmt.Printf("colLeft: %d, colRight: %d\n", colLeft, colRight)
		// fmt.Printf("rowBottom: %d, rowTop: %d\n", rowBottom, rowTop)
		colLength = int(math.Abs(float64(colRight - colLeft)))
		rowLength = int(math.Abs(float64(rowTop - rowBottom)))

		// fmt.Printf("colLength: %d, rowLength: %d\n", colLength, rowLength)

		fmt.Printf("\n")
		fmt.Printf("\n")
		marker := "."
		for c := colLength; c >= colLeft; c-- {
			for r := rowBottom; r <= rowLength; r++ {
				marker = "."
				if c == 0 && r == 0 {
					marker = "s"
				}
				for x := 9; x >= 1; x-- {
					if positions[x].Col == c && positions[x].Row == r {
						marker = fmt.Sprintf("%d", x)
					}
					if positions[0].Col == c && positions[0].Row == r {
						marker = "H"
					}
				}
				fmt.Printf("%s", marker)
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func displayGridP1(headAt *Coord, tailAt *Coord) {

	// Seriously, what was I thinking
	longestCol := 6
	highestRow := 6

	if headAt.Col > tailAt.Col {
		if headAt.Col > highestRow {
			highestRow = headAt.Col
		} else {
			if tailAt.Col > highestRow {
				highestRow = tailAt.Col
			}
		}
	}
	if headAt.Row > tailAt.Row {
		if headAt.Row > longestCol {
			longestCol = headAt.Row
		} else {
			if tailAt.Row > longestCol {
				longestCol = tailAt.Row
			}
		}
	}

	fmt.Printf("\n")
	fmt.Printf("\n")
	marker := "."
	for c := highestRow; c >= 0; c-- {
		for r := 0; r <= longestCol; r++ {
			marker = "."
			if c == 0 && r == 0 {
				marker = "s"
			}
			if tailAt.Col == c && tailAt.Row == r {
				marker = "T"
			}
			if headAt.Col == c && headAt.Row == r {
				marker = "H"
			}
			fmt.Printf("%s", marker)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func distanceAway(headAt *Coord, tailAt *Coord) Coord {
	newCoord := Coord{Row: 0, Col: 0}

	newCoord.Col = headAt.Col - tailAt.Col
	newCoord.Row = headAt.Row - tailAt.Row
	// fmt.Printf("DIST row/col: %d,%d\n", newCoord.Row, newCoord.Col)
	return newCoord
}

func secondaryMovement(distance int) int {

	if distance >= 1 {
		return 1
	}
	if distance <= -1 {
		return -1
	}
	return 0
}

func tailMovement(headAt *Coord, tailAt *Coord, visits map[Coord]bool, hMark string, tMark string) Coord {
	newCoord := Coord{Row: 0, Col: 0}

	tailStart := *tailAt
	// fmt.Printf("=== INCOMING: headAt %d,%d : tailAt %d,%d\n", headAt.Row, headAt.Col, tailAt.Row, tailAt.Col)
	separation := distanceAway(headAt, tailAt)
	needMoves := true
	if separation.Col > 1 {
		tailAt.Col++
		newCoord.Col = 1
		if separation.Row <= 1 {
			rowMove := secondaryMovement(separation.Row)
			tailAt.Row = tailAt.Row + rowMove
			newCoord.Row = rowMove
			needMoves = false
		}
		// if headAt.Row != tailAt.Row {
		// 	tailAt.Row = headAt.Row
		// 	newCoord.Row = headAt.Row - tailAt.Row
		// }
	}
	if separation.Col < -1 && needMoves {
		tailAt.Col--
		newCoord.Col = -1
		if separation.Row >= -1 {
			rowMove := secondaryMovement(separation.Row)
			tailAt.Row = tailAt.Row + rowMove
			newCoord.Row = rowMove
			needMoves = false
		}
		// if headAt.Row != tailAt.Row {
		// 	tailAt.Row = headAt.Row
		// 	newCoord.Row = headAt.Row - tailAt.Row
		// }
	}
	if separation.Row > 1 && needMoves {
		tailAt.Row++
		newCoord.Row = 1
		if separation.Col <= 1 {
			colMove := secondaryMovement(separation.Col)
			tailAt.Col = tailAt.Col + colMove
			newCoord.Col = colMove
			needMoves = false
		}
		// if headAt.Col != tailAt.Col {
		// 	tailAt.Col = headAt.Col
		// 	newCoord.Col = headAt.Col - tailAt.Col
		// }
	}
	if separation.Row < -1 && needMoves {
		tailAt.Row--
		newCoord.Row = -1
		if separation.Col >= -1 {
			colMove := secondaryMovement(separation.Col)
			tailAt.Col = tailAt.Col + colMove
			newCoord.Col = colMove
		}
		// if headAt.Col != tailAt.Col {
		// 	tailAt.Col = headAt.Col
		// 	newCoord.Col = headAt.Col - tailAt.Col
		// }
	}
	fmt.Printf("=== MOVEMENT: (%s) %d,%d moves (%s) %d,%d -> %d,%d\n", hMark, headAt.Row, headAt.Col, tMark, tailStart.Row, tailStart.Col, tailAt.Row, tailAt.Col)
	if tMark == "T" || tMark == "9" {
		fmt.Printf("PING: End Moved %d,%d\n", tailAt.Row, tailAt.Col)
		visits[*tailAt] = true
	}
	return newCoord
}

func moveRightP2(positions map[int]Coord, visits map[Coord]bool, distance int) {

	for r := 0; r < distance; r++ {
		positions[0] = Coord{Row: positions[0].Row + 1, Col: positions[0].Col}
		for k := 0; k < 9; k++ {
			headAt := positions[k]
			tailAt := positions[k+1]
			headMarker := fmt.Sprintf("%d", k)
			if k == 0 {
				headMarker = "H"
			}
			tailMarker := fmt.Sprintf("%d", k+1)
			tailMovement(&headAt, &tailAt, visits, headMarker, tailMarker)
			positions[k+1] = tailAt
			// fmt.Printf("Head %d At rol/col: %d,%d\n", k, headAt.Row, headAt.Col)
			// fmt.Printf("Tail %d At rol/col: %d,%d\n", k+1, tailAt.Row, tailAt.Col)
		}
		displayGridP2(positions)
	}
}
func moveLeftP2(positions map[int]Coord, visits map[Coord]bool, distance int) {

	for r := distance - 1; r >= 0; r-- {
		positions[0] = Coord{Row: positions[0].Row - 1, Col: positions[0].Col}
		for k := 0; k < 9; k++ {
			headAt := positions[k]
			tailAt := positions[k+1]
			headMarker := fmt.Sprintf("%d", k)
			if k == 0 {
				headMarker = "H"
			}
			tailMarker := fmt.Sprintf("%d", k+1)
			tailMovement(&headAt, &tailAt, visits, headMarker, tailMarker)
			positions[k+1] = tailAt
			// fmt.Printf("Head %d At rol/col: %d,%d\n", k, headAt.Row, headAt.Col)
			// fmt.Printf("Tail %d At rol/col: %d,%d\n", k+1, tailAt.Row, tailAt.Col)
		}
		displayGridP2(positions)
	}
}

func moveUpP2(positions map[int]Coord, visits map[Coord]bool, distance int) {

	for c := 0; c < distance; c++ {
		positions[0] = Coord{Row: positions[0].Row, Col: positions[0].Col + 1}
		for k := 0; k < 9; k++ {
			headAt := positions[k]
			tailAt := positions[k+1]
			headMarker := fmt.Sprintf("%d", k)
			if k == 0 {
				headMarker = "H"
			}
			tailMarker := fmt.Sprintf("%d", k+1)
			tailMovement(&headAt, &tailAt, visits, headMarker, tailMarker)
			positions[k+1] = tailAt
			// fmt.Printf("Head %d At rol/col: %d,%d\n", k, headAt.Row, headAt.Col)
			// fmt.Printf("Tail %d At rol/col: %d,%d\n", k+1, tailAt.Row, tailAt.Col)
		}
		displayGridP2(positions)
	}

}

func moveDownP2(positions map[int]Coord, visits map[Coord]bool, distance int) {

	for c := distance - 1; c >= 0; c-- {
		positions[0] = Coord{Row: positions[0].Row, Col: positions[0].Col - 1}
		for k := 0; k < 9; k++ {
			headAt := positions[k]
			tailAt := positions[k+1]
			headMarker := fmt.Sprintf("%d", k)
			if k == 0 {
				headMarker = "H"
			}
			tailMarker := fmt.Sprintf("%d", k+1)
			tailMovement(&headAt, &tailAt, visits, headMarker, tailMarker)
			positions[k+1] = tailAt
			// fmt.Printf("Head %d At rol/col: %d,%d\n", k, headAt.Row, headAt.Col)
			// fmt.Printf("Tail %d At rol/col: %d,%d\n", k+1, tailAt.Row, tailAt.Col)
		}
		displayGridP2(positions)
	}
}

func moveRightP1(headAt *Coord, tailAt *Coord, visits map[Coord]bool, distance int) {

	for r := 0; r < distance; r++ {
		headAt.Row++
		tailMovement(headAt, tailAt, visits, "H", "T")
		displayGridP1(headAt, tailAt)
		// fmt.Printf("HeadAt rol/col: %d,%d\n", headAt.Row, headAt.Col)
		// fmt.Printf("TailAt rol/col: %d,%d\n", tailAt.Row, tailAt.Col)
	}
}
func moveLeftP1(headAt *Coord, tailAt *Coord, visits map[Coord]bool, distance int) {

	for r := distance - 1; r >= 0; r-- {
		headAt.Row--
		tailMovement(headAt, tailAt, visits, "H", "T")
		displayGridP1(headAt, tailAt)
		// fmt.Printf("HeadAt rol/col: %d,%d\n", headAt.Row, headAt.Col)
		// fmt.Printf("TailAt rol/col: %d,%d\n", tailAt.Row, tailAt.Col)
	}
}

func moveUpP1(headAt *Coord, tailAt *Coord, visits map[Coord]bool, distance int) {

	for c := 0; c < distance; c++ {
		headAt.Col++
		tailMovement(headAt, tailAt, visits, "H", "T")
		displayGridP1(headAt, tailAt)
		// fmt.Printf("HeadAt rol/col: %d,%d\n", headAt.Row, headAt.Col)
		// fmt.Printf("TailAt rol/col: %d,%d\n", tailAt.Row, tailAt.Col)
	}

}

func moveDownP1(headAt *Coord, tailAt *Coord, visits map[Coord]bool, distance int) {

	for c := distance - 1; c >= 0; c-- {
		headAt.Col--
		tailMovement(headAt, tailAt, visits, "H", "T")
		displayGridP1(headAt, tailAt)
		// fmt.Printf("HeadAt rol/col: %d,%d\n", headAt.Row, headAt.Col)
		// fmt.Printf("TailAt rol/col: %d,%d\n", tailAt.Row, tailAt.Col)
	}
}

func moveHeadP2(positions map[int]Coord, visits map[Coord]bool, direction string, distance int) {

	switch direction {
	// Movement within Row
	case "R":
		fmt.Printf("--- RIGHT: %d\n", distance)
		moveRightP2(positions, visits, distance)
	case "L":
		fmt.Printf("--- LEFT: %d\n", distance)
		moveLeftP2(positions, visits, distance)

	// Movement within Col
	case "U":
		fmt.Printf("--- UP: %d\n", distance)
		moveUpP2(positions, visits, distance)
	case "D":
		fmt.Printf("--- DOWN: %d\n", distance)
		moveDownP2(positions, visits, distance)
	}

}

func moveHeadP1(headAt *Coord, tailAt *Coord, visits map[Coord]bool, direction string, distance int) {

	switch direction {
	// Movement within Row
	case "R":
		fmt.Printf("--- RIGHT: %d\n", distance)
		moveRightP1(headAt, tailAt, visits, distance)
	case "L":
		fmt.Printf("--- LEFT: %d\n", distance)
		moveLeftP1(headAt, tailAt, visits, distance)

	// Movement within Col
	case "U":
		fmt.Printf("--- UP: %d\n", distance)
		moveUpP1(headAt, tailAt, visits, distance)
	case "D":
		fmt.Printf("--- DOWN: %d\n", distance)
		moveDownP1(headAt, tailAt, visits, distance)
	}

}

func partTwo(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tailVisits := make(map[Coord]bool)
	positions := make(map[int]Coord, 10)
	for x := 0; x < 10; x++ {
		positions[x] = Coord{Row: 0, Col: 0}
	}
	tailVisits[positions[9]] = true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			commands := strings.Fields(scanner.Text())
			dist, _ := strconv.Atoi(commands[1])
			moveHeadP2(positions, tailVisits, commands[0], dist)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ANSWER: %d", len(tailVisits))

}
func partOne(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tailVisits := make(map[Coord]bool)
	headPosition := Coord{Row: 0, Col: 0.}
	tailPosition := Coord{Row: 0, Col: 0.}
	tailVisits[tailPosition] = true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			commands := strings.Fields(scanner.Text())
			dist, _ := strconv.Atoi(commands[1])
			moveHeadP1(&headPosition, &tailPosition, tailVisits, commands[0], dist)
			// displayGridP1(&headPosition, &tailPosition)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ANSWER: %d", len(tailVisits))

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
