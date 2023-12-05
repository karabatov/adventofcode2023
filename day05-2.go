package main

import (
	"fmt"
	"slices"
)

func day5part2(filename string) (string, error) {
	var seeds []int
	var ranges []day5Range
	if err := forLineError(filename, func(line string) error {
		if len(seeds) == 0 {
			matches := day5SeedsRe.FindStringSubmatch(line)
			seedRanges := parseNumberLine(matches[1])
			// TODO: Map ranges instead of individual items.
			for i := 0; i < len(seedRanges)-1; i += 2 {
				for j := seedRanges[i]; j <= seedRanges[i]+seedRanges[i+1]; j++ {
					seeds = append(seeds, j)
				}
			}
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
