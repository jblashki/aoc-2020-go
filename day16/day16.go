package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 16"
const inputFile = "./day16/tickets"

// RunDay runs Advent of Code Day 16 Puzzle
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

type validRange struct {
	from int
	to   int
}

type rule struct {
	name   string
	ranges [2]validRange
}

type document struct {
	rules []rule

	myTicket      []int
	nearbyTickets [][]int
}

type state int

const (
	stateRules state = iota
	stateMy
	stateNearby
)

func readDocument(inputFile string) (*document, error) {
	d := new(document)
	d.rules = make([]rule, 0)
	d.myTicket = make([]int, 0)
	d.nearbyTickets = make([][]int, 0)

	lines, err := filereader.ReadLines(inputFile)
	if err != nil {
		return nil, err
	}

	s := stateRules
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		switch s {
		case stateRules:
			if line == "" {
				s = stateMy
				i++
			} else {
				r1 := strings.Split(line, ": ")
				ruleName := r1[0]
				r2 := strings.Split(r1[1], " or ")
				r3 := strings.Split(r2[0], "-")
				r4 := strings.Split(r2[1], "-")

				f1, _ := strconv.Atoi(r3[0])
				f2, _ := strconv.Atoi(r4[0])

				t1, _ := strconv.Atoi(r3[1])
				t2, _ := strconv.Atoi(r4[1])

				rx := rule{name: ruleName, ranges: [2]validRange{
					{from: f1, to: t1},
					{from: f2, to: t2},
				}}

				d.rules = append(d.rules, rx)
			}

		case stateMy:
			if line == "" {
				s = stateNearby
				i++
			} else {
				csvValues := strings.Split(line, ",")
				for _, value := range csvValues {
					num, _ := strconv.Atoi(value)
					d.myTicket = append(d.myTicket, num)

				}
			}

		case stateNearby:
			numbers := make([]int, 0)
			csvValues := strings.Split(line, ",")
			for _, value := range csvValues {
				num, _ := strconv.Atoi(value)
				numbers = append(numbers, num)

			}
			d.nearbyTickets = append(d.nearbyTickets, numbers)
		}
	}

	return d, nil
}

func validateRule(r rule, v int) bool {
	if (v >= r.ranges[0].from && v <= r.ranges[0].to) ||
		(v >= r.ranges[1].from && v <= r.ranges[1].to) {
		return true
	}

	return false
}

func a() (int, error) {
	d, err := readDocument(inputFile)
	if err != nil {
		return 0, err
	}

	invalidValues := make([]int, 0)

	for _, ticket := range d.nearbyTickets {
		for _, num := range ticket {
			valid := false
			for _, r := range d.rules {
				if validateRule(r, num) {
					valid = true
				}
			}
			if !valid {
				invalidValues = append(invalidValues, num)
			}
		}
	}

	total := 0
	for _, num := range invalidValues {
		total += num
	}

	return total, nil
}

func b() (int, error) {
	d, err := readDocument(inputFile)
	if err != nil {
		return 0, err
	}

	validTickets := make([][]int, 0)
	for _, ticket := range d.nearbyTickets {
		validTicket := true
		for _, num := range ticket {
			validNum := false
			for _, r := range d.rules {
				if validateRule(r, num) {
					validNum = true
				}
			}
			if !validNum {
				validTicket = false
				break
			}
		}
		if validTicket {
			validTickets = append(validTickets, ticket)
		}
	}

	valid := make([][]bool, len(d.rules))

	for i, r := range d.rules {
		valid[i] = make([]bool, len(validTickets[0]))
		for j := 0; j < len(validTickets[0]); j++ {
			valid[i][j] = true
		}
		for _, ticket := range validTickets {
			for j, num := range ticket {
				if !validateRule(r, num) {
					valid[i][j] = false
				}
			}
		}
	}

	want := len(d.rules)
	ruleIndex := make([]int, want)
	found := 0
	for found < want {
		for r := 0; r < len(valid); r++ {
			lastIdx := -1
			count := 0
			for idx := 0; idx < len(valid[r]); idx++ {
				if valid[r][idx] {
					count++
					lastIdx = idx
				}
			}
			if count == 1 {
				for r2 := 0; r2 < len(valid); r2++ {
					valid[r2][lastIdx] = false
				}
				ruleIndex[r] = lastIdx
				found++
			}
		}
	}

	total := 1
	for i := 0; i < want; i++ {
		match, _ := regexp.MatchString("departure.*", d.rules[i].name)
		if match {
			total *= d.myTicket[ruleIndex[i]]
		}
	}

	return total, nil
}
