package year2015

import (
	"fmt"
)

type Day01 struct{}

func (p Day01) PartA(lines []string) any {
	floorNumber := 0
	for _, line := range lines {
		for _, character := range line {
			switch character {
			case '(':
				floorNumber++
			case ')':
				floorNumber--
			default:
				fmt.Printf("Character not recognized: %s", string(character))
			}
		}
	}
	return floorNumber
}

func (p Day01) PartB(lines []string) any {
	floorNumber := 0
	for _, line := range lines {
		for pos, character := range line {
			switch character {
			case '(':
				floorNumber++
			case ')':
				floorNumber--
			default:
				fmt.Printf("Character not recognized: %s", string(character))
			}
			if floorNumber == -1 {
				return pos + 1
			}
		}
	}
	return 0
}
