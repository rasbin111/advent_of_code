package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func readFile2Arr(filepath string) ([][]string, int) {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lines := [][]string{}
	firstLine := ""
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if count == 0 {
			firstLine = line
		}
		lineSlice := strings.Split(line, "")
		count++
		lines = append(lines, lineSlice)
	}

	return lines, strings.Index(firstLine, "S")
}

func D7P1(grid [][]string, startPos int) {
	numOfSplits := 0
	splitPos := []int{startPos}
	for _, val := range grid[1:] {
		for j, fv := range val {
			if "^" == fv {
				for si, sp := range splitPos {
					if j == sp {
						numOfSplits += 1
						splitPos = append(splitPos[:si], splitPos[si+1:]...)
						splitPos = append(splitPos, j-1)
						splitPos = append(splitPos, j+1)
						splitPos = slices.Compact(splitPos)
					}
				}
			}
		}
	}
	fmt.Println("Number of splits: ", numOfSplits)
}

func Day7() {
	grid, startPos := readFile2Arr("inputs/day7.txt")
	D7P1(grid, startPos)
}
