package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func d22() {
	file, err := os.Open("./inputs/d2_input_data.txt")

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		fmt.Println("Numbers: ", numbers)
	}

}
