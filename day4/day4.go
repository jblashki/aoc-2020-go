package day4

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 4"
const inputFile = "./day4/passports"

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

// RunDay runs Advent of Code Day 4 Puzzle
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
	validPassportCount := 0

	lines, err := filereader.ReadLines(inputFile)
	if err != nil {
		return 0, err
	}

	total := 0
	for p, readLines, err := readNextPassport(lines); len(lines) > 0; p, readLines, err = readNextPassport(lines) {
		if err != nil {
			return 0, err
		}
		lines = lines[readLines:]

		if validPassport(p) {
			validPassportCount++
		}
		total++
	}

	if verbose {
		fmt.Printf("%v/%v passwords validated\n", validPassportCount, total)
	}

	return validPassportCount, nil
}

func b(verbose bool) (int, error) {
	validPassportCount := 0

	lines, err := filereader.ReadLines(inputFile)
	if err != nil {
		return 0, err
	}

	total := 0
	for p, readLines, err := readNextPassport(lines); len(lines) > 0; p, readLines, err = readNextPassport(lines) {
		if err != nil {
			return 0, err
		}
		lines = lines[readLines:]

		if validPassport2(p) {
			validPassportCount++
		}
		total++
	}

	if verbose {
		fmt.Printf("%v/%v passwords validated\n", validPassportCount, total)
	}

	return validPassportCount, nil
}

func validPassport(p passport) bool {
	if p.byr == "" {
		return false
	}
	if p.iyr == "" {
		return false
	}
	if p.eyr == "" {
		return false
	}
	if p.hgt == "" {
		return false
	}
	if p.hcl == "" {
		return false
	}
	if p.ecl == "" {
		return false
	}
	if p.pid == "" {
		return false
	}
	return true
}

func validateByr(byr string) bool {
	if byr == "" {
		return false
	}

	if len(byr) != 4 {
		return false
	}
	byrNum, err := strconv.Atoi(byr)
	if err != nil {
		return false
	}
	if byrNum < 1920 || byrNum > 2002 {
		return false
	}
	return true
}

func validateIyr(iyr string) bool {
	if iyr == "" {
		return false
	}

	if len(iyr) != 4 {
		return false
	}
	iyrNum, err := strconv.Atoi(iyr)
	if err != nil {
		return false
	}
	if iyrNum < 2010 || iyrNum > 2020 {
		return false
	}
	return true
}

func validateEyr(eyr string) bool {
	if eyr == "" {
		return false
	}

	if len(eyr) != 4 {
		return false
	}
	eyrNum, err := strconv.Atoi(eyr)
	if err != nil {
		return false
	}
	if eyrNum < 2020 || eyrNum > 2030 {
		return false
	}

	return true
}

func validateHgt(hgt string) bool {
	if hgt == "" {
		return false
	}

	hgtMetric := hgt[len(hgt)-2:]
	if !(hgt[len(hgt)-2] == 'c' && hgt[len(hgt)-1] == 'm') && !(hgt[len(hgt)-2] == 'i' && hgt[len(hgt)-1] == 'n') {
		return false
	}
	hgtNum, err := strconv.Atoi(hgt[:len(hgt)-2])
	if err != nil {
		return false
	}
	switch hgtMetric {
	case "cm":
		if hgtNum < 150 || hgtNum > 193 {
			return false
		}
	case "in":
		if hgtNum < 59 || hgtNum > 76 {
			return false
		}
	}
	return true
}

func validateHcl(hcl string) bool {
	if hcl == "" {
		return false
	}

	if len(hcl) != 7 {
		return false
	}
	if hcl[0] != '#' {
		return false
	}
	for i := 1; i < len(hcl); i++ {
		switch hcl[i] {
		case '0':
		case '1':
		case '2':
		case '3':
		case '4':
		case '5':
		case '6':
		case '7':
		case '8':
		case '9':
		case 'a':
		case 'b':
		case 'c':
		case 'd':
		case 'e':
		case 'f':
		default:
			return false
		}
	}

	return true
}

func validateEcl(ecl string) bool {
	if ecl == "" {
		return false
	}

	switch ecl {
	case "amb":
	case "blu":
	case "brn":
	case "gry":
	case "grn":
	case "hzl":
	case "oth":
	default:
		return false
	}

	return true
}

func validatePid(pid string) bool {
	if pid == "" {
		return false
	}

	if len(pid) != 9 {
		return false
	}
	for i := 0; i < len(pid); i++ {
		switch pid[i] {
		case '0':
		case '1':
		case '2':
		case '3':
		case '4':
		case '5':
		case '6':
		case '7':
		case '8':
		case '9':
		default:
			return false
		}
	}

	return true
}

func validPassport2(p passport) bool {
	switch {
	case !validateByr(p.byr):
		return false

	case !validateIyr(p.iyr):
		return false

	case !validateEyr(p.eyr):
		return false

	case !validateHgt(p.hgt):
		return false

	case !validateHcl(p.hcl):
		return false

	case !validateEcl(p.ecl):
		return false

	case !validatePid(p.pid):
		return false
	}

	return true
}

func (p passport) String() string {
	return fmt.Sprintf("{\n\tbyr = %v\n\tiyr = %v\n\teyr = %v\n\thgt = %v\n\thcl = %v\n\tecl = %v\n\tpid = %v\n\tcid = %v\n}",
		p.byr, p.iyr, p.eyr, p.hgt, p.hcl, p.ecl, p.pid, p.cid)
}

func readNextPassport(lines []string) (passport, int, error) {
	var p passport
	linesRead := 0

	for i := 0; i < len(lines); i++ {
		linesRead++
		if lines[i] == "" {
			break
		}
		// Parse Line here
		fields := strings.Split(lines[i], " ")
		for _, field := range fields {
			val := strings.Split(field, ":")
			if val[0] == "" {
				errormsg := fmt.Sprintf("Invalid key value \"\" on line %v\n", i)
				return p, linesRead, errors.New(errormsg)
			}
			switch val[0] {
			case "byr":
				p.byr = val[1]

			case "iyr":
				p.iyr = val[1]

			case "eyr":
				p.eyr = val[1]

			case "hgt":
				p.hgt = val[1]

			case "hcl":
				p.hcl = val[1]

			case "ecl":
				p.ecl = val[1]

			case "pid":
				p.pid = val[1]

			case "cid":
				p.cid = val[1]
			}
		}
	}

	return p, linesRead, nil
}
