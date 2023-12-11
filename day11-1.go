package main

import (
	"fmt"
	"log"
	"slices"
)

type day11Map [][]byte

type day11Pos struct {
	y, x int
}

func day11part1(filename string) (string, error) {
	u, err := day11ReadMap(filename)
	if err != nil {
		return "", err
	}
	u = day11ExpandMapVertical(u)
	day11ExpandMapHorizontal(u)
	for _, vy := range u {
		for _, vx := range vy {
			fmt.Print(string(vx))
		}
		fmt.Println()
	}
	gal := day11FindGalaxies(u)
	pairs := day11AllPairs(gal)
	log.Print(len(pairs))
	var total int
	for _, p := range pairs {
		total += day11PathLen(u, p[0], p[1])
	}
	return fmt.Sprint(total), nil
}

func day11ReadMap(filename string) (day11Map, error) {
	var u day11Map
	if err := forLineError(filename, func(line string) error {
		u = append(u, []byte(line))
		return nil
	}); err != nil {
		return u, err
	}
	return u, nil
}

func day11ExpandMapVertical(u day11Map) day11Map {
	var res day11Map
	for _, r := range u {
		res = append(res, r)
		if !slices.Contains(r, '#') {
			res = append(res, r)
		}
	}
	return res
}

func day11ExpandMapHorizontal(u day11Map) {
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
		for y, v := range u {
			u[y] = slices.Insert(v, x, '.')
		}
		x += 2
	}
}

func day11FindGalaxies(u day11Map) []day11Pos {
	var res []day11Pos
	for y, vy := range u {
		for x, vx := range vy {
			if vx == '#' {
				res = append(res, day11Pos{y, x})
			}
		}
	}
	return res
}

func day11AllPairs(g []day11Pos) [][]day11Pos {
	var res [][]day11Pos
	for i, v := range g {
		for j := i + 1; j < len(g); j++ {
			var p []day11Pos
			p = append(p, v)
			p = append(p, g[j])
			res = append(res, p)
		}
	}
	return res
}

func (u day11Map) isValidPos(p day11Pos) bool {
	return p.y >= 0 && p.y < len(u) && p.x >= 0 && p.x < len(u[p.y])
}

type day11Node struct {
	p day11Pos
	l int
}

func day11PathLen(u day11Map, from, to day11Pos) int {
	rowMove := []int{-1, 0, 0, 1}
	colMove := []int{0, -1, 1, 0}
	visited := make(map[day11Pos]bool)
	visited[from] = true

	q := make([]day11Node, 0)
	q = append(q, day11Node{p: from, l: 0})

	for {
		if len(q) == 0 {
			break
		}
		currNode := q[0]
		pt := currNode.p
		if pt.y == to.y && pt.x == to.x {
			return currNode.l
		}
		q = q[1:]

		for i := 0; i < 4; i++ {
			p := day11Pos{pt.y + colMove[i], pt.x + rowMove[i]}
			if u.isValidPos(p) && !visited[p] {
				visited[p] = true
				q = append(q, day11Node{p, currNode.l + 1})
			}
		}
	}

	return -1
}
