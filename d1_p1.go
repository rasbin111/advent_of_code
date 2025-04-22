package main

import (
	"fmt"
	"sort"
    "advent_of_code/utils"
)

func d11() {
	var arr1, arr2 []int

	utils.File_data_to_arrays(&arr1, &arr2)

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
