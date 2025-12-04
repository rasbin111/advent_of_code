package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readCSVFile(filepath string) []string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return records[0]
}

func patternValidityCheckR1(num int64) bool {

	numOfDigits := len(strconv.FormatInt(num, 10))

	if numOfDigits == 1 {
		return false
	}

	if numOfDigits%2 != 0 {
		return false
	} else {
		halfND := numOfDigits / 2
		factor := int64(math.Pow(10, float64(halfND)))

		return num%factor == num/factor
	}

}

func repeatedNumberSum(record string) int64 {
	var sumRepeatedNumber int64 = 0
	firstNumber, err := strconv.ParseInt(strings.Split(record, "-")[0], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	lastNumber, err := strconv.ParseInt(strings.Split(record, "-")[1], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	for i := firstNumber; i <= lastNumber; i++ {
		rptStatus := patternValidityCheckR1(i)

		if rptStatus {
			sumRepeatedNumber += i
		}
	}

	return sumRepeatedNumber
}

func patternValidityCheckR2(num int64) bool {

	numString := strconv.FormatInt(num, 10)
	numOfDigits := len(numString)

	if numOfDigits == 1 {
		return false
	}

	halfND := numOfDigits / 2
	factor := int64(math.Pow(10, float64(halfND)))

	if numOfDigits%2 == 0 {

		if num%factor == num/factor {
			return true
		}
	}

	for i := 1; i <= halfND; i++ {
		if numOfDigits%i == 0 {

			pattern := numString[:i]
			isRepeating := true

			for j := i; j < numOfDigits; j += i {
				matchingPattern := numString[j : i+j]
				if matchingPattern != pattern {
					isRepeating = false
					break
				}
			}
			if isRepeating {
				return true
			}
		}

	}
	return false
}

func repeatedNumberSum2(record string) int64 {
	var sumRepeatedNumber int64 = 0
	firstNumber, err := strconv.ParseInt(strings.Split(record, "-")[0], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	lastNumber, err := strconv.ParseInt(strings.Split(record, "-")[1], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	for i := firstNumber; i <= lastNumber; i++ {
		rptStatus := patternValidityCheckR2(i)

		if rptStatus {
			sumRepeatedNumber += i
		}
	}

	return sumRepeatedNumber
}

func Day2Part1() {
	records := readCSVFile("inputs/day2.txt")
	var sumOfAllRanges int64 = 0

	for _, record := range records {
		sumOfAllRanges += repeatedNumberSum(record)
	}

	fmt.Println("Total invalid ids: ", sumOfAllRanges)

}

func Day2Part2() {
	records := readCSVFile("inputs/day2.txt")
	var sumOfAllRanges int64 = 0

	for _, record := range records {
		sumOfAllRanges += repeatedNumberSum2(record)
	}

	fmt.Println("Total invalid ids: ", sumOfAllRanges) // answer: 20942028255

}

func Day2() {
	Day2Part1()
	Day2Part2()
}
