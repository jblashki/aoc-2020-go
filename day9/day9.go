package day9

import (
	"errors"
	"fmt"
	"sort"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 9"
const inputFile = "./day9/numbers"

// RunDay runs Advent of Code Day 9 Puzzle
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

	bResult, err = b(aResult)
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

	preamble := 25

	for i := preamble; i < len(nums); i++ {
		found := false
		for j := i - preamble; j < i-1; j++ {
			for k := j + 1; k < i; k++ {
				if nums[j]+nums[k] == nums[i] {
					found = true
				}
			}
		}

		if !found {
			return nums[i], nil
		}
	}

	return 0, errors.New("No Number Matched")
}

func b(magicNum int) (int, error) {
	nums, err := filereader.ReadAllInts(inputFile)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(nums); i++ {
		answer := nums[i]
		list := make([]int, 0, len(nums))
		list = append(list, nums[i])
		for j := i + 1; j < len(nums); j++ {
			answer += nums[j]
			list = append(list, nums[j])
			if answer == magicNum {
				sort.Ints(list)
				answer = list[0] + list[len(list)-1]
				return answer, nil
			}
		}
	}
	return 0, errors.New("Couldn't find")
}
