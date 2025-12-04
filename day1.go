package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fileDataToArr(filepath string) []string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	datas := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		datas = append(datas, line)
	}

	return datas
}

func Day1Part1() {
	datas := fileDataToArr("inputs/day1.txt")

	var position = 50

	password := 0

	for _, value := range datas {
		numberString := value[1:]
		rotationAmount, err := strconv.Atoi(numberString)

		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(value, "R") {
			position += rotationAmount

			if position > 99 {
				position = position % 100
			}

		} else {
			position -= rotationAmount

			if position < 0 {
				position = ((position % 100) + 100) % 100
			}
		}

		if position == 0 {
			password += 1
		}
	}

	fmt.Println("Password: ", password)
}

func Day1Part2() {
	datas := fileDataToArr("inputs/day1.txt")

	var position = 50

	password := 0

	for _, value := range datas {
		numberString := value[1:]
		rotationAmount, err := strconv.Atoi(numberString)

		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(value, "R") {
			position += rotationAmount

			if position > 99 {
				password += position / 100
				position = position % 100
			}

		} else {
			start := (100 - position) % 100
			password += (start+(rotationAmount%100))/100 + rotationAmount/100
			position -= rotationAmount
			if position < 0 {
				position = ((position % 100) + 100) % 100
			}
		}

	}

	fmt.Println("Password: ", password) // should be 6634
}

func Day1() {
	Day1Part1()
	Day1Part2()
}
