package day2

import (
	//"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jblashki/aoc-filereader-go"
)

const name = "Day 2"

const input_file = "./day2/passwords"

func RunDay(verbose bool) {
	var aResult int
	var bResult int
	var err error

	if verbose {
		fmt.Printf("\n%v Output:\n", name)
	}

	aResult, err = a(verbose)
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

func a(verbose bool) (int, error) {
	lines, _ := filereader.ReadLines(input_file)

	validPasswordCount := 0
	for _, line := range lines {
		min, max, char, password := parseLine(line)

		if verbose {
			fmt.Printf("Password %q with %v occuring between %v and %v times: ", password, char, min, max)
		}

		if validPasswordA(password, min, max, char) {
			validPasswordCount++
		}
	}

	return validPasswordCount, nil
}

func b(verbose bool) (int, error) {
	lines, _ := filereader.ReadLines(input_file)

	validPasswordCount := 0
	for _, line := range lines {
		idx1, idx2, char, password := parseLine(line)

		if verbose {
			fmt.Printf("Password %q with %q @ index %v xor %v: ", password, char, idx1, idx2)
		}

		if validPasswordB(password, idx1, idx2, char) {
			if verbose {
				fmt.Printf("VALID\n")
			}
			validPasswordCount++
		} else if verbose {
			fmt.Printf("INVALID\n")
		}
	}

	return validPasswordCount, nil
}

func parseLine(line string) (num1 int, num2 int, char byte, password string) {
	s := strings.Split(line, ":")
	rules := strings.Split(s[0], " ")
	count := strings.Split(rules[0], "-")
	num1, _ = strconv.Atoi(count[0])
	num2, _ = strconv.Atoi(count[1])

	char = rules[1][0]

	password = strings.Trim(s[1], " ")

	return
}

func validPasswordA(password string, min int, max int, char byte) bool {
	count := 0
	for i := 0; i < len(password); i++ {
		if char == password[i] {
			count++
		}
	}

	if count < min || count > max {
		return false
	}

	return true
}

func validPasswordB(password string, idx1 int, idx2 int, char byte) bool {

	count := 0
	if password[idx1-1] == char {
		count++
	}

	if password[idx2-1] == char {
		count++
	}

	if count == 1 {
		return true
	}

	return false
}
