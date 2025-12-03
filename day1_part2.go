package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fileDataToArr2(filepath string) []string {
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

func Day1Part2() {
	datas := fileDataToArr2("inputs/day1.txt")

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
