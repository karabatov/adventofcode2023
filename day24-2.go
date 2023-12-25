package main

import "fmt"

func day24part2(filename string) (string, error) {
	hails, err := day24ReadList(filename)
	if err != nil {
		return "", err
	}
	minP := day24Vec3{200000000000000, 200000000000000, 0}
	maxP := day24Vec3{400000000000000, 400000000000000, 0}
	var longs []day24Hail
	for _, v := range hails {
		s := v.extendToArea(minP, maxP)
		longs = append(longs, s)
	}
	var total int
	for i := 0; i < len(longs); i++ {
		for ii := i + 1; ii < len(longs); ii++ {
			if xPoint, ok := longs[i].intersect2D(longs[ii]); ok {
				if xPoint.inArea2D(minP, maxP) {
					total += 1
				}
			}
		}
	}
	return fmt.Sprint(total), nil
}
