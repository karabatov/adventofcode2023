package main

import (
	"fmt"
	"log"
)

func day16part2(filename string) (string, error) {
	f, err := day16ReadMap(filename)
	if err != nil {
		return "", err
	}
	var e int
	starts := f.startPositions()
	for i, v := range starts {
		log.Printf("Run %d of %d...", i+1, len(starts))
		e = max(e, f.runFrom(v))
	}
	return fmt.Sprint(e), nil
}

func (f day16Field) startPositions() []day16Step {
	var starts []day16Step
	// Horizontal beams, inward from both sides.
	for y := 0; y < len(f); y++ {
		starts = append(starts, day16Step{
			day16Pos{y, 0},
			day16East,
		})
		starts = append(starts, day16Step{
			day16Pos{y, len(f[y]) - 1},
			day16West,
		})
	}
	// Vertical beams, inward from both sides.
	for x := 0; x < len(f[0]); x++ {
		starts = append(starts, day16Step{
			day16Pos{0, x},
			day16South,
		})
		starts = append(starts, day16Step{
			day16Pos{len(f) - 1, x},
			day16North,
		})
	}
	return starts
}

func (f day16Field) runFrom(start day16Step) int {
	e := make(day16Ener)
	eLen := len(e)
	eSame := 0
	q := day16Queue{start}
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
	return len(e)
}
