package day10

import (
	"fmt"
	"sort"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 10"

const inputFile = "./day10/jolts"

// RunDay runs Advent of Code Day 10 Puzzle
func RunDay(verbose bool) {
	var aResult int
	var bResult int
	var err error

	if verbose {
		fmt.Printf("\n%v Output:\n", name)
	}

	aResult, err = a()
	if err != nil {
		fmt.Printf("%va: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%va: Program Result = %v\n", name, aResult)
	}

	bResult, err = b()
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: Program Result = %v\n", name, bResult)
	}
}

func a() (int, error) {
	nums, err := filereader.ReadAllInts(inputFile)
	if err != nil {
		return 0, err
	}

	sort.Ints(nums)

	ratingDiffCount := make([]int, 3)
	for i, num := range nums {
		idx := 0
		if i == 0 {
			idx = num - 1
		} else {
			idx = num - nums[i-1] - 1
		}
		ratingDiffCount[idx]++
	}

	// Last rating
	ratingDiffCount[2]++

	answer := ratingDiffCount[0] * ratingDiffCount[2]

	return answer, nil
}

func b() (int, error) {
	nums, err := filereader.ReadAllInts(inputFile)
	if err != nil {
		return 0, err
	}

	nums = append(nums, 0)
	sort.Ints(nums)

	countMap := make(map[int]int, len(nums)+1)

	for _, num := range nums {
		if num == 0 {
			countMap[num] = 1
		} else {
			countMap[num] = 0
		}
	}

	for _, num := range nums {
		i, ok := countMap[num+1]
		if ok {
			countMap[num+1] = i + countMap[num]
		}

		i, ok = countMap[num+2]
		if ok {
			countMap[num+2] = i + countMap[num]
		}

		i, ok = countMap[num+3]
		if ok {
			countMap[num+3] = i + countMap[num]
		}
	}

	finalNum := nums[len(nums)-1]

	count := countMap[finalNum]

	return count, nil
}
