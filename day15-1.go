package main

import (
	"fmt"
	"strings"
)

func day15part1(filename string) (string, error) {
	var l string
	if err := forLineError(filename, func(line string) error {
		l = line
		return nil
	}); err != nil {
		return "", err
	}
	s := day15Splits(l)
	var sum int
	for _, v := range s {
		sum += day15Hash(v)
	}
	return fmt.Sprint(sum), nil
}

func day15Splits(l string) [][]byte {
	s := strings.Split(l, ",")
	var res [][]byte
	for _, v := range s {
		res = append(res, []byte(v))
	}
	return res
}

func day15Hash(b []byte) int {
	var res int
	for _, v := range b {
		res += int(v)
		res *= 17
		res %= 256
	}
	return res
}
