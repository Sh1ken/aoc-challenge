package year2015

import "fmt"

type Day03 struct{}

type Pair struct {
	X int
	Y int
}

func processDirections(directions string) (houses int, err error) {
	// We define a 'set'
	m := make(map[Pair]bool)
	currentHouse := Pair{0, 0}

	// First house always gets a present
	m[currentHouse] = true

	for _, direction := range directions {
		// move in a direction
		switch direction {
		case '^':
			currentHouse.Y += 1
		case '>':
			currentHouse.X += 1
		case 'v':
			currentHouse.Y -= 1
		case '<':
			currentHouse.X -= 1
		default:
			return 0, fmt.Errorf("character %c not recognized", direction)
		}

		// and submit to set
		m[currentHouse] = true
	}

	return len(m), nil
}

func (p Day03) PartA(lines []string) any {
	if len(lines) == 0 {
		return 0
	}

	totalHouses, err := processDirections(lines[0])
	if err != nil {
		return err
	}

	return totalHouses
}

func (p Day03) PartB(lines []string) any {
	return "implement_me"
}
