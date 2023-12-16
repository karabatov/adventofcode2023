package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

// A labelled lens.
type day15Lens struct {
	lbl string
	foc int
}

type day15Conf map[int][]day15Lens

func day15part2(filename string) (string, error) {
	var l string
	if err := forLineError(filename, func(line string) error {
		l = line
		return nil
	}); err != nil {
		return "", err
	}
	s := strings.Split(l, ",")
	c, err := day15RunInit(s)
	if err != nil {
		return "", err
	}
	var total int
	for k, v := range c {
		total += day15FocusPower(k, v)
	}
	return fmt.Sprint(total), nil
}

func day15FocusPower(box int, v []day15Lens) int {
	var res int
	for i, l := range v {
		res += (box + 1) * (i + 1) * l.foc
	}
	return res
}

func day15NewConf() day15Conf {
	c := make(day15Conf)
	for i := 0; i < 256; i++ {
		c[i] = make([]day15Lens, 0)
	}
	return c
}

func day15RunInit(p []string) (day15Conf, error) {
	c := day15NewConf()
	for _, v := range p {
		if m := strings.Index(v, "-"); m > 0 {
			c.removeLens(v[:m])
		} else if e := strings.Index(v, "="); e > 0 {
			f, err := strconv.Atoi(v[e+1:])
			if err != nil {
				return nil, err
			}
			l := day15Lens{
				lbl: v[:e],
				foc: f,
			}
			c.insertLens(l)
		} else {
			log.Fatal("Expected op code")
		}
	}
	return c, nil
}

func (c day15Conf) removeLens(lbl string) {
	box := day15Hash([]byte(lbl))
	if i := slices.IndexFunc(c[box], func(dl day15Lens) bool {
		return dl.lbl == lbl
	}); i >= 0 {
		c[box] = slices.Delete(c[box], i, i+1)
	}
}

func (c day15Conf) insertLens(l day15Lens) {
	box := day15Hash([]byte(l.lbl))
	if i := slices.IndexFunc(c[box], func(dl day15Lens) bool {
		return dl.lbl == l.lbl
	}); i >= 0 {
		c[box][i] = l
	} else {
		c[box] = append(c[box], l)
	}
}
