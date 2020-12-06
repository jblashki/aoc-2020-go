package day3

import (
	"fmt"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 3"
const inputFile = "./day3/map"

// RunDay runs Advent of Code Day 3 Puzzle
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

func countTrees(slopeMap []string, right int, down int) (count int) {
	for xpos, ypos := 0, 0; ypos < len(slopeMap); xpos, ypos = xpos+right, ypos+down {
		if xpos >= len(slopeMap[ypos]) {
			xpos = xpos - len(slopeMap[ypos])
		}

		if slopeMap[ypos][xpos] == '#' {
			count++
		}
	}

	return
}

func a() (int, error) {
	slopeMap, err := filereader.ReadLines(inputFile)
	if err != nil {
		return 0, err
	}

	count := countTrees(slopeMap, 3, 1)

	return count, nil
}

func b() (int, error) {
	slopeMap, err := filereader.ReadLines(inputFile)
	if err != nil {
		return 0, err
	}

	countRun1 := countTrees(slopeMap, 1, 1)
	countRun2 := countTrees(slopeMap, 3, 1)
	countRun3 := countTrees(slopeMap, 5, 1)
	countRun4 := countTrees(slopeMap, 7, 1)
	countRun5 := countTrees(slopeMap, 1, 2)

	totalCount := countRun1 * countRun2 * countRun3 * countRun4 * countRun5

	return totalCount, nil
}
