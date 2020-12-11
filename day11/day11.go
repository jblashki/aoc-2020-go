package day11

import (
	"fmt"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 11"
const inputFile = "./day11/seats"

// RunDay runs Advent of Code Day 11 Puzzle
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

type direction int

const (
	N direction = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func canSeeOccupiedSeat(seats [][]rune, i int, j int, d direction, maxRadius int, curRadius int) bool {
	switch d {
	case N:
		if i <= 0 {
			return false
		}
		i--
	case NE:
		if i <= 0 || j >= len(seats[i])-1 {
			return false
		}
		i--
		j++
	case E:
		if j >= len(seats[i])-1 {
			return false
		}
		j++
	case SE:
		if i >= len(seats)-1 || j >= len(seats[i])-1 {
			return false
		}
		i++
		j++
	case S:
		if i >= len(seats)-1 {
			return false
		}
		i++
	case SW:
		if i >= len(seats)-1 || j <= 0 {
			return false
		}
		i++
		j--
	case W:
		if j <= 0 {
			return false
		}
		j--
	case NW:
		if i <= 0 || j <= 0 {
			return false
		}
		i--
		j--
	}
	curRadius++

	if seats[i][j] == '.' {
		if maxRadius != -1 && curRadius == maxRadius {
			return false
		} else {
			return canSeeOccupiedSeat(seats, i, j, d, maxRadius, curRadius)
		}
	} else if seats[i][j] == '#' {
		return true
	}

	return false
}

func countSurroundingOccupiedSeats(seats [][]rune, i int, j int, partA bool) int {
	count := 0

	for d := N; d <= NW; d++ {
		if partA {
			if canSeeOccupiedSeat(seats, i, j, d, 1, 0) {
				count++
			}
		} else {
			if canSeeOccupiedSeat(seats, i, j, d, -1, 0) {
				count++
			}
		}
	}

	return count
}

func processSeats(seats [][]rune, partA bool) ([][]rune, bool) {
	newSeats := make([][]rune, len(seats))

	changed := false
	for i := range seats {
		newSeats[i] = make([]rune, len(seats[i]))
		for j := range seats[i] {
			if seats[i][j] == 'L' {
				// Check surrounding seats
				count := countSurroundingOccupiedSeats(seats, i, j, partA)
				if count == 0 {
					newSeats[i][j] = '#'
					changed = true
				} else {
					newSeats[i][j] = seats[i][j]
				}
			} else if seats[i][j] == '#' {
				// Check surrounding seats
				count := countSurroundingOccupiedSeats(seats, i, j, partA)
				surroundingSeatCount := 4
				if !partA {
					surroundingSeatCount = 5
				}
				if count >= surroundingSeatCount {
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

	changed := true
	for changed == true {
		seats, changed = processSeats(seats, true)
	}

	total := totalOccupied(seats)

	return total, nil
}

func b() (int, error) {
	lines, _ := filereader.ReadLines(inputFile)

	seats := make([][]rune, len(lines))
	for i, line := range lines {
		seats[i] = []rune(line)
	}

	changed := true
	for changed == true {
		seats, changed = processSeats(seats, false)
	}

	total := totalOccupied(seats)

	return total, nil
}
