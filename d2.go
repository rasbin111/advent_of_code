package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func d2(){
    file, err := os.Open("./d2_input_data.txt")

    if (err != nil){
        fmt.Println("Error opening file", err)
        return 
    }
    scanner := bufio.NewScanner(file)
    // safeLineCount := 0
    for scanner.Scan(){
        line := scanner.Text() // get line
        // numbers := strings.Split(line, " ") // this also works
        numbers := strings.Fields(line)
        isIncreasing := true 
        isDecreasing := true 
        isSafe := true
        max := strconv.Atoi(numbers[0])
        min := strconv.Atoi(numbers[0])
        prevValue := strconv.Atoi(numbers[0]) 
        for _, value, err := strconv.Atoi(value) := range numbers{
            if max <= value{
                isIncreasing = false
            } else {
                max = value
            }
            if min >= value {
                isDecreasing = false
            } else {
                min = value
            }
            if !isIncreasing && !isDecreasing{
                isSafe = false
                break
            } else {
                if value - prevValue > 3{
                    isSafe = false
                    break
                } 
           }
            prevValue = value
            safeLineCount += 1
        }

        fmt.Println("Safe Line Count", safeLineCount)
    }

    
    
}   

