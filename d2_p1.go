package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func d21() {
	file, err := os.Open("./inputs/d2_input_data.txt")

	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	scanner := bufio.NewScanner(file)
	safeLineCount := 0
	for scanner.Scan() {
		line := scanner.Text() // get line

		numbers := strings.Fields(line) // numbers := strings.Split(line, " ") // this also works
		isIncreasing := false
		isDecreasing := false
		isSafe := true
		diff := 0 // 4 3  2 1
		prevValue, _ := strconv.Atoi(numbers[0])
		for _, value := range numbers[1:] {
			num, _ := strconv.Atoi(value)
			if num == prevValue {
				isSafe = false
				break
			} else if num > prevValue {
				diff = num - prevValue
				if diff >= 1 && diff <= 3 {
					isIncreasing = true
					if isDecreasing {
						isSafe = false
						break
					}
				} else {
					isSafe = false
					break
				}

			} else {
				diff = prevValue - num
				if diff >= 1 && diff <= 3 {
					isDecreasing = true
					if isIncreasing {
						isSafe = false
						break
					}
				} else {
					isSafe = false
					break
				}
			}

			prevValue = num
		}
		if isSafe {
			safeLineCount += 1
		}

	}
	fmt.Println("Final safe line count: ", safeLineCount)
}
