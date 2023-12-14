package main

import (
	"fmt"
	"slices"
)

type day14Map [][]byte

func day14part1(filename string) (string, error) {
	m, err := day14ReadMap(filename)
	if err != nil {
		return "", err
	}
	var total int
	for x := 0; x < len(m[0]); x++ {
		col := m.column(x)
		day14SlideLeft(col)
		total += day14Weight(col)
	}
	return fmt.Sprint(total), nil
}

func day14Slide(s []byte, cmp func(byte, byte) int) {
	var i, l int
	for i < len(s) {
		if l == i {
			switch s[i] {
			case 'O', '.':
				l = i
			case '#':
				l = i + 1
			}
			i += 1
			continue
		}
		if s[i] == '#' {
			slices.SortFunc(s[l:i], cmp)
			l = i + 1
		} else if i == len(s)-1 {
			slices.SortFunc(s[l:i+1], cmp)
		}
		i += 1
	}
}

func day14SlideLeft(s []byte) {
	day14Slide(s, func(a, b byte) int {
		return int(b) - int(a)
	})
}

func (m day14Map) column(x int) []byte {
	var res []byte
	for _, vy := range m {
		res = append(res, vy[x])
	}
	return res
}

func day14ReadMap(filename string) (day14Map, error) {
	var m day14Map
	if err := forLineError(filename, func(line string) error {
		m = append(m, []byte(line))
		return nil
	}); err != nil {
		return m, err
	}
	return m, nil
}

// Weight as function of reverse distance from end.
func day14Weight(s []byte) int {
	l := len(s)
	var res int
	for i, v := range s {
		if v == 'O' {
			res += l - i
		}
	}
	return res
}
