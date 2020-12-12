package day12

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 12"
const inputFile = "./day12/instructions"

// RunDay runs Advent of Code Day 12 Puzzle
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

type op int

const (
	opNorth   op = 'N'
	opSouth   op = 'S'
	opEast    op = 'E'
	opWest    op = 'W'
	opLeft    op = 'L'
	opRight   op = 'R'
	opForward op = 'F'
)

func (o op) String() string {
	switch o {
	case opNorth:
		return "North"
	case opSouth:
		return "South"
	case opEast:
		return "East"
	case opWest:
		return "West"
	case opLeft:
		return "Left"
	case opRight:
		return "Right"
	case opForward:
		return "Forward"
	}
	return ""
}

func a() (int, error) {
	answer, err := calcDistance(false)

	return answer, err
}

func b() (int, error) {
	answer, err := calcDistance(true)

	return answer, err
}

func calcDistance(partB bool) (int, error) {
	lines, err := filereader.ReadLines(inputFile)
	if err != nil {
		return 0, err
	}

	pos := complex(0, 0)

	way := complex(1, 0)
	if partB {
		way = complex(10, 1)
	}

	for _, line := range lines {
		op := op(line[0])
		num, _ := strconv.Atoi(line[1:])

		switch op {
		case opNorth:
			if !partB {
				pos += complex(float64(0), float64(num))
			}
			way += complex(float64(0), float64(num))

		case opSouth:
			if !partB {
				pos -= complex(float64(0), float64(num))
			}
			way -= complex(float64(0), float64(num))

		case opEast:
			if !partB {
				pos += complex(float64(num), float64(0))
			}
			way += complex(float64(num), float64(0))

		case opWest:
			if !partB {
				pos -= complex(float64(num), float64(0))
			}
			way -= complex(float64(num), float64(0))

		case opLeft:
			rotations := num / 90
			diff := (way - pos) * cmplx.Pow(0+1i, complex(float64(rotations), 0))
			way = pos + diff

		case opRight:
			rotations := num / 90
			diff := (way - pos) * cmplx.Pow(0-1i, complex(float64(rotations), 0))
			way = pos + diff

		case opForward:
			diff := (way - pos)
			pos += diff * complex(float64(num), 0)
			way = pos + diff
		}
	}

	answer := math.Abs(real(pos)) + math.Abs(imag(pos))
	answerInt := int(answer)

	return answerInt, nil
}
