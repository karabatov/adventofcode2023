package main

import (
	"fmt"
	"log"
	"slices"
)

func day22part2(filename string) (string, error) {
	bricks, err := day22ReadPlan(filename)
	if err != nil {
		return "", err
	}
	fall, _ := day22Fall(bricks)
	var total int
	for i := range fall {
		newFall := slices.Clone(fall)
		newFall = slices.Delete(newFall, i, i+1)
		_, fallen := day22Fall(newFall)
		log.Printf("Brick %d causes %d bricks to fall", i+1, fallen)
		total += fallen
	}
	return fmt.Sprint(total), nil
}
