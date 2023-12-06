package main

import (
	"fmt"
	"regexp"
	"slices"
)

var (
	day5SeedsRe = regexp.MustCompile(`^seeds: (.+)$`)
)

type day5Range struct {
	dst, src, len int
}

// Returns value and true if seed in range, otherwise 0 and false.
func (r day5Range) mapSeed(seed int) (int, bool) {
	return r.dst + (seed - r.src), seed >= r.src && seed < r.src+r.len
}

func day5part1(filename string) (string, error) {
	var seeds []int
	var ranges []day5Range
	if err := forLineError(filename, func(line string) error {
		if len(seeds) == 0 {
			matches := day5SeedsRe.FindStringSubmatch(line)
			seeds = parseNumberLine(matches[1])
			return nil
		}

		if line == "" {
			if len(ranges) == 0 {
				return nil
			}
			for i, s := range seeds {
				seeds[i] = day5RunMap(s, ranges)
			}
			return nil
		}

		if line[len(line)-1:] == ":" {
			ranges = nil
			return nil
		}

		p := parseNumberLine(line)
		ranges = append(ranges, day5Range{dst: p[0], src: p[1], len: p[2]})

		return nil
	}); err != nil {
		return "", err
	}
	// Run one final time after the last line.
	for i, s := range seeds {
		seeds[i] = day5RunMap(s, ranges)
	}

	return fmt.Sprint(slices.Min(seeds)), nil
}

func day5RunMap(seed int, ranges []day5Range) int {
	for _, r := range ranges {
		if mapped, ok := r.mapSeed(seed); ok {
			return mapped
		}
	}

	return seed
}
