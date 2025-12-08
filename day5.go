package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type ValidIdsRange struct {
	start int64
	end   int64
}

func readFromFile(filepath string) []string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func validIdsPerLine(line string) ValidIdsRange {

	start, err := strconv.ParseInt(strings.Split(line, "-")[0], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	end, err := strconv.ParseInt(strings.Split(line, "-")[1], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	rangesPerLine := ValidIdsRange{
		start: start,
		end:   end,
	}

	return rangesPerLine
}

func validIdsListStruct(lines []string) []ValidIdsRange {
	ranges := []ValidIdsRange{}

	for _, line := range lines {
		if line == "" {
			break
		}
		ragesPerLine := validIdsPerLine(line)

		ranges = append(ranges, ragesPerLine)

	}

	return ranges
}

func IdsList(lines []string) []int64 {
	ids := []int64{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "-") {
			continue
		}
		num, err := strconv.ParseInt(line, 10, 64)

		if err != nil {
			log.Fatal(err)
		}

		ids = append(ids, num)
	}

	return ids
}

func D5P1(lines []string) {

	validIds := validIdsListStruct(lines)

	ids := IdsList(lines)

	freshIds := []int64{}

	for _, id := range ids {
		for _, vi := range validIds {
			if id >= vi.start && id <= vi.end {
				if slices.Contains(freshIds, id) {
					continue
				} else {
					freshIds = append(freshIds, id)
				}
			}
		}
	}

	fmt.Println("Ids: ", len(ids))
	fmt.Println("Number of fresh ingredients: ", len(freshIds))
}

func D5P2(lines []string) {
	ranges := validIdsListStruct(lines)

	slices.SortFunc(ranges, func(a, b ValidIdsRange) int {
		return int(a.start - b.start)
	})

	var freshIds int64 = 0
	currentStart := ranges[0].start
	currentEnd := ranges[0].end

	for _, r := range ranges[1:] {
		if r.start <= currentEnd+1 {
			if r.end > currentEnd {
				currentEnd = r.end
			}
		} else {
			freshIds += (currentEnd - currentStart + 1)
			currentStart = r.start
			currentEnd = r.end
		}
	}

	freshIds += (currentEnd - currentStart + 1)

	fmt.Println("Number of fresh ingredients:", freshIds)

}

func Day5() {
	lines := readFromFile("inputs/day5.txt")
	D5P1(lines) // 513
	D5P2(lines) // 339668510830757
}
