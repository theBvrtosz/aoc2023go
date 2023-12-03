package aoc

import (
	"math"
	"strconv"
	"strings"
)

func parseGamesFromInput(input []string) map[int]map[string]float64 {
	parsedGames := make(map[int]map[string]float64)
	for _, game := range input {
		gameSplitted := strings.Split(game, ":")
		gameId, _ := strconv.Atoi(strings.ReplaceAll(gameSplitted[0], "Game ", ""))
		cubeReveals := strings.ReplaceAll(gameSplitted[1], ";", ",")
		parsedGames[gameId] = make(map[string]float64)
		for _, reveal := range strings.Split(cubeReveals, ",") {
			revealSplitted := strings.Split(strings.TrimSpace(reveal), " ")
			cubesCount, _ := strconv.ParseFloat(revealSplitted[0], 64)
			cubesColour := revealSplitted[1]
			currentCubesCount, ok := parsedGames[gameId][cubesColour]
			if ok {
				parsedGames[gameId][cubesColour] = math.Max(cubesCount, currentCubesCount)
			} else {
				parsedGames[gameId][cubesColour] = cubesCount
			}
		}
	} 
	return parsedGames
}

func getCubesCountForCubeColour(cubesPlayed map[string]float64, cubeColour string) float64 {
	var currectCubeCount float64
	if cubeCount, ok := cubesPlayed[cubeColour]; ok {
		currectCubeCount = cubeCount
	} else {
		currectCubeCount = 1.0
	}
	return currectCubeCount
}

func Day2Part1(isTestRun bool) int {
	input := readDayInput("2", isTestRun)
	parsedInput := parseGamesFromInput(input)
	redCubesCount := 12.0
	greenCubesCount := 13.0
	blueCubesCount := 14.0
	gameIdSum := 0
	for gameId, maxCubesPlayed := range parsedInput{
		if maxCubesPlayed["red"] <= redCubesCount &&
		maxCubesPlayed["green"] <= greenCubesCount &&
		maxCubesPlayed["blue"] <= blueCubesCount {
			gameIdSum = gameIdSum + gameId
		}
	}
	return gameIdSum
}

func Day2Part2(isTestRun bool) int {
	input := readDayInput("2", isTestRun)
	parsedInput := parseGamesFromInput(input)
	sumOfGamesPower := 0
	for _, maxCubesPlayed := range parsedInput {
		greenCubeCount := getCubesCountForCubeColour(maxCubesPlayed, "green")
		blueCubeCount := getCubesCountForCubeColour(maxCubesPlayed, "blue")
		redCubeCount := getCubesCountForCubeColour(maxCubesPlayed, "red")
		sumOfGamesPower = sumOfGamesPower + int(greenCubeCount*blueCubeCount*redCubeCount)
	}
	return sumOfGamesPower
}