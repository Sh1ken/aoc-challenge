package year2015

import (
	"fmt"
	"strconv"
	"strings"
)

type Day06 struct{}

const GridSize = 1000

type (
	Lights      [GridSize][GridSize]int
	Instruction struct {
		action string
		startX int
		startY int
		endX   int
		endY   int
	}
)

func parseAction(line string) (currentInstruction Instruction, err error) {
	offset := 0
	tokens := strings.Split(line, " ")

	switch tokens[0] {
	case "toggle":
		currentInstruction.action = "toggle"
	case "turn":
		currentInstruction.action = tokens[1]
		offset = 1
	default:
		return Instruction{}, fmt.Errorf("instruction %s not supported", line)
	}

	currentStart := strings.Split(tokens[1+offset], ",")
	currentEnd := strings.Split(tokens[3+offset], ",")

	currentInstruction.startX, _ = strconv.Atoi(currentStart[0])
	currentInstruction.startY, _ = strconv.Atoi(currentStart[1])
	currentInstruction.endX, _ = strconv.Atoi(currentEnd[0])
	currentInstruction.endY, _ = strconv.Atoi(currentEnd[1])

	return currentInstruction, nil
}

func changeBrightness(action string, currentValue int) int {
	switch action {
	case "toggle":
		currentValue += 2
	case "on":
		currentValue += 1
	case "off":
		if currentValue >= 1 {
			currentValue -= 1
		}
	}
	return currentValue
}

func updateLightState(action string, currentValue int) int {
	switch action {
	case "toggle":
		if currentValue == 0 {
			currentValue = 1
		} else {
			currentValue = 0
		}
	case "on":
		currentValue = 1
	case "off":
		currentValue = 0
	}
	return currentValue
}

func executeInstruction(lights *Lights, instruction Instruction, isBrightness bool) error {
	for i := instruction.startX; i <= instruction.endX; i++ {
		for j := instruction.startY; j <= instruction.endY; j++ {
			if isBrightness {
				lights[i][j] = changeBrightness(instruction.action, lights[i][j])
			} else {
				lights[i][j] = updateLightState(instruction.action, lights[i][j])
			}
		}
	}
	return nil
}

func countLit(lights *Lights) int {
	count := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			count += lights[i][j]
		}
	}

	return count
}

func (p Day06) PartA(lines []string) any {
	var lights Lights

	for _, line := range lines {
		if line == "" {
			continue
		}

		currentInstruction, err := parseAction(line)
		if err != nil {
			return err
		}

		executeInstruction(&lights, currentInstruction, false)

	}
	return countLit(&lights)
}

func (p Day06) PartB(lines []string) any {
	var lights Lights
	for _, line := range lines {
		if line == "" {
			continue
		}

		currentInstruction, err := parseAction(line)
		if err != nil {
			return err
		}

		executeInstruction(&lights, currentInstruction, true)

	}
	return countLit(&lights)
}
