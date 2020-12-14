package day14

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 14"
const fileInput = "./day14/mask"

// RunDay runs Advent of Code Day 14 Puzzle
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
	lines, err := filereader.ReadLines(fileInput)
	if err != nil {
		return 0, err
	}

	numbers := make(map[int]int)

	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	for _, line := range lines {
		lineValues := strings.Split(line, " = ")

		if lineValues[0] == "mask" {
			mask = lineValues[1]
		} else {
			mem := strings.Split(lineValues[0], "[")
			address, _ := strconv.Atoi(strings.Split(mem[1], "]")[0])
			value, _ := strconv.Atoi(lineValues[1])

			outValue := 0
			for i, maskChar := range mask {
				num := 0
				numCharVal := value >> (len(mask) - (i + 1)) & 0x01
				if maskChar != 'X' {
					maskCharVal := 0
					if maskChar == '1' {
						maskCharVal = 1
					}
					num = maskCharVal
				} else {
					num = numCharVal
				}

				outValue += num << (len(mask) - (i + 1))
			}
			numbers[address] = outValue
		}
	}

	total := 0
	for _, num := range numbers {
		total += num
	}

	return total, nil
}

func getAddresses(address int, mask string) []int {
	retAddresses := make([]int, 0)
	digitPos := make([]int, 0, len(mask))
	for i, c := range mask {
		if c == '1' {
			address |= 1 << (len(mask) - (i + 1))
		} else if c == 'X' {
			digitPos = append(digitPos, i)
		}
	}
	maxNum := math.Pow(2, float64(len(digitPos))) - 1
	maxNumBits := strconv.FormatInt(int64(maxNum), 2)

	a := 0
	for i := 0; i <= int(maxNum); i++ {
		a = address
		for j := 0; j < len(maxNumBits); j++ {
			v := i >> j & 0x01
			bitShift := len(mask) - (digitPos[j] + 1)
			if v == 1 {
				a |= 1 << bitShift
			} else {
				a &= ^(1 << bitShift)
			}

		}
		retAddresses = append(retAddresses, a)
	}

	return retAddresses
}

func b() (int, error) {
	lines, err := filereader.ReadLines(fileInput)
	if err != nil {
		return 0, err
	}

	numbers := make(map[int]int)

	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	for _, line := range lines {
		lineValues := strings.Split(line, " = ")

		if lineValues[0] == "mask" {
			mask = lineValues[1]
		} else {
			mem := strings.Split(lineValues[0], "[")
			address, _ := strconv.Atoi(strings.Split(mem[1], "]")[0])
			value, _ := strconv.Atoi(lineValues[1])

			addresses := getAddresses(address, mask)

			for _, a := range addresses {
				numbers[a] = value
			}
		}
	}

	total := 0
	for _, num := range numbers {
		total += num
	}

	return total, nil
}
