package year2015

import (
	"fmt"
)

type Day01 struct{}

const BasementFloor = -1

func processLine(line string, findBasement bool) (floor int, basementPos int) {
	for pos, char := range line {
		switch char {
		case '(':
			floor++
		case ')':
			floor--
		default:
			fmt.Printf("Character not recognized: %c\n", char)
		}

		if findBasement && floor == BasementFloor {
			basementPos = pos + 1
			break
		}
	}
	return floor, basementPos
}

func (d Day01) PartA(lines []string) any {
	if len(lines) == 0 {
		return 0
	}
	floor, _ := processLine(lines[0], false)
	return floor
}

func (d Day01) PartB(lines []string) any {
	if len(lines) == 0 {
		return 0
	}
	_, position := processLine(lines[0], true)
	return position
}
