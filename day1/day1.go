package day1

import (
	"errors"
	"fmt"

	"github.com/jblashki/aoc-filereader-go"
)

const name = "Day 1"

const input_file = "./day1/numbers"

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
		fmt.Printf("%va: Result = %v\n", name, aResult)
	}

	bResult, err = b()
	if err != nil {
		fmt.Printf("%vb: **** Error: %q ****\n", name, err)
	} else {
		fmt.Printf("%vb: Result = %v\n", name, bResult)
	}
}

func a() (int, error) {
	numbers, err := filereader.ReadAllInts(input_file)
	if err != nil {
		errormsg := fmt.Sprintf("Failed to read from #q: %v", input_file, err)
		return 0, errors.New(errormsg)
	}

	for i := 0; i < len(numbers); i++ {
		numA := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			numB := numbers[j]

			if numB+numA == 2020 {
				return numA * numB, nil
			}
		}
	}
	return 0, errors.New("Couldn't find numbers")
}

func b() (int, error) {
	numbers, err := filereader.ReadAllInts(input_file)
	if err != nil {
		errormsg := fmt.Sprintf("Failed to read from #q: %v", input_file, err)
		return 0, errors.New(errormsg)
	}

	for i := 0; i < len(numbers); i++ {
		numA := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			numB := numbers[j]
			for k := j + 1; k < len(numbers); k++ {
				numC := numbers[k]
				if numB+numA+numC == 2020 {
					return numA * numB * numC, nil
				}
			}
		}
	}
	return 0, errors.New("Couldn't find numbers")
}
