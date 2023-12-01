package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type DayFunc = func(string) (string, error)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Pass the day and input file name as arguments.")
	}

	var dayFn DayFunc
	switch os.Args[1] {
	case "11":
		dayFn = day1part1
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
		return fmt.Errorf("Could not open file: %s", filename)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		run(scanner.Text())
	}

	return nil
}