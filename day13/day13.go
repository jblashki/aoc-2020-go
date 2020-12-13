package day13

import (
	"fmt"
	"strconv"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 13"
const inputFile = "./day13/timetable"

// RunDay runs Advent of Code Day 13 Puzzle
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

	bResult, err = b(verbose)
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: Program Result = %v\n", name, bResult)
	}
}

func a() (int, error) {
	values, err := filereader.ReadCSVStrings(inputFile)

	if err != nil {
		return 0, err
	}

	earliest, _ := strconv.Atoi(values[0])
	values = values[1:]

	lowest := -1
	lowestNum := -1
	for _, value := range values {
		if value != "x" {
			num, _ := strconv.Atoi(value)

			result := num - (earliest % num)

			if lowest == -1 || result < lowest {
				lowest = result
				lowestNum = num
			}
		}
	}

	return lowestNum * lowest, nil
}

type bus struct {
	id           int
	minutesAfter int
}

func b(verbose bool) (int, error) {
	values, err := filereader.ReadCSVStrings(inputFile)

	if err != nil {
		return 0, err
	}

	values = values[1:]
	buses := make([]bus, 0, len(values))
	for i, value := range values {
		if value != "x" {
			num, _ := strconv.Atoi(value)
			buses = append(buses, bus{id: num, minutesAfter: i})
		}
	}

	inc := buses[0].id
	i := inc
	curIdx := 1
	if verbose {
		fmt.Printf("Found departure @ %v: Bus ID = %v Minutes After = %v\n", i, buses[0].id, buses[0].minutesAfter)
	}
	for {
		bus := buses[curIdx]
		if (i+bus.minutesAfter)%bus.id == 0 {
			// found
			if verbose {
				fmt.Printf("Found departure @ %v: Bus ID = %v Minutes After = %v\n", i, bus.id, bus.minutesAfter)
			}
			if curIdx == len(buses)-1 {
				//found all
				return i, nil
			}
			curIdx++
			inc *= bus.id
		}
		i += inc
	}
}
