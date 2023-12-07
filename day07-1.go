package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type day7Type int

const (
	day7TypeHighCard day7Type = iota + 1
	day7Type1Pair
	day7Type2Pair
	day7Type3OfAKind
	day7TypeFullHouse
	day7Type4OfAKind
	day7Type5OfAKind
)

var day7CardValues = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type day7Hand struct {
	cards string
	class day7Type
}

func (hand day7Hand) valueAt(idx int) int {
	return day7CardValues[hand.cards[idx]]
}

func (hand day7Hand) stronger(other day7Hand) int {
	if res := hand.class - other.class; res != 0 {
		return int(res)
	}

	for i := 0; i < len(hand.cards); i++ {
		if res := hand.valueAt(i) - other.valueAt(i); res != 0 {
			return res
		}
	}

	return -1
}

type day7Move struct {
	hand day7Hand
	bid  int
}

func day7part1(filename string) (string, error) {
	var moves []day7Move
	if err := forLineError(filename, func(line string) error {
		m, err := day7ParseMove(line, day7ParseHand)
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

func day7HandType(cards string) day7Type {
	chars := day7Chars(cards)
	counts := day7CountsFromChars(chars)
	return day7TypeFromCounts(counts)
}

func day7Chars(cards string) map[rune]int {
	chars := make(map[rune]int)
	for _, b := range cards {
		chars[b] += 1
	}
	return chars
}

func day7CountsFromChars(chars map[rune]int) []int {
	counts := make([]int, 0)
	for _, v := range chars {
		counts = append(counts, v)
	}
	slices.Sort(counts)
	slices.Reverse(counts)
	return counts
}

func day7TypeFromCounts(counts []int) day7Type {
	var t day7Type
	if slices.Equal(counts, []int{5}) {
		t = day7Type5OfAKind
	} else if slices.Equal(counts, []int{4, 1}) {
		t = day7Type4OfAKind
	} else if slices.Equal(counts, []int{3, 2}) {
		t = day7TypeFullHouse
	} else if slices.Equal(counts, []int{3, 1, 1}) {
		t = day7Type3OfAKind
	} else if slices.Equal(counts, []int{2, 2, 1}) {
		t = day7Type2Pair
	} else if slices.Equal(counts, []int{2, 1, 1, 1}) {
		t = day7Type1Pair
	} else {
		t = day7TypeHighCard
	}
	return t
}

func day7ParseHand(cards string) day7Hand {
	return day7Hand{
		cards: cards,
		class: day7HandType(cards),
	}
}

func day7ParseMove(line string, parseHand func(string) day7Hand) (day7Move, error) {
	splits := strings.Split(line, " ")
	move := day7Move{
		hand: parseHand(splits[0]),
	}
	bid, err := strconv.Atoi(splits[1])
	if err != nil {
		return move, err
	}
	move.bid = bid
	return move, nil
}
