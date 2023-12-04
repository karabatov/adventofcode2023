package main

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
)

var (
	day4CardRe = regexp.MustCompile(`^Card\s+(\d+): (.*) \| (.*)$`)
	day4NumRe  = regexp.MustCompile(`\d+`)
)

func day4part1(filename string) (string, error) {
	var total int

	if err := forLine(filename, func(line string) {
		matches := day4CardRe.FindStringSubmatch(line)
		winning := day4NumberSlice(matches[2])
		pow := -1
		for _, n := range day4NumberSlice(matches[3]) {
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

func day4NumberSlice(str string) []int {
	nums := make([]int, 0)
	for _, s := range day4NumRe.FindAllString(str, -1) {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("could not convert to number: '%s'", s)
		}
		nums = append(nums, n)
	}
	return nums
}
