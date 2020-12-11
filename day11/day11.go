package day11

import (
	"errors"
	"fmt"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 11"
const inputFile = "./day11/seats"

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

func checkSeat(seats [][]rune, i int, j int) bool {
	if i < 0 {
		return false
	}

	if i >= len(seats) {
		return false
	}

	if j < 0 {
		return false
	}

	if j >= len(seats[i]) {
		return false
	}

	if seats[i][j] == '#' {
		return true
	}

	return false
}

func countOccupiedSeats(seats [][]rune, i int, j int) int {
	count := 0

	if checkSeat(seats, i-1, j-1) {
		count++
	}

	if checkSeat(seats, i-1, j) {
		count++
	}

	if checkSeat(seats, i-1, j+1) {
		count++
	}

	if checkSeat(seats, i, j-1) {
		count++
	}

	if checkSeat(seats, i, j+1) {
		count++
	}

	if checkSeat(seats, i+1, j-1) {
		count++
	}

	if checkSeat(seats, i+1, j) {
		count++
	}

	if checkSeat(seats, i+1, j+1) {
		count++
	}

	return count
}

func processSeats(seats [][]rune) ([][]rune, bool) {
	newSeats := make([][]rune, len(seats))

	changed := false
	for i := range seats {
		newSeats[i] = make([]rune, len(seats[i]))
		for j := range seats[i] {
			if seats[i][j] == 'L' {
				// Check surrounding seats
				count := countOccupiedSeats(seats, i, j)
				if count == 0 {
					newSeats[i][j] = '#'
					changed = true
				} else {
					newSeats[i][j] = seats[i][j]
				}
			} else if seats[i][j] == '#' {
				// Check surrounding seats
				count := countOccupiedSeats(seats, i, j)
				if count >= 4 {
					newSeats[i][j] = 'L'
					changed = true
				} else {
					newSeats[i][j] = seats[i][j]
				}
			} else {
				newSeats[i][j] = seats[i][j]
			}
		}
	}

	return newSeats, changed
}

func totalOccupied(seats [][]rune) int {
	count := 0
	for i := range seats {
		for j := range seats[i] {
			if seats[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func a() (int, error) {
	lines, _ := filereader.ReadLines(inputFile)

	seats := make([][]rune, len(lines))
	for i, line := range lines {
		seats[i] = []rune(line)
	}

	//fmt.Printf("%v\n", seats)
	changed := true
	for changed == true {
		seats, changed = processSeats(seats)
	}

	total := totalOccupied(seats)

	// for i := range seats {
	// 	for j := range seats[i] {
	// 		fmt.Printf("%c ", seats[i][j])
	// 	}
	// 	fmt.Printf("\n")
	// }
	return total, nil
}

func b() (int, error) {
	return 0, errors.New("Not Complete Yet")
}
