package main

import (
	"fmt"
	"slices"
	"strconv"
)

func day4part2(filename string) (string, error) {
	// Card number to number of cards.
	cards := map[int]int{}
	if err := forLineError(filename, func(line string) error {
		matches := day4CardRe.FindStringSubmatch(line)
		cardNum, err := strconv.Atoi(matches[1])
		if err != nil {
			return err
		}
		cards[cardNum] += 1

		numMatches := day4NumberOfMatches(day4NumberSlice(matches[3]), day4NumberSlice(matches[2]))

		if numMatches == 0 {
			return nil
		}

		for i := cardNum + 1; i <= cardNum+numMatches; i++ {
			cards[i] += cards[cardNum]
		}
		return nil
	}); err != nil {
		return "", err
	}

	var total int
	for _, v := range cards {
		total += v
	}
	return fmt.Sprint(total), nil
}

func day4NumberOfMatches(card []int, winning []int) int {
	num := 0
	for _, n := range card {
		if slices.Contains(winning, n) {
			num += 1
		}
	}
	return num
}
