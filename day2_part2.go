package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func patternValidityCheckR2(num int64) bool {

	numOfDigits := len(strconv.FormatInt(num, 10))

	if numOfDigits == 1 {
		return false
	}

	halfND := numOfDigits / 2
	factor := int64(math.Pow(10, float64(halfND)))

	if num%factor == num/factor {
		return true
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

func Day2Part2() {
	records := readCSVFile("inputs/day2.txt")
	var sumOfAllRanges int64 = 0

	for _, record := range records {
		sumOfAllRanges += repeatedNumberSum2(record)
	}

	fmt.Println("Total invalid ids: ", sumOfAllRanges)

}
