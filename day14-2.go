package main

import (
	"fmt"
)

func day14part2(filename string) (string, error) {
	m, err := day14ReadMap(filename)
	if err != nil {
		return "", err
	}
	var cycle int
	var totals []int
	for cycle < 1000 {
		m.cycle()
		t := m.northLoad()
		totals = append(totals, t)
		cycle += 1
	}
	cycleLen, cycleStart := day14Cycle(totals)
	diff := (1000000000 - cycleStart - 1) % cycleLen
	return fmt.Sprint(totals[cycleStart+diff]), nil
}

// Returns length of cycle and start index.
func day14Cycle(s []int) (int, int) {
	t := 1
	h := 2
	for s[t] != s[h] {
		t += 1
		h += 2
	}

	mu := 0
	t = 0
	for s[t] != s[h] {
		t += 1
		h += 1
		mu += 1
	}

	lam := 1
	h = t + 1
	for s[t] != s[h] {
		h += 1
		lam += 1
	}

	return lam, mu
}

func (m day14Map) northLoad() int {
	var total int
	for x := 0; x < len(m[0]); x++ {
		total += m.weightInColumn(x)
	}
	return total
}

func (m day14Map) weightInColumn(x int) int {
	var res int
	for y, vy := range m {
		if vy[x] == 'O' {
			res += len(m) - y
		}
	}
	return res
}

func (m day14Map) cycle() {
	m.north()
	m.west()
	m.south()
	m.east()
}

func day14SlideRight(s []byte) {
	day14Slide(s, func(a, b byte) int {
		return int(a) - int(b)
	})
}

func (m day14Map) north() {
	for x := 0; x < len(m[0]); x++ {
		col := m.column(x)
		day14SlideLeft(col)
		for y, vy := range m {
			vy[x] = col[y]
		}
	}
}

func (m day14Map) west() {
	for _, vy := range m {
		day14SlideLeft(vy)
	}
}

func (m day14Map) south() {
	for x := 0; x < len(m[0]); x++ {
		col := m.column(x)
		day14SlideRight(col)
		for y, vy := range m {
			vy[x] = col[y]
		}
	}
}

func (m day14Map) east() {
	for _, vy := range m {
		day14SlideRight(vy)
	}
}
