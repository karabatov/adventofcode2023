package main

import (
	"fmt"
	"log"
)

type day10Pipe byte

type day10Ground [][]day10Pipe

// 0-based indices to bytes in the ground.
type day10Pos struct {
	y, x int
}

// Direction of travel.
type day10Dir day10Pos

var (
	day10DirEnd = day10Dir{0, 0}
	day10DirN   = day10Dir{-1, 0}
	day10DirS   = day10Dir{1, 0}
	day10DirW   = day10Dir{0, -1}
	day10DirE   = day10Dir{0, 1}
)

var day10Directions = []day10Dir{day10DirN, day10DirS, day10DirE, day10DirW}

func (p day10Pos) move(d day10Dir) day10Pos {
	return day10Pos{p.y + d.y, p.x + d.x}
}

func day10part1(filename string) (string, error) {
	var ground day10Ground
	if err := forLineError(filename, func(line string) error {
		ground = append(ground, []day10Pipe(line))
		return nil
	}); err != nil {
		return "", err
	}

	startPos := ground.findStartPos()
	startDir := ground.findStartDir(startPos)
	steps := ground.fullWalk(startPos, startDir, func(i int, dp day10Pos) bool {
		return false
	})

	return fmt.Sprint(steps / 2), nil
}

func (g day10Ground) fullWalk(startPos day10Pos, startDir day10Dir, stop func(int, day10Pos) bool) int {
	currDir := startDir
	currPos := startPos
	steps := 0
	for {
		nextDir, ok := g.move(currDir, currPos)
		if !ok {
			log.Fatalf("We broke our leg on step %d", steps)
		}
		steps += 1
		if nextDir == day10DirEnd {
			break
		}
		currPos = currPos.move(currDir)
		currDir = nextDir
		if stop(steps, currPos) {
			break
		}
	}
	return steps
}

func (g day10Ground) findStartDir(start day10Pos) day10Dir {
	var res day10Dir
	for _, startDir := range day10Directions {
		if _, canMove := g.move(startDir, start); canMove {
			res = startDir
			break
		}
	}
	return res
}

func (g day10Ground) findStartPos() day10Pos {
	var p day10Pos
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == 'S' {
				p.y = i
				p.x = j
			}
		}
	}
	return p
}

func (g day10Ground) isValidPos(p day10Pos) bool {
	return p.y >= 0 && p.y < len(g) && p.x >= 0 && p.x < len(g[p.y])
}

func (g day10Ground) pipeAt(p day10Pos) (day10Pipe, bool) {
	if g.isValidPos(p) {
		return g[p.y][p.x], true
	}

	return '.', false
}

// Returns true if can go that direction, and the next direction.
// Returns day10DirEnd and true if we reached the start position.
func (g day10Ground) move(dir day10Dir, from day10Pos) (day10Dir, bool) {
	nextPipe, ok := g.pipeAt(from.move(dir))
	if !ok {
		return day10DirEnd, false
	}

	return day10CanMove(dir, nextPipe)
}

var day10DirMapping = map[day10Dir]map[day10Pipe]day10Dir{
	day10DirN: {
		'|': day10DirN,
		'7': day10DirW,
		'F': day10DirE,
	},
	day10DirS: {
		'|': day10DirS,
		'L': day10DirE,
		'J': day10DirW,
	},
	day10DirE: {
		'-': day10DirE,
		'J': day10DirN,
		'7': day10DirS,
	},
	day10DirW: {
		'-': day10DirW,
		'L': day10DirN,
		'F': day10DirS,
	},
}

func day10CanMove(dir day10Dir, next day10Pipe) (day10Dir, bool) {
	if next == 'S' {
		return day10DirEnd, true
	}
	nextDir, ok := day10DirMapping[dir][next]
	return nextDir, ok
}
