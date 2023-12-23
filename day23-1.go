package main

import (
	"fmt"
)

type day23Forest [][]byte

type day23Pos struct {
	y, x int
}

func day23part1(filename string) (string, error) {
	forest, err := day23ReadMap(filename)
	if err != nil {
		return "", err
	}
	start := forest.findStartPos()
	end := forest.findEndPos()
	long := forest.longest(start, end)
	return fmt.Sprint(long - 1), nil
}

var day23Len int = -1

func (f day23Forest) longest(from, to day23Pos) int {
	seen := make(map[day23Pos]int)
	out := make(map[day23Pos]int)
	f.walk(seen, out, 0, to, from)
	return out[to]
}

// N, S, E, W
// 0, 1, 2, 3

var day23DirY = []int{-1, 1, 0, 0}
var day23DirX = []int{0, 0, 1, -1}

var day23DirMap = map[byte]int{
	'^': 0,
	'v': 1,
	'>': 2,
	'<': 3,
}

func (f day23Forest) walk(seen, out map[day23Pos]int, steps int, to, point day23Pos) {
	if !f.isValid(point) {
		return
	} else if point == to {
		if seen[point] == 0 && day23Len < steps {
			day23Len = steps
			for dp := range out {
				delete(out, dp)
			}
			for k, v := range seen {
				out[k] = v
			}
			out[point] = steps + 1
		}
	}
	if seen[point] > 0 {
		return
	}
	var next []int
	switch f[point.y][point.x] {
	case '#':
		return
	case '.':
		next = []int{0, 1, 2, 3}
	default:
		next = []int{day23DirMap[f[point.y][point.x]]}
	}
	seen[point] += 1
	for _, v := range next {
		nextPoint := point
		nextPoint.y += day23DirY[v]
		nextPoint.x += day23DirX[v]
		f.walk(seen, out, steps+1, to, nextPoint)
	}
	seen[point] = 0
}

func (f day23Forest) isValid(p day23Pos) bool {
	return p.y >= 0 && p.y < len(f) && p.x >= 0 && p.x < len(f[p.y])
}

func (f day23Forest) findStartPos() day23Pos {
	return day23Pos{0, 1}
}

func (f day23Forest) findEndPos() day23Pos {
	return day23Pos{len(f) - 1, len(f[0]) - 2}
}

func day23ReadMap(filename string) (day23Forest, error) {
	var res day23Forest
	if err := forLine(filename, func(s string) {
		res = append(res, []byte(s))
	}); err != nil {
		return res, err
	}
	return res, nil
}
