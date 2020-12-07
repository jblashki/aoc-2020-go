package day7

import (
	"fmt"
	"strconv"
	"strings"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 7"

const inputFile = "./day7/bags"

var fileStrings []string

// RunDay runs Advent of Code Day 7 Puzzle
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

func bagContainsColour(mainBag *bag, colourName string) bool {
	for _, subBag := range mainBag.bags {
		if subBag.name == colourName {
			return true
		}
		if bagContainsColour(subBag, colourName) == true {
			return true
		}
	}

	return false
}

func a(verbose bool) (int, error) {
	bags, err := loadBags()
	if err != nil {
		return 0, err
	}

	count := 0
	for _, bag := range bags {
		if bagContainsColour(bag, "shiny gold") {
			if verbose {
				fmt.Printf("%q has at least 1 shiny gold bag\n", bag.name)
			}
			count++
		}
	}

	return count, nil
}

func countBag(b *bag, verbose bool) int {
	count := 1

	for i, b2 := range b.bags {
		count2 := countBag(b2, verbose)

		count += (b.bagCount[i] * count2)
	}

	if verbose {
		fmt.Printf("Bag: %q contains %v bags\n", b.name, count-1)
	}

	return count
}

func b(verbose bool) (int, error) {
	bags, err := loadBags()
	if err != nil {
		return 0, err
	}

	b := bags["shiny gold"]
	count := countBag(b, verbose)

	count--

	return count, nil
}

type state int

const (
	statePrimaryColour state = iota
	stateContains
	stateSubCount
	stateNoBags
	stateSubColour
)

type bag struct {
	name     string
	bags     []*bag
	bagCount []int
}

func loadBags() (map[string]*bag, error) {
	err := readFile()
	if err != nil {
		return nil, err
	}

	returnMap := make(map[string]*bag)

	state := statePrimaryColour
	primaryColour := ""
	subCount := 0
	subColour := ""
	for token := getNextToken(); token != ""; token = getNextToken() {
		switch state {
		case statePrimaryColour:
			switch token {
			case "bags":
				newBag, ok := returnMap[primaryColour]
				if !ok {
					//fmt.Printf("	%v: %q\n", subCount, subColour)
					newBag = new(bag)
					newBag.name = primaryColour
					newBag.bags = nil
					newBag.bagCount = nil
					returnMap[primaryColour] = newBag
				}

				getNextToken() // Read off "contain"
				state = stateContains
			default:
				primaryColour += " " + token
				primaryColour = strings.TrimSpace(primaryColour)
			}

		case stateContains:
			switch token {
			case "no":
				getNextToken() // Read off "other"
				getNextToken() // Read off "bags."

				//fmt.Printf("%q -- NO OTHER BAGS\n", primaryColour)
				primaryColour = ""
				subCount = 0
				subColour = ""
				state = statePrimaryColour
			default:
				subCount, err = strconv.Atoi(token)
				if err != nil {
					return nil, err
				}
				state = stateSubColour
			}

		case stateSubColour:
			switch token {
			case "bag,":
				fallthrough
			case "bags,":
				// bag description
				//fmt.Printf("	%v: %q\n", subCount, subColour)
				newBag, ok := returnMap[subColour]
				if !ok {
					//fmt.Printf("	%v: %q\n", subCount, subColour)

					newBag = new(bag)
					newBag.name = subColour
					newBag.bags = nil
					newBag.bagCount = nil
					returnMap[subColour] = newBag
				}
				bag := returnMap[primaryColour]
				bag.bags = append(bag.bags, newBag)
				bag.bagCount = append(bag.bagCount, subCount)
				returnMap[primaryColour] = bag

				//fmt.Printf("bag.bags len = %v\n", bag.bags[0].name)

				subCount = 0
				subColour = ""
				state = stateContains

			case "bag.":
				fallthrough
			case "bags.":
				// Last sub bag description
				// fmt.Printf("	%v: %q\n", subCount, subColour)
				// fmt.Printf("%q: ^^^^^^^^\n", primaryColour)

				newBag, ok := returnMap[subColour]
				if !ok {
					newBag = new(bag)
					newBag.name = subColour
					newBag.bags = nil
					newBag.bagCount = nil
					returnMap[subColour] = newBag
				}
				bag := returnMap[primaryColour]
				bag.bags = append(bag.bags, newBag)
				bag.bagCount = append(bag.bagCount, subCount)

				primaryColour = ""
				subCount = 0
				subColour = ""
				state = statePrimaryColour

			default:
				subColour += " " + token
				subColour = strings.TrimSpace(subColour)
			}
		}
		//fmt.Printf("Token: %v\n", token)
	}

	return returnMap, nil
}

func readFile() error {
	lines, err := filereader.ReadLines(inputFile)
	if err != nil {
		return err
	}

	fileStrings = make([]string, 0)

	for _, line := range lines {
		lineStrings := strings.Split(line, " ")
		fileStrings = append(fileStrings, lineStrings...)
	}

	return nil
}

func getNextToken() string {
	if len(fileStrings) <= 0 {
		return ""
	}

	s := fileStrings[0]
	fileStrings = fileStrings[1:]
	return s
}
