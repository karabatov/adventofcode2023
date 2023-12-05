package main

import (
	"fmt"
	"regexp"
	"slices"
)

var (
	day4CardRe = regexp.MustCompile(`^Card\s+(\d+): (.*) \| (.*)$`)
)

func day4part1(filename string) (string, error) {
	var total int

	if err := forLine(filename, func(line string) {
		matches := day4CardRe.FindStringSubmatch(line)
		winning := parseNumberLine(matches[2])
		pow := -1
		for _, n := range parseNumberLine(matches[3]) {
			if slices.Contains(winning, n) {
				pow += 1
			}
		}
		if pow >= 0 {
			total += 1 << pow
		}
	}); err != nil {
		return "", err
	}

	return fmt.Sprint(total), nil
}
