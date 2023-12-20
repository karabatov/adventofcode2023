package main

import (
	"fmt"
	"slices"
)

func day20part2(filename string) (string, error) {
	f, err := day20ReadMod(filename)
	if err != nil {
		return "", err
	}
	st := day20State{
		sys:  f,
		flip: map[string]bool{},
		conj: map[string]map[string]int{},
	}
	var dirty = map[string]int{}
	for _, v := range st.dirtyModules() {
		dirty[v] = 0
	}
	var count int
	for {
		count += 1
		st.pressButton(func(to string, val int) bool {
			if _, ok := dirty[to]; ok && val == 0 {
				dirty[to] = count
			}
			return false
		})
		var cyc int
		for _, v := range dirty {
			if v > 1 {
				cyc += 1
			}
		}
		if cyc == len(dirty) {
			break
		}
	}
	res := 1
	for _, v := range dirty {
		res *= v
	}
	return fmt.Sprint(res), nil
}

func (st day20State) dirtyModules() []string {
	for _, v := range st.sys.mods {
		if slices.Contains(v.output, "rx") {
			return st.sys.inp[v.name]
		}
	}
	return nil
}
