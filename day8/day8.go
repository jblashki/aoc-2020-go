package day8

import (
	"fmt"
	"strconv"
	"strings"

	filereader "github.com/jblashki/aoc-filereader-go"
)

const name = "Day 8"
const inputFile = "./day8/program"

// RunDay runs Advent of Code Day 8 Puzzle
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

type operation int

const (
	opNop operation = iota
	opAcc
	opJmp
)

type instruction struct {
	op    operation
	value int
}

func parseLine(line string) (inst instruction) {
	words := strings.Split(line, " ")
	switch words[0] {
	case "nop":
		inst.op = opNop
	case "acc":
		inst.op = opAcc
	case "jmp":
		inst.op = opJmp
	}

	if words[1][0] == '+' {
		words[1] = words[1][1:]
	}

	inst.value, _ = strconv.Atoi(words[1])

	return
}

func runCode(change int) (bool, int) {
	lines, err := filereader.ReadLines(inputFile)
	if err != nil {
		return false, 0
	}

	program := make([]instruction, 0, len(lines))
	n := 0
	for i := range lines {
		inst := parseLine(lines[i])

		switch inst.op {
		case opNop:
			n++
			if n == change {
				inst.op = opJmp
			}
		case opJmp:
			n++
			if n == change {
				inst.op = opNop
			}
		}

		program = append(program, inst)
	}

	acc := 0
	seen := make(map[int]bool)

	for progPos := 0; progPos < len(program); {
		if seen[progPos] {
			return false, acc
		}
		seen[progPos] = true
		inst := program[progPos]

		switch inst.op {
		case opJmp:
			progPos += inst.value

		case opAcc:
			acc += inst.value
			fallthrough
		case opNop:
			fallthrough
		default:
			progPos++
		}
	}

	return true, acc
}

func a() (int, error) {
	_, acc := runCode(0)

	return acc, nil
}

func b() (int, error) {
	ok := false
	n := 0
	acc := 0
	for !ok {
		ok, acc = runCode(n)
		n++
	}

	return acc, nil
}
