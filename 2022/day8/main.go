package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type GridMap map[CoordPair]Grid

type Grid struct {
	TreeNumber int
	Height     int
}

type CoordPair struct {
	Row int
	Col int
}

func getTreeScore(gm GridMap, srcCp CoordPair, gs int) int {

	score := 0

	// Score Left
	leftScore := 0
	for col := srcCp.Col - 1; col >= 0; col-- {
		cp := CoordPair{Col: col, Row: srcCp.Row}
		if gm[cp].Height >= gm[srcCp].Height {
			leftScore++
			break
		}
		leftScore++
	}
	fmt.Printf("Row/Col: %d,%d  LeftScore: %d\n", srcCp.Row, srcCp.Col, leftScore)

	// Score Right
	rightScore := 0
	for col := srcCp.Col + 1; col < gs; col++ {
		cp := CoordPair{Col: col, Row: srcCp.Row}
		if gm[cp].Height >= gm[srcCp].Height {
			rightScore++
			break
		}
		rightScore++
	}
	fmt.Printf("Row/Col: %d,%d  RightScore: %d\n", srcCp.Row, srcCp.Col, rightScore)

	// Score Up
	upScore := 0
	for row := srcCp.Row - 1; row >= 0; row-- {
		cp := CoordPair{Col: srcCp.Col, Row: row}
		if gm[cp].Height >= gm[srcCp].Height {
			upScore++
			break
		}
		upScore++
	}
	fmt.Printf("Row/Col: %d,%d  UpScore: %d\n", srcCp.Row, srcCp.Col, upScore)

	//Score Down
	downScore := 0
	for row := srcCp.Row + 1; row < gs; row++ {
		cp := CoordPair{Col: srcCp.Col, Row: row}
		if gm[cp].Height >= gm[srcCp].Height {
			downScore++
			break
		}
		downScore++
	}
	fmt.Printf("Row/Col: %d,%d  DownScore: %d\n", srcCp.Row, srcCp.Col, downScore)

	score = leftScore * rightScore * upScore * downScore

	fmt.Printf("Row/Col: %d,%d  TotalScore: %d\n", srcCp.Row, srcCp.Col, score)
	return score
}
func getBestSpot(gm GridMap, gs int) (CoordPair, int) {

	bestScore := 0
	bestTree := CoordPair{}
	for row := 0; row < gs; row++ {
		for col := 0; col < gs; col++ {
			cp := CoordPair{Col: col, Row: row}
			score := getTreeScore(gm, cp, gs)
			if score > bestScore {
				bestScore = score
				bestTree = cp
			}
		}
	}
	return bestTree, bestScore

}

func getVisibleTrees(gm GridMap, gs int, vt map[int]bool) {

	fmt.Println("--- Row Left to Right ---")
	// Row -> left to right (col 0 -> col <gs>)
	for row := 1; row < gs-1; row++ {
		highestRowTree := 0
		for col := 1; col < gs-1; col++ {
			cp := CoordPair{Col: col, Row: row}
			// fmt.Printf("Row/Col: %d,%d  TreeNumber: %d  TreeSize: %d\n", cp.Row, cp.Col, gm[cp].TreeNumber, gm[cp].Height)
			prevTree := gm[CoordPair{Col: col - 1, Row: row}]
			if prevTree.Height > highestRowTree {
				highestRowTree = prevTree.Height
			}
			if highestRowTree < gm[cp].Height {
				fmt.Printf("Row/Col: %d,%d  TreeNumber: %d  TreeSize: %d\n", cp.Row, cp.Col, gm[cp].TreeNumber, gm[cp].Height)
				vt[gm[cp].TreeNumber] = true
			}
		}
	}
	fmt.Println("--- Row Right to Left ---")
	// Row -> right to left (col <gs> -> col 0)
	for row := 1; row < gs-1; row++ {
		highestRowTree := 0
		for col := gs - 2; col > 0; col-- {
			cp := CoordPair{Col: col, Row: row}
			// fmt.Printf("Row/Col: %d,%d  TreeNumber: %d  TreeSize: %d\n", cp.Row, cp.Col, gm[cp].TreeNumber, gm[cp].Height)
			prevTree := gm[CoordPair{Col: col + 1, Row: row}]
			if prevTree.Height > highestRowTree {
				highestRowTree = prevTree.Height
			}
			if highestRowTree < gm[cp].Height {
				fmt.Printf("Row/Col: %d,%d  TreeNumber: %d  TreeSize: %d\n", cp.Row, cp.Col, gm[cp].TreeNumber, gm[cp].Height)
				vt[gm[cp].TreeNumber] = true
			}
		}
	}

	fmt.Println("--- Col Top to Bottom ---")
	// Col -> top to bottom (row 0 -> row <gs>)
	for col := 1; col < gs-1; col++ {
		highestColTree := 0
		for row := 1; row < gs-1; row++ {
			cp := CoordPair{Col: col, Row: row}
			// fmt.Printf("Row/Col: %d,%d  TreeNumber: %d  TreeSize: %d\n", cp.Row, cp.Col, gm[cp].TreeNumber, gm[cp].Height)
			prevTree := gm[CoordPair{Col: col, Row: row - 1}]
			if prevTree.Height > highestColTree {
				highestColTree = prevTree.Height
			}
			if highestColTree < gm[cp].Height {
				fmt.Printf("Row/Col: %d,%d  TreeNumber: %d  TreeSize: %d\n", cp.Row, cp.Col, gm[cp].TreeNumber, gm[cp].Height)
				vt[gm[cp].TreeNumber] = true
			}
		}
	}

	fmt.Println("--- Col Bottom to Top ---")
	// Col -> top to bottom (row 0 -> row <gs>)
	// Col -> bottom to top (row <gs> -> row 0)
	for col := 1; col < gs-1; col++ {
		highestColTree := 0
		for row := gs - 2; row > 0; row-- {
			cp := CoordPair{Col: col, Row: row}
			// fmt.Printf("Row/Col: %d,%d  TreeNumber: %d  TreeSize: %d\n", cp.Row, cp.Col, gm[cp].TreeNumber, gm[cp].Height)
			prevTree := gm[CoordPair{Col: col, Row: row + 1}]
			if prevTree.Height > highestColTree {
				highestColTree = prevTree.Height
			}
			if highestColTree < gm[cp].Height {
				fmt.Printf("Row/Col: %d,%d  TreeNumber: %d  TreeSize: %d\n", cp.Row, cp.Col, gm[cp].TreeNumber, gm[cp].Height)
				vt[gm[cp].TreeNumber] = true
			}
		}
	}
}

func parse(gm GridMap, line []byte, num int) int {

	gridSize := len(line)
	fmt.Printf("GridSize: %dx%d\n", gridSize, gridSize)
	for i, v := range line {
		treeHeight, _ := strconv.Atoi(string(v))
		fmt.Printf("Row: %d, Tree: %d, Height: %d\n", num, num*gridSize+i, treeHeight)
		// gridSize = 5 (5x5 grid)
		// Row (num)
		// Col (i)
		// Row = 0, Col = 0, TreeNumber = 0
		// Row = 0, Col = 1, TreeNumber = 1
		// Row = 4, Col = 4, TreeNumber = 24 ()
		treeId := Grid{
			TreeNumber: num*gridSize + i,
			Height:     treeHeight,
		}
		cp := CoordPair{Col: i, Row: num}
		gm[cp] = treeId
	}

	return gridSize
}

func partTwo(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hGrid := GridMap{}
	gridSize := 0
	lineNum := 0

	for scanner.Scan() {
		gridSize = parse(hGrid, scanner.Bytes(), lineNum)
		lineNum++
	}

	coord, score := getBestSpot(hGrid, gridSize)

	fmt.Printf("Tree Row/Col: %d,%d  Score: %d\n", coord.Row, coord.Col, score)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
func partOne(fileName string) {

	// rows multi-dim array contains (height,tree number - count across/down)
	// cols multi-dim array contains (height,tree number - count across/down)
	// Check checking rows/columns, visible trees are added to a map using the tree number, that way no counting twice
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hGrid := GridMap{}
	gridSize := 0
	scanner := bufio.NewScanner(file)
	lineNum := 0
	visibleTrees := make(map[int]bool)

	for scanner.Scan() {
		gridSize = parse(hGrid, scanner.Bytes(), lineNum)
		lineNum++
	}

	getVisibleTrees(hGrid, gridSize, visibleTrees)

	exteriorVisible := (gridSize * 2) + ((gridSize - 2) * 2)
	fmt.Printf("Exterior Visible Trees: %d\n", exteriorVisible)
	fmt.Printf("Interior Visible Trees: %d\n", len(visibleTrees))
	fmt.Printf("ANSWER: %d\n", exteriorVisible+len(visibleTrees))
	// for k, v := range hGrid {
	// fmt.Printf("Key: %d,%d  TreeNumber: %d  TreeSize: %d\n", k.X, k.Y, v.TreeNumber, v.Height)
	// }

	// for row := 0; row < gridSize; row++ {
	// 	for col := 0; col < gridSize; col++ {
	// 		cp := CoordPair{Col: col, Row: row}
	// 		fmt.Printf("Row/Col: %d,%d  TreeNumber: %d  TreeSize: %d\n", cp.Row, cp.Col, hGrid[cp].TreeNumber, hGrid[cp].Height)
	// 	}
	// }

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
