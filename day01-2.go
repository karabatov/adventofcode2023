package main

import (
	"fmt"
	"strings"
)

func day1part2(filename string) (string, error) {
	var total int

	if err := forLine(filename, func(line string) {
		firstIdx := len(line)
		lastIdx := -1
		var first, last int

		nums := map[string]int{
			"1":     1,
			"2":     2,
			"3":     3,
			"4":     4,
			"5":     5,
			"6":     6,
			"7":     7,
			"8":     8,
			"9":     9,
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
		}

		for k, v := range nums {
			if f := strings.Index(line, k); f >= 0 && f < firstIdx {
				firstIdx = f
				first = v
			}
			if l := strings.LastIndex(line, k); l >= 0 && l > lastIdx {
				lastIdx = l
				last = v
			}
		}

		total += first*10 + last
	}); err != nil {
		return "", err
	}

	return fmt.Sprint(total), nil
}
