package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func d2() {
	file, err := os.Open("./d2_input_data.txt")

	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	scanner := bufio.NewScanner(file)
	safeLineCount := 0
	for scanner.Scan() {
		line := scanner.Text() // get line

		numbers := strings.Fields(line) // numbers := strings.Split(line, " ") // this also works
		isIncreasing := true
		isDecreasing := true
		isDecreasingOrIncreasing := true
		isSafe := true
		max, _ := strconv.Atoi(numbers[0])
		min, _ := strconv.Atoi(numbers[0])
		prevValue, _ := strconv.Atoi(numbers[0])
		for _, value := range numbers {
			num, _ := strconv.Atoi(value)
			if max <= num {
				isIncreasing = false
			} else {
				max = num
			}
			if min >= num {
				isDecreasing = false
			} else {
				min = num
			}
			if !isIncreasing && !isDecreasing {
				isDecreasingOrIncreasing = false
				break
			} else {
				if num-prevValue > 3 {
					isSafe = false
					break
				}
			}
			prevValue = num
		}

		if isDecreasingOrIncreasing && isSafe {
			safeLineCount += 1
		}

	}
	fmt.Println("Safe Line Count", safeLineCount)

}
