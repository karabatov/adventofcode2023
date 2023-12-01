package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Pass the input file name as argument.")
	}

	var total int

	if err := forLine(os.Args[1], func(line string) {
		first := -1
		last := -1
		for _, r := range line {
			if unicode.IsDigit(r) {
				if first < 0 {
					first = int(r - '0')
				}
				last = int(r - '0')
			}
		}
		total += first*10 + last
	}); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total: %d", total)
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
