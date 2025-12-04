package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFileData(filepath string) []string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func convertStringToDigits(line string) []int {
	digits := []int{}

	digitsInString := strings.Split(line, "")

	for _, v := range digitsInString {
		digit, err := strconv.Atoi(v)

		if err != nil {
			log.Fatal(err)
		}

		digits = append(digits, digit)
	}
	return digits
}

func findMaximumNumber(nums []int, excludePos int) (int, int) {
	max := nums[0]
	position := 0

	if len(nums) == 1 {
		return max, position
	}

	for i := 1; i < len(nums)-excludePos; i++ {
		if (nums[i] > max) && (position < i) {
			max = nums[i]
			position = i
		}
	}

	return max, position
}

func findJoltage(line string, numDigits int) int {

	digits := convertStringToDigits(line)

	number := 0

	for i := 1; i <= numDigits; i++ {
		multiFactor := int(math.Pow(10, float64(numDigits-i)))
		max := 0

		if len(digits) == 0 {
			break
		}

		max, pos := findMaximumNumber(digits, numDigits-i)
		digits = digits[pos+1:]

		number += max * multiFactor
	}

	fmt.Println("Number:", number)

	return number
}

func D3P1() {
	batteryLines := readFileData("inputs/day3.txt")

	totalJoltage := 0

	for _, bl := range batteryLines {
		totalJoltage += findJoltage(bl, 2)
	}

	fmt.Println("Total output joltage: ", totalJoltage) // 17343
}

func D3P2() {
	batteryLines := readFileData("inputs/day3.txt")

	totalJoltage := 0

	for _, bl := range batteryLines {
		totalJoltage += findJoltage(bl, 12)
	}

	fmt.Println("Total output joltage: ", totalJoltage) // 17343
}

func Day3() {
	D3P1()
	D3P2()
}
