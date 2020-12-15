package day15

import (
	"fmt"
)

const name = "Day 15"

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
	numbers := make([]int, 0)
	numbers = append(numbers, 1)
	numbers = append(numbers, 0)
	numbers = append(numbers, 15)
	numbers = append(numbers, 2)
	numbers = append(numbers, 10)
	numbers = append(numbers, 13)

	spoken := make(map[int]int, 0)

	turn := 0
	lastNum := -1
	num := -1
	for {
		if turn < len(numbers) {
			num = numbers[turn]
			spoken[num] = turn + 1
		} else {
			lastTurn, ok := spoken[lastNum]
			if !ok {
				num = 0
			} else {
				num = turn - lastTurn
			}
			spoken[lastNum] = turn
		}
		turn++

		lastNum = num

		if turn == 2020 {
			return num, nil
		}
	}
}

func b() (int, error) {
	numbers := make([]int, 0)
	numbers = append(numbers, 1)
	numbers = append(numbers, 0)
	numbers = append(numbers, 15)
	numbers = append(numbers, 2)
	numbers = append(numbers, 10)
	numbers = append(numbers, 13)

	spoken := make(map[int]int, 0)

	turn := 0
	lastNum := -1
	num := -1
	for {
		if turn < len(numbers) {
			num = numbers[turn]
			spoken[num] = turn + 1
		} else {
			lastTurn, ok := spoken[lastNum]
			if !ok {
				num = 0
			} else {
				num = turn - lastTurn
			}
			spoken[lastNum] = turn
		}
		turn++

		lastNum = num

		if turn == 30000000 {
			return num, nil
		}
	}
}
