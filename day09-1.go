package main

import (
	"fmt"
)

func day9part1(filename string) (string, error) {
	total := 0
	if err := forLineError(filename, func(line string) error {
		seq := parseNumberLine(line)
		next := day9NextValue(seq)
		total += next
		return nil
	}); err != nil {
		return "", err
	}
	return fmt.Sprint(total), nil
}

func day9NextValue(seq []int) int {
	allZeros := true
	var subSeq []int
	for i := 1; i < len(seq); i++ {
		num := seq[i] - seq[i-1]
		allZeros = allZeros && num == 0
		subSeq = append(subSeq, num)
	}
	if allZeros {
		return seq[len(seq)-1]
	}
	return seq[len(seq)-1] + day9NextValue(subSeq)
}
