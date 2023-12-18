package main

import (
	"fmt"
	"strconv"
)

func day18part2(filename string) (string, error) {
	f, err := day18ReadPlan2(filename)
	if err != nil {
		return "", err
	}
	tr := day18Draw(f)
	pp := day18Points(tr)
	area := day18Area(pp)
	return fmt.Sprint(area), nil
}

func day18ReadPlan2(filename string) ([]day18Dig, error) {
	cm := map[byte]string{
		'0': "R",
		'1': "D",
		'2': "L",
		'3': "U",
	}
	var res []day18Dig
	if err := forLineError(filename, func(line string) error {
		c := day18Regex.FindStringSubmatch(line)[3]
		num, err := strconv.ParseInt(c[:5], 16, 32)
		if err != nil {
			return err
		}
		res = append(res, day18Dig{
			dir: cm[c[5]],
			num: int(num),
			col: c,
		})
		return nil
	}); err != nil {
		return res, err
	}
	return res, nil
}
