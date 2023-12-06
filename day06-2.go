package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day6part2(filename string) (string, error) {
	var time, dst int
	if err := forLineError(filename, func(line string) error {
		matches := day6LineRe.FindStringSubmatch(line)[1]
		numStr := strings.ReplaceAll(matches, " ", "")
		parsed, err := strconv.Atoi(numStr)
		if err != nil {
			return err
		}
		if time == 0 {
			time = parsed
			return nil
		}
		if dst == 0 {
			dst = parsed
			return nil
		}
		return nil
	}); err != nil {
		return "", err
	}

	total := day6TimesToBeatQuick(time, dst)
	return fmt.Sprint(total), nil
}

func day6TimesToBeatQuick(time, dst int) int {
	res := 0

	for i := 1; i < time/2; i++ {
		if d := i * (time - i); d > dst {
			return time - i*2 + 1
		}
	}

	return res
}
