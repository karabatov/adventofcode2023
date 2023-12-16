package main

import (
	"fmt"
	"slices"
)

type day16Field [][]byte

type day16Pos struct {
	y, x int
}

var (
	day16North day16Pos = day16Pos{-1, 0}
	day16South day16Pos = day16Pos{1, 0}
	day16East  day16Pos = day16Pos{0, 1}
	day16West  day16Pos = day16Pos{0, -1}
)

type day16Step struct {
	pos, dir day16Pos
}

type day16Queue []day16Step

type day16Ener map[day16Pos]bool

func day16part1(filename string) (string, error) {
	f, err := day16ReadMap(filename)
	if err != nil {
		return "", err
	}
	e := f.run()
	return fmt.Sprint(len(e)), nil
}

func (f day16Field) run() day16Ener {
	e := make(day16Ener)
	eLen := len(e)
	eSame := 0
	q := day16Queue{day16Step{day16Pos{0, 0}, day16East}}
	for len(q) > 0 && eSame < 10000000 {
		next := q[0]
		q = q[1:]
		for _, s := range f.step(next) {
			q = append(q, s)
		}
		e[next.pos] = true
		if eLen == len(e) {
			eSame += 1
		} else {
			eLen = len(e)
			eSame = 0
		}
	}
	return e
}

func (f day16Field) isValid(p day16Pos) bool {
	return p.y >= 0 && p.y < len(f) && p.x >= 0 && p.x < len(f[p.y])
}

func (f day16Field) step(s day16Step) []day16Step {
	var res []day16Step
	switch f[s.pos.y][s.pos.x] {
	case '.':
		return f.stepsEmpty(s)
	case '/':
		return f.stepsMirrorNE(s)
	case '\\':
		return f.stepsMirrorSW(s)
	case '|':
		return f.stepsSplitterNS(s)
	case '-':
		return f.stepsSplitterEW(s)
	}
	return res
}

func (p day16Pos) move(d day16Pos) day16Pos {
	return day16Pos{p.y + d.y, p.x + d.x}
}

func (f day16Field) stepsSplitterEW(s day16Step) []day16Step {
	var res []day16Step
	switch s.dir {
	case day16North, day16South:
		res = []day16Step{
			{s.pos.move(day16East), day16East},
			{s.pos.move(day16West), day16West},
		}
	case day16East, day16West:
		return f.stepsEmpty(s)
	}
	return slices.DeleteFunc(res, func(ds day16Step) bool {
		return !f.isValid(ds.pos)
	})
}

func (f day16Field) stepsSplitterNS(s day16Step) []day16Step {
	var res []day16Step
	switch s.dir {
	case day16North, day16South:
		return f.stepsEmpty(s)
	case day16East, day16West:
		res = []day16Step{
			{s.pos.move(day16North), day16North},
			{s.pos.move(day16South), day16South},
		}
	}
	return slices.DeleteFunc(res, func(ds day16Step) bool {
		return !f.isValid(ds.pos)
	})
}

func (f day16Field) stepsMirrorSW(s day16Step) []day16Step {
	next := s.pos
	dir := s.dir
	switch s.dir {
	case day16North:
		dir = day16West
	case day16South:
		dir = day16East
	case day16East:
		dir = day16South
	case day16West:
		dir = day16North
	}
	next = next.move(dir)
	if f.isValid(next) {
		return []day16Step{{next, dir}}
	}
	return []day16Step{}
}

func (f day16Field) stepsMirrorNE(s day16Step) []day16Step {
	next := s.pos
	dir := s.dir
	switch s.dir {
	case day16North:
		dir = day16East
	case day16South:
		dir = day16West
	case day16East:
		dir = day16North
	case day16West:
		dir = day16South
	}
	next.y += dir.y
	next.x += dir.x
	if f.isValid(next) {
		return []day16Step{{next, dir}}
	}
	return []day16Step{}
}

func (f day16Field) stepsEmpty(s day16Step) []day16Step {
	next := day16Pos{
		y: s.pos.y + s.dir.y,
		x: s.pos.x + s.dir.x,
	}
	if f.isValid(next) {
		return []day16Step{{next, s.dir}}
	}

	return []day16Step{}
}

func day16ReadMap(filename string) (day16Field, error) {
	var m day16Field
	if err := forLineError(filename, func(line string) error {
		m = append(m, []byte(line))
		return nil
	}); err != nil {
		return m, err
	}
	return m, nil
}
