package main

import (
	"fmt"
)

func day17part2(filename string) (string, error) {
	f, err := day17ReadMap(filename)
	if err != nil {
		return "", err
	}
	to := day17Pos{len(f) - 1, len(f[0]) - 1}
	s := f.shortest4(day17Pos{0, 0}, to, 4, 10)
	return fmt.Sprint(s), nil
}
