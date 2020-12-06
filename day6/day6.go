package day6

import (
	"fmt"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 6"

const inputFile = "./day6/answers"

// RunDay runs Advent of Code Day 6 Puzzle
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
	lines, err := filereader.ReadLines(inputFile)
	if err != nil {
		return 0, err
	}

	answerCount := make(map[rune]int)
	line := ""
	totalCount := 0
	for _, line = range lines {
		if line == "" {
			for _, count := range answerCount {
				if count > 0 {
					totalCount++
				}
			}
			answerCount = make(map[rune]int)
		}

		for j := 0; j < len(line); j++ {
			answerCount[rune(line[j])]++
		}
	}

	for _, count := range answerCount {
		if count > 0 {
			totalCount++
		}
	}

	return totalCount, nil
}

func b() (int, error) {
	lines, err := filereader.ReadLines(inputFile)
	if err != nil {
		return 0, err
	}

	answerCount := make(map[rune]int)
	line := ""
	totalCount := 0
	groupSize := 0
	for _, line = range lines {
		if line == "" {
			for _, count := range answerCount {
				if count == groupSize {
					totalCount++
				}
			}
			answerCount = make(map[rune]int)
			groupSize = 0
		} else {
			groupSize++
		}
		for j := 0; j < len(line); j++ {
			answerCount[rune(line[j])]++
		}
	}

	for _, count := range answerCount {
		if count == groupSize {
			totalCount++
		}
	}

	return totalCount, nil
}
