package main

import (
	"fmt"
)

func day3part2(filename string) (string, error) {
	lines, err := allLines(filename)
	if err != nil {
		return "", err
	}

	maybeGears := make(map[day3Point][]day3Number)
	for _, num := range day3AllNumbers(lines) {
		day3WalkAroundNumber(num, lines, func(y, x int, b byte) bool {
			if b == '*' {
				loc := day3Point{y, x}
				maybeGears[loc] = append(maybeGears[loc], num)
				return false
			}

			return true
		})
	}

	var total int
	for _, v := range maybeGears {
		if len(v) == 2 {
			v0, err0 := v[0].value(lines)
			v1, err1 := v[1].value(lines)
			if err0 != nil || err1 != nil {
				return "", fmt.Errorf("bad numbers: %v, %v", err0, err1)
			}
			total += v0 * v1
		}
	}
	return fmt.Sprint(total), nil
}
