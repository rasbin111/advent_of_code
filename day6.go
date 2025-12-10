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

func readFile2SlWithSpaceStr(filepath string, numStrings int) [][]string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lines := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.SplitN(line, " ", numStrings)
		lines = append(lines, lineSlice)
	}
	return lines
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

func D6P2(numStrings int) {

	grid := readFile2SlWithSpaceStr("inputs/day6.txt", numStrings)
	totalCalc := 0
	for i := numStrings - 1; i >= 0; i-- {
		calc := 0
		maxIndRow := len(grid) - 1
		// col := strings.TrimSpace(grid[maxIndRow][i])

		maxLen := len(grid[0][i])
		for _, row := range grid[1:maxIndRow] {
			if len(row[i]) > maxLen {
				maxLen = len(row[i])
			}
		}

		nums := [][]string{}

		for ilx := range maxLen {
			_ = ilx
			dNums := make([]string, maxIndRow)
			nums = append(nums, dNums)

		}

		for j, row := range grid[:maxIndRow] {
			selectedNum := row[i]
			// lenItem := len(selectedNum)
			digitsSlice := strings.Split(selectedNum, "")

			for li := range len(selectedNum) {

				nums[li][j] = digitsSlice[li]
				// nums[li][j] = digitsSlice[lenItem-li-1]
				// nums[maxLen-li-1][j] = digitsSlice[li]
			}
		}
		fmt.Println(nums)

		// switch col {
		// case "*":
		// 	calc = 1

		// 	for i, num := range nums {
		// 		calc *= num
		// 	}

		// case "+":
		// 	for i, num := range nums {
		// 		calc += num
		// 	}
		// }

		totalCalc += calc
	}
	fmt.Println("Total amount: ", totalCalc)
}

func Day6() {
	grid := readFile2SlNS("inputs/day6.txt")
	// D6P1(grid)
	D6P2(len(grid[0]))
}
