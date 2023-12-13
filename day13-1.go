package main

import (
	"fmt"
	"log"
	"slices"
)

type day13Map [][]byte

func day13part1(filename string) (string, error) {
	maps, err := day13ReadMaps(filename)
	if err != nil {
		return "", err
	}
	var cols, rows int
	for i, m := range maps {
		day13PrintMap(m)
		if row, ok := m.findRowReflection(); ok {
			log.Print(i, " Row ", row)
			rows += 100 * row
			continue
		}
		if column, ok := m.findColumnReflection(); ok {
			log.Print(i, " Column ", column)
			cols += column
			continue
		}
	}
	return fmt.Sprint(cols + rows), nil
}

func day13PrintMap(u day13Map) {
	for _, vy := range u {
		for _, vx := range vy {
			fmt.Print(string(vx))
		}
		fmt.Println()
	}
}

func day13ReadMaps(filename string) ([]day13Map, error) {
	var maps []day13Map
	var m day13Map
	if err := forLineError(filename, func(line string) error {
		if len(line) == 0 {
			maps = append(maps, m)
			m = make(day13Map, 0)
			return nil
		}
		m = append(m, []byte(line))
		return nil
	}); err != nil {
		return maps, err
	}
	maps = append(maps, m)
	return maps, nil
}

// Returns 0-based index of row and true if reflected.
func (m day13Map) findRowReflection() (int, bool) {
	row := 1
	for row < len(m) {
		res := true
		for b, t := row-1, row; b >= 0 && t < len(m); b, t = b-1, t+1 {
			res = res && slices.Equal(m[b], m[t])
			if !res {
				break
			}
		}
		if res {
			return row, true
		}
		row += 1
	}
	return 0, false
}

// Returns 0-based index of column and true if reflected.
func (m day13Map) findColumnReflection() (int, bool) {
	column := 1
	for column < len(m[0]) {
		res := true
		for l, r := column-1, column; l >= 0 && r < len(m[0]); l, r = l-1, r+1 {
			res = res && m.colsEqual(l, r)
			if !res {
				break
			}
		}
		if res {
			return column, true
		}
		column += 1
	}
	return 0, false
}

func (m day13Map) colsEqual(l, r int) bool {
	for y := 0; y < len(m); y++ {
		if m[y][l] != m[y][r] {
			return false
		}
	}

	return true
}
