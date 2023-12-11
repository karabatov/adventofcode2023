package main

import (
	"fmt"
	"slices"
)

func day11part2(filename string) (string, error) {
	u, err := day11ReadMap(filename)
	if err != nil {
		return "", err
	}
	day11ExpandMap(u)
	day11PrintMap(u)
	gal := day11FindGalaxies(u)
	pairs := day11AllPairs(gal)
	var total int
	for _, p := range pairs {
		total += day11PathLen(u, p[0], p[1], func(dp day11Pos) int {
			return u.weight(dp, 1000000)
		})
	}
	return fmt.Sprint(total), nil
}

func day11ExpandMap(u day11Map) {
	for y, vy := range u {
		if !slices.Contains(vy, '#') {
			for x := range vy {
				u[y][x] = '*'
			}
		}
	}

	x := 0
	for {
		if x == len(u[0]) {
			break
		}
		colIsEmpty := true
		for _, v := range u {
			colIsEmpty = colIsEmpty && v[x] != '#'
		}
		if !colIsEmpty {
			x += 1
			continue
		}
		for y := range u {
			u[y][x] = '*'
		}
		x += 1
	}
}

func (u day11Map) weight(p day11Pos, exp int) int {
	if u[p.y][p.x] == '*' {
		return exp
	}

	return 1
}
