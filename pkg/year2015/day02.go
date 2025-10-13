package year2015

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day02 struct{}

func parseMeasurements(line string) ([]int, error) {
	var measurements []int

	// parse input
	parsedLine := strings.Split(line, "x")
	if len(parsedLine) != 3 {
		return nil, fmt.Errorf("expected 3 measurements, got %d", len(parsedLine))
	}

	// cast to int
	for _, parsedMeasurement := range parsedLine {
		val, err := strconv.Atoi(parsedMeasurement)
		if err != nil {
			return nil, err
		}
		measurements = append(measurements, val)
	}

	// sort input
	slices.Sort(measurements)

	return measurements, nil
}

func (p Day02) PartA(lines []string) any {
	totalSum := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		// obtain, parse and sort input
		measurements, err := parseMeasurements(line)
		if err != nil {
			return fmt.Errorf("failed to parse input: %w ", err)
		}

		// split in sorted dimensions
		l, w, h := measurements[0], measurements[1], measurements[2]

		// get the paper
		totalSum += 2*l*w + 2*w*h + 2*h*l

		// also some extra paper
		totalSum += l * w
	}

	return totalSum
}

func (p Day02) PartB(lines []string) any {
	totalSum := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		// obtain, parse and sort input
		measurements, err := parseMeasurements(line)
		if err != nil {
			return fmt.Errorf("failed to parse input: %w", err)
		}

		// split in sorted dimensions
		l, w, h := measurements[0], measurements[1], measurements[2]

		// get the perimeter of the smallest face
		totalSum += 2*l + 2*w

		// also the extra for the ribbon
		totalSum += l * w * h
	}

	return totalSum
}
