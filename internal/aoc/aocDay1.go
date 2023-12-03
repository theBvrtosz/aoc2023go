package aoc

import (
	"regexp"
	"strconv"
	"strings"
)

func getOnlyNumericsFromRow(row string) string {
	onlyNumericsRegexp := regexp.MustCompile("[^0-9]")
	onlyNumericsRow := onlyNumericsRegexp.ReplaceAllString(row, "")
	return onlyNumericsRow
}

func getFirstAndLastCharacterFromRow(row string) (string, string) {
	rowLen := len(row)
	firstDigit := string(row[0:1])
	lastDigit := string(row[rowLen-1 : rowLen])
	return firstDigit, lastDigit
}

func Day1Part1(isTestRun bool) int {
	dayInput := readDayInput("1", isTestRun)
	calibrationSum := 0
	for _, row := range dayInput {
		onlyNumericsRow := getOnlyNumericsFromRow(row)
		firstDigit, lastDigit := getFirstAndLastCharacterFromRow(onlyNumericsRow)
		coordinate, _ := strconv.Atoi(firstDigit + lastDigit)
		calibrationSum = calibrationSum + coordinate
	}
	return calibrationSum
}

func Day1Part2(isTestRun bool) int {
	spelled_digits := map[string]string{
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}
	dayInput := readDayInput("1", isTestRun)
	coordinateSum := 0
	for _, row := range dayInput {
		for key, val := range spelled_digits {
			if strings.Contains(row, key) {
				row = strings.ReplaceAll(row, key, val)
			}
		}
		onlyNumericsRow := getOnlyNumericsFromRow(row)
		firstDigit, lastDigit := getFirstAndLastCharacterFromRow(onlyNumericsRow)
		coordinate, _ := strconv.Atoi(firstDigit + lastDigit)
		coordinateSum = coordinateSum + coordinate
	}
	return coordinateSum
}
