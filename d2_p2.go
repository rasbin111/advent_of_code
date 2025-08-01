package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkIsIncreadingOrIsDecreasing(numbers []int) bool {
	isIncreasing := true
	isDecreasing := true

	smallest := numbers[0]
	largest := numbers[0]

	if numbers[0] > numbers[1] {
		isDecreasing = false
		isIncreasing = true
	} else if numbers[0] > numbers[1] {
		isDecreasing = true
		isIncreasing = false
	}

	for _, v := range numbers[1:] {
		if isIncreasing {

		}
	}
	return isIncreasing || isDecreasing
}

func checkDifferenceBetweenAdj(numbers []int) bool {
	return true
}

func checkSafety(numbers []int) bool {
	check_rule1 := checkIsIncreadingOrIsDecreasing(numbers)
	check_rule2 := true
	if check_rule1 {
		check_rule2 = checkDifferenceBetweenAdj(numbers)
	}

	return check_rule1 && check_rule2
}

func d22() {
	file, err := os.Open("./inputs/d2_input_data.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	safe_levels_count := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		num_in_string_sl := strings.Fields(line)
		numbers := []int{}
		for _, v := range num_in_string_sl {
			i, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Conversion to int err")
			}
			numbers = append(numbers, i)

		}
		isSafe := checkSafety(numbers)
		if isSafe {
			safe_levels_count += 1
		}
	}

}
