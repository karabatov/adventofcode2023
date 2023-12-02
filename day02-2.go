package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func day2part2(filename string) (string, error) {
	var total int

	if err := forLine(filename, func(line string) {
		matches := day2GameIDRe.FindStringSubmatch(line)
		minCubes := day2MinCubesForGame(matches[2])
		total += day2GamePower(minCubes)
	}); err != nil {
		return "", err
	}

	return fmt.Sprint(total), nil
}

func day2GamePower(cubes day2CubeSet) int {
	res := 1
	for _, v := range cubes {
		res *= v
	}
	return res
}

func day2MinCubesForGame(game string) day2CubeSet {
	var minCubes day2CubeSet
	// A handful of cubes.
	for _, pick := range strings.Split(game, "; ") {
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
			if minCubes[idx] < num {
				minCubes[idx] = num
			}
		}
	}
	return minCubes
}
