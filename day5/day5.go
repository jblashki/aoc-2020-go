package day5

import (
	"fmt"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 5"
const inputFile = "./day5/seats"

// RunDay runs Advent of Code Day 5 Puzzle
func RunDay(verbose bool) {
	var aResult int
	var bResult int
	var err error

	if verbose {
		fmt.Printf("\n%v Output:\n", name)
	}

	aResult, bResult, err = calc()
	if err != nil {
		fmt.Printf("%va: **** Error: %q ****\n", name, err)
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%va: Program Result = %v\n", name, aResult)
		fmt.Printf("%vb: Program Result = %v\n", name, bResult)
	}
}

func calc() (int, int, error) {
	lines, _ := filereader.ReadLines(inputFile)

	seats := make([][]bool, 128)
	for i := 0; i < 128; i++ {
		seats[i] = make([]bool, 8)
	}

	highest := -1
	for _, line := range lines {
		rowMin := 0
		rowMax := 127
		colMin := 0
		colMax := 7
		for _, c := range []rune(line) {
			rowMin, rowMax, colMin, colMax = processChar(c, rowMin, rowMax, colMin, colMax)
		}

		seats[rowMin][colMin] = true

		num := (rowMin * 8) + colMin

		if highest == -1 || num > highest {
			highest = num
		}
	}

	mySeat := -1
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if !seats[i][j] {
				seatID := (i * 8) + j
				prevSeatI := (seatID - 1) / 8
				prevSeatJ := (seatID - 1) % 8
				nextSeatI := (seatID + 1) / 8
				nextSeatJ := (seatID + 1) % 8

				if prevSeatI >= 0 && prevSeatI <= 127 && prevSeatJ >= 0 && prevSeatJ <= 7 &&
					nextSeatI >= 0 && nextSeatI <= 127 && nextSeatJ >= 0 && nextSeatJ <= 7 {
					if seats[prevSeatI][prevSeatJ] && seats[nextSeatI][nextSeatJ] {
						mySeat = seatID
					}
				}
			}
		}
	}

	return highest, mySeat, nil
}

func processChar(char rune, rowMin int, rowMax int, colMin int, colMax int) (int, int, int, int) {
	rows := rowMax - rowMin
	cols := colMax - colMin
	switch char {
	case 'F':
		if rows%2 != 0 {
			rowMax = (rows / 2) + rowMin
		}
	case 'B':
		if rows%2 != 0 {
			rowMin = (rows / 2) + 1 + rowMin
		} else {
			rowMin = (rows / 2) + rowMin
		}
	case 'L':
		if cols%2 != 0 {
			colMax = (cols / 2) + colMin
		}
	case 'R':
		if cols%2 != 0 {
			colMin = (cols / 2) + 1 + colMin
		} else {
			colMin = (cols / 2) + colMin
		}
	}

	return rowMin, rowMax, colMin, colMax
}
