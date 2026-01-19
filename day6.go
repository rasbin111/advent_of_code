package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile2SlNS(filepath string) [][]string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lines := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Fields(line)
		lines = append(lines, lineSlice)
	}
	return lines
}

func readFile2SlWithSpaceStr(filepath string, maxDigits int) [][]string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lines := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for i := 0; i <= len(line)-maxDigits; i = i + maxDigits + 1 {
			num := line[i : i+maxDigits]
			row = append(row, num)
		}
		lines = append(lines, row)
	}
	return lines
}

func maxLengthOfNum(lines [][]string) int {
	maxLen := len(lines[0][0])

	for i := range lines {
		for j := range lines[0] {
			if maxLen < len(lines[i][j]) {
				maxLen = len(lines[i][j])
			}
		}
	}
	return maxLen
}

func D6P1(grid [][]string) {
	totalCalc := 0
	for i, col := range grid[len(grid)-1] {
		calc := 0

		switch col {
		case "*":
			calc = 1
			for _, row := range grid[:len(grid)-1] {
				num, err := strconv.Atoi(row[i])

				if err != nil {
					log.Fatal(err)
				}
				calc *= num
			}
		case "+":
			for _, row := range grid[:len(grid)-1] {
				num, err := strconv.Atoi(row[i])

				if err != nil {
					log.Fatal(err)
				}
				calc += num
			}
		}

		totalCalc += calc
	}
	fmt.Println("Total amount: ", totalCalc)
}

func D6P2(maxDigits int) {

	grid := readFile2SlWithSpaceStr("inputs/day6.txt", maxDigits)
	fmt.Println(grid)
	nums := make([]int, len(grid[0]))
	_ = nums
	for j := len(grid[0]) - 1; j >= 0; j-- {
		lastRowIdx := len(grid) - 1
		symbol := grid[lastRowIdx][j]
		// fmt.Println("Symbol: ", symbol)
		// countTotalNums := len(grid[0][j])
		tempNums := make([]int64, maxDigits)
		// _ = tempNums
		_ = symbol
		num := 0
		for i := 0; i < len(grid)-1; i++ {
			num := grid[i][j]
			for d := 0; d < maxDigits; d++ {
				n, err := strconv.ParseInt(string(num[d]), 10, 64)

				if err != nil {
					log.Fatal(err)
				}
				num := n
			}
		}
	}
	// fmt.Println("Total amount: ", totalCalc)
}

func Day6() {
	grid := readFile2SlNS("inputs/day6.txt")
	maxDigits := maxLengthOfNum(grid)
	D6P2(maxDigits)
}
