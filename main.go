package main

import (
	"fmt"
	"sort"

	"aoc.thebvrtosz.net/internal/aoc"
)

var dayMap = map[int][2]func(bool) int {
	1: {aoc.Day1Part1, aoc.Day1Part2},
	2: {aoc.Day2Part1, aoc.Day2Part2},
}

func printDaySolution(dayNumber int, isTestRun bool) {
	fmt.Printf("day %v:\n", dayNumber)
	fmt.Printf("\t part1 solution: %v\n", dayMap[dayNumber][0](isTestRun))
	fmt.Printf("\t part2 solution: %v\n", dayMap[dayNumber][1](isTestRun))
	fmt.Println("--------------------")
}

func main() {
	days := make([]int, 0)
	for day := range dayMap {
		days = append(days, day)
	}
	sort.Ints(days)
	for _, day := range days {
		printDaySolution(day, false)
	}
}
