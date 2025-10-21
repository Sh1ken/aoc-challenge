package year2015

import (
	"strings"
)

type Day08 struct{}

func processString(line string) (processedLine string, err error) {
	var str strings.Builder

	for i := 0; i < len(line); i++ {
		currentCharacter := line[i]

		switch currentCharacter {
		case '"':
			continue

		case '\\':
			if line[i+1] == 'x' {
				str.WriteRune('H')
				i += 3
			} else {
				str.WriteString(string(line[i+1]))
				i += 1
			}
		default:
			str.WriteByte(currentCharacter)
		}
	}

	return str.String(), nil
}

func processEscapedString(line string) (processedLine string, err error) {
	var str strings.Builder

	for i := 0; i < len(line); i++ {
		currentCharacter := line[i]
		switch currentCharacter {
		case '"':
			if i == 0 {
				str.WriteString("\"")
			}

			str.WriteString("\\\"")

			if i == len(line)-1 {
				str.WriteString("\"")
			}
		case '\\':
			str.WriteString("\\\\")
		default:
			str.WriteByte(currentCharacter)
		}
	}

	return str.String(), err
}

func (p Day08) PartA(lines []string) any {
	result := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		processedLine, err := processString(line)
		if err != nil {
			return err
		}

		result += len(line) - len(processedLine)
	}

	return result
}

func (p Day08) PartB(lines []string) any {
	result := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		processedLine, err := processEscapedString(line)
		if err != nil {
			return err
		}

		result += len(processedLine) - len(line)
	}

	return result
}
