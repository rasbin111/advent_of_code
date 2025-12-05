package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFileDataTo2DArr(filepath string) [][]string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lines := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, "")
		lines = append(lines, lineSlice)
	}
	return lines
}

func rollsCounter(grid [][]string, replace bool) int {
	lenOuterArr := len(grid)
	lenInnerArr := len(grid[0])

	gridPosition := [][]int{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}

	rollsCount := 0
	for i := range lenOuterArr {
		for j := range lenInnerArr {

			if grid[i][j] == "@" {
				adjacentCount := 0

				for gx := range len(gridPosition) {
					xPosition := i + gridPosition[gx][0]
					yPosition := j + gridPosition[gx][1]

					if (xPosition < 0) || (xPosition >= lenOuterArr) || (yPosition < 0) || (yPosition >= lenInnerArr) {
						continue
					}

					if grid[xPosition][yPosition] == "@" {
						adjacentCount += 1
					}
				}

				if adjacentCount < 4 {
					rollsCount += 1
					if replace {
						grid[i][j] = "x"
					}
				}
			}

		}
	}
	return rollsCount
}

func D4P1() {
	grid := readFileDataTo2DArr("inputs/day4.txt")

	rollsCount := rollsCounter(grid, false)

	fmt.Println("Rolls Count: ", rollsCount)
}

func D4P2() {
	grid := readFileDataTo2DArr("inputs/day4.txt")
	totalRollsCount := 0
	for {
		rollsCount := rollsCounter(grid, true)
		totalRollsCount += rollsCount

		if rollsCount == 0 {
			break
		}
	}

	fmt.Println("Rolls Count: ", totalRollsCount)
}

func Day4() {
	D4P1()
	D4P2()
}
