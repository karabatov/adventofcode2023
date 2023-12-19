package main

import (
	"fmt"
	"slices"
)

type day19Range struct {
	lo, hi []int
}

type day19Set struct {
	flow string
	val  day19Range
}

func day19part2(filename string) (string, error) {
	f, err := day19ReadPlan(filename)
	if err != nil {
		return "", err
	}
	var acc []day19Range
	q := []day19Set{{
		flow: "in",
		val: day19Range{
			lo: []int{1, 1, 1, 1},
			hi: []int{4000, 4000, 4000, 4000},
		},
	}}
	for len(q) > 0 {
		next := q[0]
		q = q[1:]
		sets := []day19Set{next}
		for _, rule := range f.rules[next.flow] {
			var nextSets []day19Set
			for _, set := range sets {
				for _, nextSet := range rule.split(set) {
					switch nextSet.flow {
					case "A":
						acc = append(acc, nextSet.val)
					case "R":
						continue
					case next.flow:
						nextSets = append(nextSets, nextSet)
					default:
						q = append(q, nextSet)
					}
				}
			}
			sets = nextSets
		}
	}
	var total int
	for _, v := range acc {
		one := 1
		for i := 0; i < len(v.lo); i++ {
			one *= v.hi[i] - v.lo[i] + 1
		}
		total += one
	}
	return fmt.Sprint(total), nil
}

func (r day19Rule) split(set day19Set) []day19Set {
	var res []day19Set
	switch r.cond {
	case day19CondPass:
		res = append(res, day19Set{
			flow: r.out,
			val:  set.val,
		})
	case day19CondLess:
		if set.val.lo[r.part] < r.cmp && set.val.hi[r.part] >= r.cmp {
			setLo := day19Set{
				flow: r.out,
				val:  set.val.with(r.part, set.val.lo[r.part], r.cmp-1),
			}
			setHi := day19Set{
				flow: set.flow,
				val:  set.val.with(r.part, r.cmp, set.val.hi[r.part]),
			}
			res = append(res, setLo, setHi)
		} else if set.val.lo[r.part] >= r.cmp {
			res = append(res, set)
		} else if set.val.hi[r.part] < r.cmp {
			res = append(res, day19Set{
				flow: r.out,
				val:  set.val,
			})
		}
	case day19CondMore:
		if set.val.lo[r.part] <= r.cmp && set.val.hi[r.part] > r.cmp {
			setLo := day19Set{
				flow: set.flow,
				val:  set.val.with(r.part, set.val.lo[r.part], r.cmp),
			}
			setHi := day19Set{
				flow: r.out,
				val:  set.val.with(r.part, r.cmp+1, set.val.hi[r.part]),
			}
			res = append(res, setLo, setHi)
		} else if set.val.lo[r.part] <= r.cmp {
			res = append(res, set)
		} else if set.val.hi[r.part] > r.cmp {
			res = append(res, day19Set{
				flow: r.out,
				val:  set.val,
			})
		}
	}
	return res
}

func (r day19Range) with(part, lo, hi int) day19Range {
	res := day19Range{
		lo: slices.Clone(r.lo),
		hi: slices.Clone(r.hi),
	}
	res.lo[part] = lo
	res.hi[part] = hi
	return res
}
