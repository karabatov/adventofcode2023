package main

import (
	"fmt"
	"log"
)

func day10part2(filename string) (string, error) {
	var ground day10Ground
	if err := forLineError(filename, func(line string) error {
		ground = append(ground, []day10Pipe(line))
		return nil
	}); err != nil {
		return "", err
	}

	path := make(map[day10Pos]int)

	startPos := ground.findStartPos()
	startDir := ground.findStartDir(startPos)
	startPipe := ground.startPipeType(startPos)
	path[startPos] = day10EvenOddPipeValue(startPipe)
	ground.fullWalk(startPos, startDir, func(i int, dp day10Pos) bool {
		pipe, _ := ground.pipeAt(dp)
		if pipe == 'S' {
			pipe = startPipe
		}
		path[dp] = day10EvenOddPipeValue(pipe)
		return false
	})

	var windings []int
	for y := 0; y < len(ground); y++ {
		for x := 0; x < len(ground[y]); x++ {
			pos := day10Pos{y, x}
			val := ground.evenOdd(pos, path)
			if _, isPath := path[pos]; isPath {
				continue
			}
			windings = append(windings, val)
		}
	}

	var total int
	for _, v := range windings {
		if v%2 != 0 {
			total += 1
		}
	}
	return fmt.Sprint(total), nil
}

func day10EvenOddPipeValue(p day10Pipe) int {
	res := 0
	switch p {
	case '.':
		res = 0
	case '7', 'L', '-', '|':
		res = 1
	case 'F', 'J':
		res = 2
	default:
		log.Fatalf("Bad pipe value %v", p)
	}
	return res
}

func (g day10Ground) evenOdd(from day10Pos, path map[day10Pos]int) int {
	res := 0
	y := from.y - 1
	x := from.x + 1
	for {
		if y < 0 || x == len(g[y]) {
			break
		}
		res += path[day10Pos{y, x}]
		y -= 1
		x += 1
	}
	return res
}

var day10StartPipes = map[day10Dir]map[day10Dir]day10Pipe{
	day10DirN: {
		day10DirS: '|',
		day10DirE: 'L',
		day10DirW: 'J',
	},
	day10DirS: {
		day10DirN: '|',
		day10DirE: 'F',
		day10DirW: '7',
	},
	day10DirE: {
		day10DirN: 'L',
		day10DirS: 'F',
		day10DirW: '-',
	},
	day10DirW: {
		day10DirN: 'J',
		day10DirS: '7',
		day10DirE: '-',
	},
}

func (g day10Ground) startPipeType(start day10Pos) day10Pipe {
	var res []day10Dir
	for _, startDir := range day10Directions {
		if _, canMove := g.move(startDir, start); canMove {
			res = append(res, startDir)
		}
	}
	if len(res) != 2 {
		log.Fatalf("We should only be able to go two directions")
	}
	pipe, ok := day10StartPipes[res[0]][res[1]]
	if !ok {
		log.Fatalf("There should be a pipe variant for ", res)
	}
	return pipe
}
