package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("./d1_input_data.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	var arr1, arr2 []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) >= 2 {
			num1, _ := strconv.Atoi(parts[0])
			num2, _ := strconv.Atoi(parts[1])

			arr1 = append(arr1, num1)
			arr2 = append(arr2, num2)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	sum := 0
	diff := 0
	for i := 0; i < len(arr1); i++ {
		if arr1[i] >= arr2[i] {
			diff = arr1[i] - arr2[i]
		} else {
			diff = arr2[i] - arr1[i]
		}
		sum += diff
	}
	fmt.Println(sum)
}
