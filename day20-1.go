package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type day20Kind int

const (
	day20KindNone = iota
	day20KindFlip
	day20KindConj
	day20KindCast
)

type day20Mod struct {
	name   string
	kind   day20Kind
	output []string
}

// 0 is low and 1 is high.
type day20Pulse struct {
	from, to string
	val      int
}

type day20System struct {
	mods map[string]day20Mod
	// Inputs for all inverters.
	inp map[string][]string
}

type day20State struct {
	sys  day20System
	flip map[string]bool
	conj map[string]map[string]int
}

func day20part1(filename string) (string, error) {
	f, err := day20ReadMod(filename)
	if err != nil {
		return "", err
	}
	st := day20State{
		sys:  f,
		flip: map[string]bool{},
		conj: map[string]map[string]int{},
	}
	res := make(map[int]int)
	for i := 0; i < 1000; i++ {
		next := st.pressButton(func(to string, val int) bool {
			return false
		})
		res[0] += next[0]
		res[1] += next[1]
	}
	return fmt.Sprintf("%d * %d = %d", res[0], res[1], res[0]*res[1]), nil
}

// Returns the number of low and high pulses with keys 0 and 1.
func (st day20State) pressButton(stop func(to string, val int) bool) map[int]int {
	res := make(map[int]int)
	q := []day20Pulse{{
		from: "button",
		to:   "broadcaster",
		val:  0,
	}}
	for len(q) > 0 {
		pul := q[0]
		q = q[1:]
		res[pul.val] += 1
		if stop(pul.to, pul.val) {
			return res
		}
		mod, ok := st.sys.mods[pul.to]
		if !ok {
			continue
		}
		switch mod.kind {
		case day20KindCast:
			q = append(q, day20Send(mod, pul.val)...)
		case day20KindFlip:
			if pul.val == 1 {
				continue
			}
			if st.flip[mod.name] {
				st.flip[mod.name] = false
				q = append(q, day20Send(mod, 0)...)
			} else {
				st.flip[mod.name] = true
				q = append(q, day20Send(mod, 1)...)
			}
		case day20KindConj:
			np := st.updateConj(mod.name, pul.val, pul.from)
			q = append(q, day20Send(mod, np)...)
		}
	}
	return res
}

// Returns the next pulse it should send.
func (st day20State) updateConj(name string, with int, from string) int {
	m, ok := st.conj[name]
	if !ok {
		m = make(map[string]int)
	}
	m[from] = with
	var pos int
	for _, v := range st.sys.inp[name] {
		if m[v] == 1 {
			pos += 1
		}
	}
	st.conj[name] = m
	if pos == len(st.sys.inp[name]) {
		return 0
	}
	return 1
}

func day20Send(mod day20Mod, val int) []day20Pulse {
	var res []day20Pulse
	for _, v := range mod.output {
		res = append(res, day20Pulse{
			from: mod.name,
			to:   v,
			val:  val,
		})
	}
	return res
}

var day20ModRegex = regexp.MustCompile(`([&%]*)([[:alpha:]]+) -> (.*)`)

func day20ReadMod(filename string) (day20System, error) {
	res := day20System{
		mods: map[string]day20Mod{},
		inp:  map[string][]string{},
	}
	if err := forLineError(filename, func(line string) error {
		m := day20ModRegex.FindStringSubmatch(line)
		mod := day20Mod{
			name:   m[2],
			kind:   day20ModKind(m[1]),
			output: strings.Split(m[3], ", "),
		}
		res.mods[m[2]] = mod
		return nil
	}); err != nil {
		return res, err
	}
	res.setInputs()
	return res, nil
}

func (s day20System) setInputs() {
	for _, v := range s.mods {
		if v.kind != day20KindConj {
			continue
		}
		var res []string
		for _, o := range s.mods {
			if slices.Contains(o.output, v.name) {
				res = append(res, o.name)
			}
		}
		s.inp[v.name] = res
	}
}

func day20ModKind(s string) day20Kind {
	switch s {
	case "":
		return day20KindCast
	case "%":
		return day20KindFlip
	case "&":
		return day20KindConj
	default:
		return day20KindNone
	}
}
