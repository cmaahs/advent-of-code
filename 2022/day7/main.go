package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	currentPath := "/"
	dirTotal := 0
	dirSizes := make(map[string]int)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "$") {
			if strings.HasPrefix(scanner.Text(), "$ cd") {
				if scanner.Text() == "$ cd /" {
					currentPath = "/"
				} else {
					pathTarget := strings.TrimPrefix(scanner.Text(), "$ cd ")
					if _, ok := dirSizes[currentPath]; !ok {
						fmt.Printf("Set Size: %s, %d\n", currentPath, dirTotal)
						dirSizes[currentPath] = dirTotal
						dirTotal = 0
					} else {
						if dirTotal != 0 {
							fmt.Printf("Set Size: %s, %d\n", currentPath, dirTotal)
							dirSizes[currentPath] = dirTotal
							dirTotal = 0
						}
					}
					if pathTarget == ".." {
						pathParts := strings.Split(currentPath, "/")
						currentPath = strings.Join(pathParts[:len(pathParts)-1], "/")
					} else {
						currentPath = fmt.Sprintf("%s/%s", currentPath, strings.Split(scanner.Text(), " ")[2])
					}
				}
				fmt.Printf("INFO: New Path: %s\n", currentPath)
			}
		} else {
			if !strings.HasPrefix(scanner.Text(), "dir") {
				fmt.Printf("file: %s\n", scanner.Text())
				fileSize, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
				dirTotal = dirTotal + fileSize
			}
		}

	}
	dirSizes[currentPath] = dirTotal
	dirTotal = 0
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	cumulativeTotals := make(map[string]int, len(dirSizes))
	keys := make([]string, 0, len(dirSizes))
	for k, v := range dirSizes {
		keys = append(keys, k)
		cumulativeTotals[k] = v
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("Sorted: %s: %d\n", k, dirSizes[k])
		parents := strings.Split(k, "/")
		key := ""
		for i := 1; i < len(parents)-1; i++ {
			if parents[i] == " " {
				key = ""
				fmt.Printf("Add to root %s, %d\n", k, dirSizes[k])
				cumulativeTotals["/"] = cumulativeTotals["/"] + dirSizes[k]
			} else {
				key = strings.Replace(fmt.Sprintf("%s/%s", key, parents[i]), "///", "//", 1)
				fmt.Printf("Add to Parent %s, %d\n", key, dirSizes[k])
				cumulativeTotals[key] = cumulativeTotals[key] + dirSizes[k]
			}
		}
	}

	usedSpace := cumulativeTotals["/"]
	freeSpace := 70000000 - usedSpace
	needFree := 30000000 - freeSpace

	fmt.Printf("used Space: %d\n", usedSpace)
	fmt.Printf("free Space: %d\n", freeSpace)
	fmt.Printf("Need Free: %d\n", needFree)

	cKeys := make([]string, 0, len(cumulativeTotals))
	for k := range cumulativeTotals {
		cKeys = append(cKeys, k)
	}
	sort.Strings(cKeys)
	smallest := usedSpace
	for _, k := range cKeys {
		if cumulativeTotals[k] >= needFree {
			fmt.Printf("Cumulative: %s: %#v\n", k, cumulativeTotals[k])
			if cumulativeTotals[k] < smallest {
				smallest = cumulativeTotals[k]
			}
		}
	}

	fmt.Printf("Anser: %d", smallest)
}
func partOne(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentPath := "/"
	dirTotal := 0
	dirSizes := make(map[string]int)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "$") {
			if strings.HasPrefix(scanner.Text(), "$ cd") {
				if scanner.Text() == "$ cd /" {
					currentPath = "/"
				} else {
					pathTarget := strings.TrimPrefix(scanner.Text(), "$ cd ")
					if _, ok := dirSizes[currentPath]; !ok {
						fmt.Printf("Set Size: %s, %d\n", currentPath, dirTotal)
						dirSizes[currentPath] = dirTotal
						dirTotal = 0
					} else {
						if dirTotal != 0 {
							fmt.Printf("Set Size: %s, %d\n", currentPath, dirTotal)
							dirSizes[currentPath] = dirTotal
							dirTotal = 0
						}
					}
					if pathTarget == ".." {
						pathParts := strings.Split(currentPath, "/")
						currentPath = strings.Join(pathParts[:len(pathParts)-1], "/")
					} else {
						currentPath = fmt.Sprintf("%s/%s", currentPath, strings.Split(scanner.Text(), " ")[2])
					}
				}
				fmt.Printf("INFO: New Path: %s\n", currentPath)
			}
		} else {
			if !strings.HasPrefix(scanner.Text(), "dir") {
				fmt.Printf("file: %s\n", scanner.Text())
				fileSize, _ := strconv.Atoi(strings.Split(scanner.Text(), " ")[0])
				dirTotal = dirTotal + fileSize
			}
		}

	}
	dirSizes[currentPath] = dirTotal
	dirTotal = 0
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	cumulativeTotals := make(map[string]int, len(dirSizes))
	keys := make([]string, 0, len(dirSizes))
	for k, v := range dirSizes {
		keys = append(keys, k)
		cumulativeTotals[k] = v
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("Sorted: %s: %d\n", k, dirSizes[k])
		parents := strings.Split(k, "/")
		key := ""
		for i := 1; i < len(parents)-1; i++ {
			if parents[i] == " " {
				key = ""
				fmt.Printf("Add to root %s, %d\n", k, dirSizes[k])
				cumulativeTotals["/"] = cumulativeTotals["/"] + dirSizes[k]
			} else {
				key = strings.Replace(fmt.Sprintf("%s/%s", key, parents[i]), "///", "//", 1)
				fmt.Printf("Add to Parent %s, %d\n", key, dirSizes[k])
				cumulativeTotals[key] = cumulativeTotals[key] + dirSizes[k]
			}
		}
	}

	cKeys := make([]string, 0, len(cumulativeTotals))
	for k := range cumulativeTotals {
		cKeys = append(cKeys, k)
	}
	sort.Strings(cKeys)
	sumSizes := 0
	for _, k := range cKeys {
		if cumulativeTotals[k] < 100000 {
			fmt.Printf("Cumulative: %s: %#v\n", k, cumulativeTotals[k])
			sumSizes = sumSizes + cumulativeTotals[k]
		}
	}

	fmt.Printf("Answer: %d", sumSizes)
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
