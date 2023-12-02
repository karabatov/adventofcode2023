package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// The order is red (0), green (1), blue (2)
type day2CubeSet [3]int

var (
	// ID of the game.
	day2GameIDRe = regexp.MustCompile(`^Game (\d+): (.*)$`)
	// One color of cube.
	day2NumColorRe = regexp.MustCompile(`(\d+) (\w+)`)
)

func day2part1(filename string) (string, error) {
	var total int

	compare := day2CubeSet{12, 13, 14}

	if err := forLine(filename, func(line string) {
		matches := day2GameIDRe.FindStringSubmatch(line)
		gameID, _ := strconv.Atoi(matches[1])
		var maxCubes day2CubeSet
		// A handful of cubes.
		for _, pick := range strings.Split(matches[2], "; ") {
			// One color of cubes within a handful.
			for _, one := range strings.Split(pick, ", ") {
				oneMatches := day2NumColorRe.FindStringSubmatch(one)
				num, _ := strconv.Atoi(oneMatches[1])
				var idx int
				switch oneMatches[2] {
				case "red":
					idx = 0
				case "green":
					idx = 1
				case "blue":
					idx = 2
				default:
					log.Fatalf("unknown color: %s", oneMatches[2])
				}
				if maxCubes[idx] < num {
					maxCubes[idx] = num
				}
			}
		}
		if day2part1CanPlay(maxCubes, compare) {
			total += gameID
		}
	}); err != nil {
		return "", err
	}

	return fmt.Sprint(total), nil
}

// day2part1CanPlay returns true if the game compare can be played with the set maxCubes.
func day2part1CanPlay(maxCubes day2CubeSet, compare day2CubeSet) bool {
	res := true
	for i := 0; i <= 2; i++ {
		res = res && (maxCubes[i] <= compare[i])
	}
	return res
}
