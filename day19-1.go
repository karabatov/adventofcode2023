package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// x, m, a, s is 0, 1, 2, 3.
type day19Part []int

type day19Cond int

const (
	day19CondPass = iota
	day19CondLess
	day19CondMore
)

type day19Rule struct {
	// Result workflow.
	out  string
	cond day19Cond
	// 0–3.
	part int
	// Comparison condition.
	cmp int
}

type day19RuleSet map[string][]day19Rule

type day19System struct {
	rules day19RuleSet
	parts []day19Part
}

func day19part1(filename string) (string, error) {
	f, err := day19ReadPlan(filename)
	if err != nil {
		return "", err
	}
	var acc []day19Part
	for _, p := range f.parts {
		if out := f.process(p); out == "A" {
			acc = append(acc, p)
		}
	}
	var total int
	for _, v := range acc {
		for _, vv := range v {
			total += vv
		}
	}
	return fmt.Sprint(total), nil
}

func (s day19System) process(p day19Part) string {
	next := "in"
	for next != "A" && next != "R" {
		for _, r := range s.rules[next] {
			if out, pass := r.result(p); pass {
				next = out
				break
			}
		}
	}
	return next
}

// Returns the result workflow and if the rule succeeded.
func (r day19Rule) result(p day19Part) (string, bool) {
	res := false
	switch r.cond {
	case day19CondPass:
		res = true
	case day19CondLess:
		res = p[r.part] < r.cmp
	case day19CondMore:
		res = p[r.part] > r.cmp
	}
	return r.out, res
}

var (
	day19RuleRegex    = regexp.MustCompile(`([[:alpha:]]+){(.*)}`)
	day19OneRuleRegex = regexp.MustCompile(`([xmas])([<>])(\d+):([[:alpha:]]+)`)
	day19PartRegex    = regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)
)

func day19ReadPlan(filename string) (day19System, error) {
	var rulesDone bool
	res := day19System{
		rules: day19RuleSet{},
		parts: []day19Part{},
	}
	if err := forLineError(filename, func(line string) error {
		if len(line) == 0 {
			rulesDone = true
			return nil
		}
		if !rulesDone {
			m := day19RuleRegex.FindStringSubmatch(line)
			rules, err := day19ReadRules(m[2])
			if err != nil {
				return err
			}
			res.rules[m[1]] = rules
			return nil
		}
		p, err := day19ReadPart(line)
		if err != nil {
			return err
		}
		res.parts = append(res.parts, p)
		return nil
	}); err != nil {
		return res, err
	}
	return res, nil
}

func day19ReadPart(s string) (day19Part, error) {
	var p day19Part
	m := day19PartRegex.FindStringSubmatch(s)
	if len(m) != 5 {
		return p, fmt.Errorf("bad part: %s", s)
	}
	for i := 1; i < 5; i++ {
		num, err := strconv.Atoi(m[i])
		if err != nil {
			return p, err
		}
		p = append(p, num)
	}
	return p, nil
}

func day19ReadRules(r string) ([]day19Rule, error) {
	var res []day19Rule
	for _, v := range strings.Split(r, ",") {
		rule, err := day19MakeRule(v)
		if err != nil {
			return nil, err
		}
		res = append(res, rule)
	}
	return res, nil
}

var day19PartMap = map[string]int{
	"x": 0,
	"m": 1,
	"a": 2,
	"s": 3,
}

var day19CondMap = map[string]day19Cond{
	"<": day19CondLess,
	">": day19CondMore,
}

func day19MakeRule(s string) (day19Rule, error) {
	var r day19Rule
	if spl := day19OneRuleRegex.FindStringSubmatch(s); len(spl) == 5 {
		r.out = spl[4]
		r.cond = day19CondMap[spl[2]]
		r.part = day19PartMap[spl[1]]
		n, err := strconv.Atoi(spl[3])
		if err != nil {
			return r, err
		}
		r.cmp = n
	} else {
		r.out = s
	}
	return r, nil
}
