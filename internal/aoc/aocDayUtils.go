package aoc

import (
	"log"
	"os"
	"strings"
)

func readDayInput(dayNumber string, isTest bool) []string {
	inputsDirectory := "./inputs/d" + dayNumber + "/"
	var fileName string
	if isTest {
		fileName = "d" + dayNumber + "-test-input.txt"
	} else {
		fileName = "d" + dayNumber + "-input.txt"
	}
	fileDirectory := inputsDirectory + fileName
	content, err := os.ReadFile(fileDirectory)
	if err != nil {
		log.Fatal(err)
	}
	contentSplitted := strings.Split(string(content), "\n")
	return contentSplitted
}
