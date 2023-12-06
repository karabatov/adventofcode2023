package main

import (
	"fmt"
	"regexp"
)

var (
	day6LineRe = regexp.MustCompile(`^.+:\s+(.+)$`)
)

func day6part1(filename string) (string, error) {
	var time, dst []int
	if err := forLineError(filename, func(line string) error {
		matches := day6LineRe.FindStringSubmatch(line)[1]
		if len(time) == 0 {
			time = parseNumberLine(matches)
			return nil
		}
		if len(dst) == 0 {
			dst = parseNumberLine(matches)
			return nil
		}
		return nil
	}); err != nil {
		return "", err
	}

	total := 1

	for i := 0; i < len(time); i++ {
		total *= day6TimesToBeat(time[i], dst[i])
	}

	return fmt.Sprint(total), nil
}

func day6TimesToBeat(time, dst int) int {
	res := 0
	for i := 1; i < dst; i++ {
		if i*(time-i) > dst {
			res += 1
		}
	}
	return res
}
