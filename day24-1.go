package main

import (
	"fmt"
	"strings"
)

type day24Vec3 struct {
	x, y, z float64
}

type day24Hail struct {
	p, dir day24Vec3
}

type day24Segment struct {
	p1, p2 day24Vec3
}

func day24part1(filename string) (string, error) {
	hails, err := day24ReadList(filename)
	if err != nil {
		return "", err
	}
	minP := day24Vec3{200000000000000, 200000000000000, 0}
	maxP := day24Vec3{400000000000000, 400000000000000, 0}
	var longs []day24Hail
	for _, v := range hails {
		s := v.extendToArea(minP, maxP)
		longs = append(longs, s)
	}
	var total int
	for i := 0; i < len(longs); i++ {
		for ii := i + 1; ii < len(longs); ii++ {
			if xPoint, ok := longs[i].intersect2D(longs[ii]); ok {
				if xPoint.inArea2D(minP, maxP) {
					total += 1
				}
			}
		}
	}
	return fmt.Sprint(total), nil
}

func (p day24Vec3) inArea2D(minP, maxP day24Vec3) bool {
	return p.x >= minP.x && p.x <= maxP.x && p.y >= minP.y && p.y <= maxP.y
}

func (h day24Hail) extendToArea(minP, maxP day24Vec3) day24Hail {
	kX := max((minP.x-h.p.x)/h.dir.x, (maxP.x-h.p.x)/h.dir.x)
	kY := max((minP.y-h.p.y)/h.dir.y, (maxP.y-h.p.y)/h.dir.y)
	kk := max(kX, kY)
	return day24Hail{
		p:   h.p,
		dir: h.dir.justMul(kk),
	}
}

/*
func (v day24Vec3) add(o day24Vec3) day24Vec3 {
	return day24Vec3{
		x: v.x + o.x,
		y: v.y + o.y,
		z: v.z + o.z,
	}
}
*/

func (v day24Vec3) justMul(by float64) day24Vec3 {
	return day24Vec3{
		x: v.x * by,
		y: v.y * by,
		z: v.z * by,
	}
}

// Returns point and if there is an intersection.
func (a day24Hail) intersect2D(b day24Hail) (day24Vec3, bool) {
	s := (-a.dir.y*(a.p.x-b.p.x) + a.dir.x*(a.p.y-b.p.y)) / (-b.dir.x*a.dir.y + a.dir.x*b.dir.y)
	t := (b.dir.x*(a.p.y-b.p.y) - b.dir.y*(a.p.x-b.p.x)) / (-b.dir.x*a.dir.y + a.dir.x*b.dir.y)

	var res day24Vec3
	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
		res.x = a.p.x + (t * a.dir.x)
		res.y = a.p.y + (t * a.dir.y)
		return res, true
	}
	return res, false
}

func day24ReadList(filename string) ([]day24Hail, error) {
	var res []day24Hail
	if err := forLineError(filename, func(line string) error {
		var h day24Hail
		spl := strings.Split(line, " @ ")
		point := parseNumberLine(spl[0])
		h.p = day24NewVec3(point)
		dir := parseNumberLine(spl[1])
		h.dir = day24NewVec3(dir)
		res = append(res, h)
		return nil
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func day24NewVec3(vv []int) day24Vec3 {
	var v day24Vec3
	v.x = float64(vv[0])
	v.y = float64(vv[1])
	v.z = float64(vv[2])
	return v
}
