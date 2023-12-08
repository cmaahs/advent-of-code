package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type ModuleInfo struct {
	Day string
}

func main() {

	if len(os.Args) < 3 {
		logrus.Fatal("failed to provide enough args")
	}
	year := os.Args[1]
	day := os.Args[2]

	createDayCode(year, day)
}

func createDayCode(year string, day string) {

	// ├── 2023
	// │   └── day1
	// │       ├── day1_suite_test.go
	// │       ├── day1_test.go
	// │       ├── go.mod
	// │       ├── go.sum
	// │       ├── input.txt
	// │       ├── main.go
	// │       ├── sample1.txt
	// │       ├── sample2.txt
	// │       └── solution
	// │           └── solution.go

	// Setup the directories
	dayDir := fmt.Sprintf("%s/day%s/solution", year, day)
	logrus.Infof("Directory: %s", dayDir)
	if _, err := os.Stat(fmt.Sprintf("./%s", dayDir)); err != nil {
		if os.IsNotExist(err) {
			mkerr := os.MkdirAll(fmt.Sprintf("./%s", dayDir), os.ModePerm)
			if mkerr != nil {
				logrus.WithError(mkerr).Fatalf("Error creating %s directory", dayDir)
			}
		}
	}

	homeDir, herr := os.UserHomeDir()
	destinationBasePath := fmt.Sprintf("%s/src/cmaahsProjects/advent-of-code/%s/day%s", homeDir, year, day)
	if herr != nil {
		logrus.WithError(herr).Fatal("failed to fetch homedir")
	}
	templateDir := fmt.Sprintf("%s/src/cmaahsProjects/advent-of-code/templater/templates", homeDir)
	createSource(destinationBasePath, templateDir, year, day)
	// templateDir = fmt.Sprintf("%s/src/cmaahsProjects/advent-of-code/templater/templates/solution", homeDir)
	// createSource(destinationBasePath, templateDir, year, day)

}

func createSource(destinationBasePath string, templateDir string, year string, day string) {

	err := filepath.Walk(templateDir,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				shortPath := strings.TrimPrefix(path, templateDir)
				shortPath = strings.TrimPrefix(shortPath, "/")
				if err != nil {
					return err
				}
				if len(shortPath) > 0 {
					newFile := strings.TrimSuffix(shortPath, ".tpl")
					newFile = fmt.Sprintf("%s/%s", destinationBasePath, newFile)
					logrus.Infof("File: %s", newFile)
					temp := template.Must(template.ParseFiles(path))
					tplVariables := ModuleInfo{
						Day: day,
					}

					f, ferr := os.Create(newFile)
					if ferr != nil {
						logrus.WithError(ferr).Fatal("create file: ", ferr)
					}

					err := temp.Execute(f, tplVariables)
					if err != nil {
						log.Fatalln(err)
					}
					f.Close()
				}
			}
			return nil
		})
	if err != nil {
		logrus.Error(err)
	}
}
