package year2015

import (
	"strings"
)

type Day05 struct{}

func isVowel(character rune) bool {
	switch character {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func containsSpecificCases(text string) bool {
	switch text {
	case "ab", "cd", "pq", "xy":
		return true
	}
	return false
}

func processNiceString(line string) (bool, error) {
	var previousCharacter rune

	vowelCounter := 0
	twiceInARow := false

	for index, character := range line {
		// Count vowels
		if isVowel(character) {
			vowelCounter++
		}

		// First character case
		if index == 0 {
			previousCharacter = character
			continue
		}

		// Check if there are two in a row
		if previousCharacter == character {
			twiceInARow = true
		}

		specificCase := string(previousCharacter) + string(character)
		if containsSpecificCases(specificCase) {
			return false, nil
		}

		previousCharacter = character
	}

	if vowelCounter >= 3 && twiceInARow {
		return true, nil
	}

	return false, nil
}

func containsPairs(line string) bool {
	for i := 0; i < len(line)-3; i++ {
		pair := line[i : i+2]
		if strings.Contains(line[i+2:], pair) {
			return true
		}
	}
	return false
}

func containsRepeatedLetter(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
}

func processNicerString(line string) (bool, error) {
	checkPairs := containsPairs(line)
	checkRepeatedLetter := containsRepeatedLetter(line)

	if checkPairs && checkRepeatedLetter {
		return true, nil
	}

	return false, nil
}

func (p Day05) PartA(lines []string) any {
	niceStringsCount := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		isStringNice, err := processNiceString(line)
		if err != nil {
			return err
		}

		if isStringNice {
			niceStringsCount++
		}
	}

	return niceStringsCount
}

func (p Day05) PartB(lines []string) any {
	niceStringsCount := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		isStringNice, err := processNicerString(line)
		if err != nil {
			return err
		}

		if isStringNice {
			niceStringsCount++
		}
	}

	return niceStringsCount
}
