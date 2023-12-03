package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Extract numbers from a line.
var day3NumbersRe = regexp.MustCompile(`\d+`)

// day3Point represents a point in the schematic through indices.
type day3Point struct {
	idxY, idxX int
}

// day3Number represents a number by its position in the schematic.
type day3Number struct {
	day3Point
	len int
}

func (d day3Number) value(lines []string) (int, error) {
	str := lines[d.idxY][d.idxX : d.idxX+d.len]
	return strconv.Atoi(str)
}

func day3part1(filename string) (string, error) {
	lines, err := allLines(filename)
	if err != nil {
		return "", err
	}
	numsWithSymbols := make([]day3Number, 0)
	for _, num := range day3AllNumbers(lines) {
		day3WalkAroundNumber(num, lines, func(b byte) bool {
			if b != '.' {
				numsWithSymbols = append(numsWithSymbols, num)
				return false
			}

			return true
		})
	}

	var total int
	for _, num := range numsWithSymbols {
		val, err := num.value(lines)
		if err != nil {
			return "", nil
		}
		total += val
	}
	return fmt.Sprint(total), nil
}

func day3AllNumbers(lines []string) []day3Number {
	nums := make([]day3Number, 0)
	for idxY, line := range lines {
		nums = append(nums, day3NumbersFromLine(line, idxY)...)
	}
	return nums
}

func day3NumbersFromLine(line string, idxY int) []day3Number {
	nums := make([]day3Number, 0)
	matches := day3NumbersRe.FindAllStringSubmatchIndex(line, -1)
	for _, match := range matches {
		nums = append(nums, day3Number{
			day3Point: day3Point{
				idxY: idxY,
				idxX: match[0],
			},
			len: match[1] - match[0],
		})
	}

	return nums
}

// Runs while test is true.
func day3WalkAroundNumber(num day3Number, lines []string, test func(byte) bool) {
	for y := max(0, num.idxY-1); y < min(len(lines), num.idxY+2); y++ {
		for x := max(0, num.idxX-1); x < min(len(lines[num.idxY]), num.idxX+num.len+1); x++ {
			if y == num.idxY && x >= num.idxX && x < num.idxX+num.len {
				continue
			}
			if !test(lines[y][x]) {
				return
			}
		}
	}
}
