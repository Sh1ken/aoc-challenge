package year2015

import "fmt"

type Day03 struct{}

type Position struct {
	X int
	Y int
}

func getNewPosition(currentPosition Position, direction rune) (Position, error) {
	switch direction {
	case '^':
		currentPosition.Y += 1
	case '>':
		currentPosition.X += 1
	case 'v':
		currentPosition.Y -= 1
	case '<':
		currentPosition.X -= 1
	default:
		return Position{}, fmt.Errorf("character %c not recognized", direction)
	}

	return currentPosition, nil
}

func processDirections(directions string, roboSantaEnabled bool) (houses int, err error) {
	// Track visited houses using a map as a set
	visitedHouses := make(map[Position]bool)
	currentSantaHouse := Position{0, 0}
	currentRobotHouse := Position{0, 0}

	// First house always gets a present
	visitedHouses[currentSantaHouse] = true

	for i, direction := range directions {
		// Check if it is robot's turn
		isRobotTurn := roboSantaEnabled && i%2 != 0

		// Move in a direction and visit house
		if isRobotTurn {
			currentRobotHouse, err = getNewPosition(currentRobotHouse, direction)
			if err != nil {
				return 0, err
			}
			visitedHouses[currentRobotHouse] = true
		} else {
			currentSantaHouse, err = getNewPosition(currentSantaHouse, direction)
			if err != nil {
				return 0, err
			}
			visitedHouses[currentSantaHouse] = true
		}
	}

	return len(visitedHouses), nil
}

func (p Day03) PartA(lines []string) any {
	if len(lines) == 0 {
		return 0
	}

	totalHouses, err := processDirections(lines[0], false)
	if err != nil {
		return err
	}

	return totalHouses
}

func (p Day03) PartB(lines []string) any {
	if len(lines) == 0 {
		return 0
	}

	totalHouses, err := processDirections(lines[0], true)
	if err != nil {
		return err
	}

	return totalHouses
}
