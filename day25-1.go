package main

import (
	"fmt"
	"log"
	"slices"
	"strings"
)

// Draw with:
// go run . 251 input25-2.txt | cluster -C2 | unflatten | dot -Tsvg > day25-1.svg
// go run . 251 input25-2.txt | cluster -C2 | neato -Tsvg > day25-3.svg

type day25Store map[string][]string

func day25part1(filename string) (string, error) {
	list, err := day25ReadList(filename)
	if err != nil {
		return "", err
	}
	// exp := list.expand([]string{"hfx", "bvb", "pzl", "nvd", "cmg", "jqt"})
	exp := list.expand([]string{"kzh", "dgt", "ddc", "rks", "tnz", "gqm"})
	// exp.print()
	one := exp.visited("kzh")
	two := exp.visited("rks")
	log.Print(one, " ", two)
	return fmt.Sprint(one * two), nil
}

func (s day25Store) visited(from string) int {
	seen := map[string]bool{}
	q := []string{from}
	for len(q) > 0 {
		next := q[0]
		q = q[1:]
		for _, v := range s[next] {
			if !seen[v] {
				q = append(q, v)
				seen[v] = true
			}
		}
	}
	return len(seen)
}

func (s day25Store) expand(disc []string) day25Store {
	res := day25Store{}
	for k, v := range s {
		hasK := slices.Contains(disc, k)
		var conn []string
		for _, vv := range v {
			if hasVV := slices.Contains(disc, vv); hasVV && hasK {
				continue
			}
			conn = append(conn, vv)
			if !slices.Contains(res[vv], k) {
				res[vv] = append(res[vv], k)
			}
		}
		for _, v := range conn {
			if !slices.Contains(res[k], v) {
				res[k] = append(res[k], v)
			}
		}
	}
	return res
}

func (s day25Store) print() {
	fmt.Println("graph {")
	for k, v := range s {
		for _, vv := range v {
			if (k == "kzh" && vv == "rks") ||
				(k == "dgt" && vv == "tnz") ||
				(k == "ddc" && vv == "gqm") {
				continue
			}
			fmt.Printf("%s -- %s\n", k, vv)
		}
	}
	fmt.Println("}")
}

func day25ReadList(filename string) (day25Store, error) {
	s := make(map[string][]string)
	if err := forLineError(filename, func(line string) error {
		spl := strings.Split(line, ": ")
		s[spl[0]] = strings.Split(spl[1], " ")
		return nil
	}); err != nil {
		return s, err
	}
	return s, nil
}
