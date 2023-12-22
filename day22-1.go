package main

import (
	"fmt"
	"regexp"
	"slices"
)

type day22Pos struct {
	x, y, z int
}

type day22Brick struct {
	a, b day22Pos
}

// X by Y.
type day22Field [][]int

type day22Supports struct {
	s, by map[int][]int
}

func day22part1(filename string) (string, error) {
	bricks, err := day22ReadPlan(filename)
	if err != nil {
		return "", err
	}
	fall, _ := day22Fall(bricks)
	sup := day22FindSupports(fall)
	destroyed := sup.canDestroy(fall)
	return fmt.Sprint(len(destroyed)), nil
}

func (sup day22Supports) canDestroy(fall []day22Brick) []int {
	var res []int
	for i := range fall {
		destroy := true
		// Going through the bricks this brick supports…
		for _, supportedBrick := range sup.s[i] {
			// If each of those bricks in ONLY supported by this brick, then we can't destroy.
			by := sup.by[supportedBrick]
			if len(by) == 1 && by[0] == i {
				destroy = destroy && false
			}
		}
		if destroy {
			res = append(res, i)
		}
	}
	return res
}

func day22NewSupports(count int) day22Supports {
	res := day22Supports{
		s:  map[int][]int{},
		by: map[int][]int{},
	}
	for i := 0; i < count; i++ {
		res.s[i] = make([]int, 0)
		res.by[i] = make([]int, 0)
	}
	return res
}

func (a day22Brick) crosses(b day22Brick) bool {
	ap := a.points()
	for _, v := range b.points() {
		if slices.Contains(ap, v) {
			return true
		}
	}
	return false
}

func (b day22Brick) points() []day22Pos {
	var bp []day22Pos
	for x := b.a.x; x <= b.b.x; x++ {
		for y := b.a.y; y <= b.b.y; y++ {
			for z := b.a.z; z <= b.b.z; z++ {
				bp = append(bp, day22Pos{x, y, z})
			}
		}
	}
	return bp
}

// Returns index to indices, supports and supported by.
func day22FindSupports(fall []day22Brick) day22Supports {
	sup := day22NewSupports(len(fall))
	for i, br := range fall {
		up := br
		up.moveZ(-1)
		for ii := i + 1; ii < len(fall); ii++ {
			if up.crosses(fall[ii]) {
				sup.s[i] = append(sup.s[i], ii)
				sup.by[ii] = append(sup.by[ii], i)
			}
		}
	}
	return sup
}

// Returns fallen bricks and count.
func day22Fall(b []day22Brick) ([]day22Brick, int) {
	var count int
	field := day22MakeField(b)
	var res []day22Brick
	for _, air := range b {
		ground := air
		z := field.maxZDiff(ground)
		ground.moveZ(z)
		field.setZ(ground)
		if z > 0 {
			count += 1
		}
		res = append(res, ground)
	}
	return res, count
}

func (b *day22Brick) moveZ(by int) {
	b.a.z -= by
	b.b.z -= by
}

func (f day22Field) setZ(b day22Brick) {
	for x := b.a.x; x <= b.b.x; x++ {
		for y := b.a.y; y <= b.b.y; y++ {
			f[x][y] = b.b.z
		}
	}
}

func (f day22Field) maxZDiff(b day22Brick) int {
	var z int
	for x := b.a.x; x <= b.b.x; x++ {
		for y := b.a.y; y <= b.b.y; y++ {
			z = max(z, f[x][y])
		}
	}
	return b.a.z - z - 1
}

func day22MakeField(b []day22Brick) day22Field {
	var x, y int
	for _, v := range b {
		x = max(x, v.b.x)
		y = max(y, v.b.y)
	}
	f := make(day22Field, x+1)
	for i := 0; i <= y; i++ {
		f[i] = make([]int, y+1)
	}
	return f
}

var day22BrickRegex = regexp.MustCompile(`(.*)~(.*)`)

func day22ReadPlan(filename string) ([]day22Brick, error) {
	var res []day22Brick
	if err := forLineError(filename, func(line string) error {
		m := day22BrickRegex.FindStringSubmatch(line)
		a := parseNumberLine(m[1])
		b := parseNumberLine(m[2])
		res = append(res, day22Brick{
			a: day22Pos{a[0], a[1], a[2]},
			b: day22Pos{b[0], b[1], b[2]},
		})
		return nil
	}); err != nil {
		return res, err
	}
	slices.SortFunc(res, func(a, b day22Brick) int {
		return a.a.z - b.a.z
	})
	return res, nil
}
