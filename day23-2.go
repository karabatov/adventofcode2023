package main

import "fmt"

func day23part2(filename string) (string, error) {
	forest, err := day23ReadMap(filename)
	if err != nil {
		return "", err
	}
	for y, vy := range forest {
		for x := range vy {
			switch forest[y][x] {
			case '^', 'v', '<', '>':
				forest[y][x] = '.'
			}
		}
	}
	start := forest.findStartPos()
	end := forest.findEndPos()
	long := forest.longest(start, end)
	return fmt.Sprint(long - 1), nil
}
