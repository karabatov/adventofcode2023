package main

import (
	"fmt"
	"slices"
)

func day7part2(filename string) (string, error) {
	day7CardValues['J'] = 1

	var moves []day7Move
	if err := forLineError(filename, func(line string) error {
		m, err := day7ParseMove(line, func(s string) day7Hand {
			return day7Hand{
				cards: s,
				class: day7HandTypeJoker(s),
			}
		})
		if err != nil {
			return err
		}
		moves = append(moves, m)
		return nil
	}); err != nil {
		return "", err
	}

	slices.SortFunc(moves, func(l, r day7Move) int {
		return l.hand.stronger(r.hand)
	})

	var total int
	for i, m := range moves {
		total += (i + 1) * m.bid
	}
	return fmt.Sprint(total), nil
}

func day7HandTypeJoker(cards string) day7Type {
	chars := day7Chars(cards)
	// Make the Joker into the card with the biggest count.
	if j := chars['J']; j > 0 {
		delete(chars, 'J')
		maxCard := 'J'
		max := 0
		for k, v := range chars {
			if v > max {
				maxCard = k
				max = v
			}
		}
		chars[maxCard] += j
	}
	counts := day7CountsFromChars(chars)
	return day7TypeFromCounts(counts)
}
