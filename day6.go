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

func readFile2SlWithSpaceStr(filepath string, maxSizes []int) [][]string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lines := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		i := 0
		for j := 0; j <= len(line)-maxSizes[len(maxSizes)-1]; j++ {
			maxDigits := maxSizes[i]
			num := line[j : j+maxDigits]
			row = append(row, num)
			j = j + maxDigits
			i++
		}
		lines = append(lines, row)
	}
	return lines
}

func maxLengthOfNum(lines [][]string) []int {
	maxSize := []int{}
	for j := range lines[0] {
		maxLen := 0
		for i := range lines {
			str := lines[i][j]
			lenStr := len(str)
			if maxLen < lenStr {
				maxLen = len(lines[i][j])
			}
		}
		maxSize = append(maxSize, maxLen)
	}
	return maxSize
}

func strToIntOrDefZero(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))

	if err != nil {
		return 0
	}
	return i
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

func D6P2Calculation(lines [][]string) {
	totalCalc := 0
	for _, value := range lines {
		symbol := strings.TrimSpace(value[len(value)-1])
		calc := 0
		switch symbol {
		case "*":
			calc = 1
			for _, fv := range value[:len(value)-1] {
				num := strToIntOrDefZero(fv)

				calc *= num
			}
		case "+":
			for _, fv := range value[:len(value)-1] {
				num := strToIntOrDefZero(fv)

				calc += num
			}
		}
		totalCalc += calc
	}
	fmt.Println("Total amount: ", totalCalc)

}

func D6P2(maxSizes []int) {

	grid := readFile2SlWithSpaceStr("inputs/day6.txt", maxSizes)
	nums := [][]string{}

	for j := len(grid[0]) - 1; j >= 0; j-- {
		lastRowIdx := len(grid) - 1
		symbol := grid[lastRowIdx][j]

		maxDigits := maxSizes[j]
		tempNums := make([]string, maxDigits+1)
		tempNums[maxDigits] = symbol
		_ = tempNums
		_ = symbol

		for i := 0; i < len(grid)-1; i++ {
			selectedNum := grid[i][j]
			// mf := 1
			for d := 0; d < maxDigits; d++ {
				num := tempNums[d]
				digit := string(selectedNum[d])
				num = num + digit
				tempNums[d] = num
				// num = n
			}
		}
		nums = append(nums, tempNums)
	}
	D6P2Calculation(nums)
}

func Day6() {
	grid := readFile2SlNS("inputs/day6.txt")
	maxSizes := maxLengthOfNum(grid)
	D6P2(maxSizes)
}
