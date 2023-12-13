package main

import (
	"fmt"
	"log"
)

type day13RC struct {
	col, row int
}

type day13Smudge struct {
	y, x int
}

func day13part2(filename string) (string, error) {
	maps, err := day13ReadMaps(filename)
	if err != nil {
		return "", err
	}
	var cols, rows int
	for i, m := range maps {
		day13PrintMap(m)
		s := day13FindSmudge(m)
		cols += s.col
		rows += s.row * 100
		log.Print(i, s)
	}
	return fmt.Sprint(cols + rows), nil
}

func (m day13Map) reflections(skip day13RC) day13RC {
	var rc day13RC
	if rows, ok := m.findRowReflection(); ok {
		for _, r := range rows {
			if r == skip.row {
				continue
			}
			rc.row = r
		}
	}
	if cols, ok := m.findColumnReflection(); ok {
		for _, c := range cols {
			if c == skip.col {
				continue
			}
			rc.col = c
		}
	}
	return rc
}

func (m day13Map) flip(s day13Smudge) {
	switch m[s.y][s.x] {
	case '.':
		m[s.y][s.x] = '#'
	case '#':
		m[s.y][s.x] = '.'
	default:
		log.Fatal("Unknown map value")
	}
}

func day13FindSmudge(m day13Map) day13RC {
	origRefl := m.reflections(day13RC{})
	for y := range m {
		for x := range m[y] {
			s := day13Smudge{y, x}
			m.flip(s)
			if newRefl := m.reflections(origRefl); !newRefl.isZero() && newRefl != origRefl {
				d := newRefl.diff(origRefl)
				return d
			}
			m.flip(s)
		}
	}
	log.Fatal("There must be a smudge")
	return origRefl
}

func (r day13RC) isZero() bool {
	return r.col == 0 && r.row == 0
}

func (r day13RC) diff(o day13RC) day13RC {
	s := r
	if r.col == o.col {
		s.col = 0
	}
	if r.row == o.row {
		s.row = 0
	}
	return s
}
