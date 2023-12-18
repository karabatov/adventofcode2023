package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var day18Regex = regexp.MustCompile(`(\w) (\d+) \(#([[:alnum:]]+)\)`)

type day18Dig struct {
	dir string
	num int
	col string
}

type day18Pos struct {
	x, y int
}

type day18GridNode struct {
	og       day18Pos
	dir, col string
}

func day18part1(filename string) (string, error) {
	f, err := day18ReadPlan(filename)
	if err != nil {
		return "", err
	}
	tr := day18Draw(f)
	pp := day18Points(tr)
	area := day18Area(pp)
	return fmt.Sprint(area), nil
}

func day18Points(p []day18GridNode) []day18Pos {
	var res []day18Pos
	for i, v := range p {
		ii := i - 1
		if ii < 0 {
			ii = len(p) - 1
		}
		prev := p[ii]
		n := v.og
		switch prev.dir + v.dir {
		case "DR", "RD":
			n.x += 1
		case "DL", "LD":
			n.x += 1
			n.y += 1
		case "UL", "LU":
			n.y += 1
		}
		res = append(res, n)
	}
	return res
}

func day18Area(p []day18Pos) int {
	var area int
	j := len(p) - 1
	for i := 0; i < len(p); i++ {
		p1 := p[j]
		p2 := p[i]
		area += p1.x*p2.y - p1.y*p2.x
		j = i
	}
	area /= 2
	if area < 0 {
		area *= -1
	}
	return area
}

func day18Draw(p []day18Dig) []day18GridNode {
	var res []day18GridNode
	var x, y int
	for _, v := range p {
		n := day18GridNode{
			og:  day18Pos{x, y},
			col: v.col,
			dir: v.dir,
		}
		switch v.dir {
		case "U":
			y -= v.num
		case "D":
			y += v.num
		case "L":
			x -= v.num
		case "R":
			x += v.num
		}
		res = append(res, n)
	}
	return res
}

func day18ReadPlan(filename string) ([]day18Dig, error) {
	var res []day18Dig
	if err := forLineError(filename, func(line string) error {
		m := day18Regex.FindStringSubmatch(line)
		num, err := strconv.Atoi(m[2])
		if err != nil {
			return err
		}
		res = append(res, day18Dig{
			dir: m[1],
			num: num,
			col: m[3],
		})
		return nil
	}); err != nil {
		return res, err
	}
	return res, nil
}
