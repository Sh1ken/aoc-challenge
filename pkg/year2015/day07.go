package year2015

import (
	"fmt"
	"strconv"
	"strings"
)

const WireA uint16 = 3176

type Day07 struct{}

type WireNode struct {
	value      uint16
	operation  string
	leftValue  string
	rightValue string
}
type Wires map[string]WireNode

func convertToUint16(value string) (parsedValue uint16, err error) {
	currentValue, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("%s is not a number", value)
	}

	return uint16(currentValue), nil
}

func parseLine(line string) (string, string, error) {
	currentValues := strings.Split(line, " -> ")
	if len(currentValues) != 2 {
		return "", "", fmt.Errorf("error while parsing line %s", line)
	}
	return currentValues[1], currentValues[0], nil
}

func parseInstruction(wires Wires, identifier, instruction string) (err error) {
	instructionArray := strings.Split(instruction, " ")

	switch len(instructionArray) {
	// case 1 is just the value
	case 1:
		currentValue, err := convertToUint16(instructionArray[0])
		if err == nil {
			wires[identifier] = WireNode{currentValue, "VALUE", "", ""}
		} else {
			wires[identifier] = WireNode{0, "PASS", instructionArray[0], ""}
		}

	// case 2 is just for the NOT clause
	case 2:
		wires[identifier] = WireNode{0, "NOT", instructionArray[1], ""}
	// case 3 is for more complex clauses
	case 3:
		wires[identifier] = WireNode{0, instructionArray[1], instructionArray[0], instructionArray[2]}
	default:
		return fmt.Errorf("error while parsing instruction %s", instruction)
	}

	return nil
}

func processNode(wires Wires, wireID string) (value uint16, err error) {
	currentNode := wires[wireID]

	if currentNode.operation == "VALUE" {
		return currentNode.value, nil
	}

	var leftValue, rightValue uint16

	if currentNode.leftValue != "" {
		leftValue, err = convertToUint16(currentNode.leftValue)
		if err != nil {
			// if it's not a number, go grab the value
			leftValue, err = processNode(wires, currentNode.leftValue)
			if err != nil {
				return 0, fmt.Errorf("error processing uint16 value for %s", currentNode.leftValue)
			}
		}
	}

	if currentNode.rightValue != "" {
		rightValue, err = convertToUint16(currentNode.rightValue)
		if err != nil {
			// if it's not a number, go grab the value
			rightValue, err = processNode(wires, currentNode.rightValue)
			if err != nil {
				return 0, fmt.Errorf("error processing uint16 value for %s", currentNode.rightValue)
			}
		}
	}

	switch currentNode.operation {
	case "PASS":
		value = leftValue
	case "NOT":
		value = ^leftValue
	case "OR":
		value = leftValue | rightValue
	case "AND":
		value = leftValue & rightValue
	case "LSHIFT":
		value = leftValue << rightValue
	case "RSHIFT":
		value = leftValue >> rightValue
	default:
		return 0, fmt.Errorf("operation %s not supported", currentNode.operation)
	}

	wires[wireID] = WireNode{value, "VALUE", wires[wireID].leftValue, wires[wireID].rightValue}
	return value, nil
}

func parseWires(lines []string) (Wires, error) {
	wires := make(Wires)

	for _, line := range lines {
		if line == "" {
			continue
		}

		identifier, instruction, err := parseLine(line)
		if err != nil {
			return wires, err
		}

		err = parseInstruction(wires, identifier, instruction)
		if err != nil {
			return wires, err
		}
	}

	return wires, nil
}

func (p Day07) PartA(lines []string) any {
	wires, err := parseWires(lines)
	if err != nil {
		return err
	}

	if _, ok := wires["a"]; ok {
		result, err := processNode(wires, "a")
		if err != nil {
			return err
		}
		return result
	}

	return fmt.Errorf("no element mapped to wire a")
}

func (p Day07) PartB(lines []string) any {
	wires, err := parseWires(lines)
	if err != nil {
		return err
	}

	if _, ok := wires["a"]; ok {
		wires["b"] = WireNode{WireA, "VALUE", "", ""}

		result, err := processNode(wires, "a")
		if err != nil {
			return err
		}
		return result
	}

	return fmt.Errorf("no element mapped to wire a")
}
