package main

import (
	"fmt"
	"unicode"
)

func day1part2(filename string) (string, error) {
	var total int

	if err := forLine(filename, func(line string) {
		first := -1
		last := -1
		for _, r := range line {
			if unicode.IsDigit(r) {
				if first < 0 {
					first = int(r - '0')
				}
				last = int(r - '0')
			}
		}
		total += first*10 + last
	}); err != nil {
		return "", err
	}

	return fmt.Sprint(total), nil
}
