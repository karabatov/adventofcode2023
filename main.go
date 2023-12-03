package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// DayFunc takes the input filename as input and returns the answer and error.
type DayFunc = func(string) (string, error)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Pass the day and input file name as arguments.")
	}

	var dayFn DayFunc
	switch os.Args[1] {
	case "11":
		dayFn = day1part1
	case "12":
		dayFn = day1part2
	case "21":
		dayFn = day2part1
	case "22":
		dayFn = day2part2
	case "31":
		dayFn = day3part1
	default:
		log.Fatalf("Invalid day identifier '%s'", os.Args[1])
	}

	startTime := time.Now()
	res, err := dayFn(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	endTime := time.Now()

	log.Printf("Answer: %s", res)
	log.Printf("Run time: %v", endTime.Sub(startTime))
}

func forLine(filename string, run func(string)) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not open file: %s", filename)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		run(scanner.Text())
	}

	return nil
}

func allLines(filename string) ([]string, error) {
	lines := make([]string, 0)

	if err := forLine(filename, func(line string) {
		lines = append(lines, line)
	}); err != nil {
		return nil, err
	}

	return lines, nil
}
