package main
import "fmt"

func d12() {
	var arr1, arr2 []int
	file_data_to_arrays(&arr1, &arr2)
    valueCountMap := make(map[int]int)
    var similarity_score = 0
    for _, v1 := range arr1{
        count := 0
        for _, v2 := range arr2{
            if v1 == v2{
                count += 1
        }
        valueCountMap[v1] = count
        }
    }

    for key, value := range valueCountMap{
        similarity_score += key * value
    }
    fmt.Println("Similarity Index", similarity_score)

}
