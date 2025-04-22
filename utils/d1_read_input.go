package utils 

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func File_data_to_arrays(arr1 *[]int, arr2 *[]int) {

	file, err := os.Open("./inputs/d1_input_data.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) >= 2 {
			num1, _ := strconv.Atoi(parts[0])
			num2, _ := strconv.Atoi(parts[1])

			*arr1 = append(*arr1, num1)
			*arr2 = append(*arr2, num2)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}

}
