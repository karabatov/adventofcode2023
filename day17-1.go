package main

import (
	"fmt"
	"slices"
	"strconv"
)

type day17Field [][]int

type day17Pos struct {
	y, x int
}

type day17Move struct {
	pos, dir day17Pos
}

type day17Acc struct {
	move day17Move
	loss int
}

var (
	day17North day17Pos = day17Pos{-1, 0}
	day17South day17Pos = day17Pos{1, 0}
	day17East  day17Pos = day17Pos{0, 1}
	day17West  day17Pos = day17Pos{0, -1}
)

var day17AllDir = []day17Pos{day17North, day17South, day17East, day17West}

func day17part1(filename string) (string, error) {
	f, err := day17ReadMap(filename)
	if err != nil {
		return "", err
	}
	to := day17Pos{len(f) - 1, len(f[0]) - 1}
	s := f.shortest4(day17Pos{0, 0}, to, 1, 3)
	return fmt.Sprint(s), nil
}

func (f day17Field) shortest4(from, to day17Pos, minRoll, maxRoll int) int {
	dist := make(map[day17Move]int)
	vert := []day17Acc{{}}

	for len(vert) > 0 {
		slices.SortFunc(vert, func(a, b day17Acc) int {
			return a.loss - b.loss
		})

		v := vert[0]
		vert = vert[1:]

		if v.move.pos == to {
			return v.loss
		}

		if d, ok := dist[v.move]; ok && v.loss > d {
			continue
		}

		for _, m := range f.nextMoves(v, minRoll, maxRoll) {
			prev, ok := dist[m.move]
			if !ok || m.loss < prev {
				dist[m.move] = m.loss
				vert = append(vert, m)
			}
		}
	}

	return 0
}

func (f day17Field) at(p day17Pos) int {
	return f[p.y][p.x]
}

func (f day17Field) nextMoves(s day17Acc, minMoves, maxMoves int) []day17Acc {
	var n []day17Acc
	for _, d := range day17AllDir {
		if d.sameOrReverse(s.move.dir) {
			continue
		}
		nextLoss := s.loss
		for i := 1; i <= maxMoves; i++ {
			nextPos := day17Pos{s.move.pos.y + i*d.y, s.move.pos.x + i*d.x}
			if !f.isValid(nextPos) {
				continue
			}
			nextLoss += f.at(nextPos)
			if i < minMoves {
				continue
			}
			nextMove := day17Move{
				pos: nextPos,
				dir: d,
			}
			n = append(n, day17Acc{nextMove, nextLoss})
		}
	}
	return n
}

func (f day17Field) isValid(p day17Pos) bool {
	return p.y >= 0 && p.y < len(f) && p.x >= 0 && p.x < len(f[p.y])
}

func (d day17Pos) sameOrReverse(o day17Pos) bool {
	return d == o || (d.y == -o.y && d.x == -o.x)
}

func day17ReadMap(filename string) (day17Field, error) {
	var m day17Field
	if err := forLineError(filename, func(line string) error {
		var r []int
		for _, v := range line {
			n, err := strconv.Atoi(string(v))
			if err != nil {
				return err
			}
			r = append(r, n)
		}
		m = append(m, r)
		return nil
	}); err != nil {
		return m, err
	}
	return m, nil
}
