package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type pattern [][]bool

type rule struct {
	input  pattern
	output pattern
}

func main() {
	rules := parse(getInput())
	m := newMap(rules)

	art := [][][]bool{{{false, true, false}, {false, false, true}, {true, true, true}}}

	for i := 0; i < 5; i++ {
		art = next(m, art)
	}
	fmt.Println("Star 1:", count(art))
	for i := 5; i < 18; i++ {
		art = next(m, art)
	}
	fmt.Println("Star 2:", count(art))
}

func count(art [][][]bool) int {
	count := 0
	for _, p := range art {
		for _, row := range p {
			for _, b := range row {
				if b {
					count++
				}
			}
		}
	}
	return count
}

func next(m map[uint16][][]bool, in [][][]bool) (out [][][]bool) {
	for _, pat := range in {
		if len(pat) == 4 {
			p1 := make([][]bool, 2)
			p1[0] = pat[0][:2]
			p1[1] = pat[1][:2]
			p1 = m[encode(p1)]

			p2 := make([][]bool, 2)
			p2[0] = pat[0][2:]
			p2[1] = pat[1][2:]
			p2 = m[encode(p2)]

			p3 := make([][]bool, 2)
			p3[0] = pat[2][:2]
			p3[1] = pat[3][:2]
			p3 = m[encode(p3)]

			p4 := make([][]bool, 2)
			p4[0] = pat[2][2:]
			p4[1] = pat[3][2:]
			p4 = m[encode(p4)]

			combined := make([][]bool, 6)
			for i := 0; i < 3; i++ {
				combined[i] = p1[i]
				combined[i] = append(combined[i], p2[i]...)
			}
			for i := 3; i < 6; i++ {
				combined[i] = p3[i-3]
				combined[i] = append(combined[i], p4[i-3]...)
			}
			out = append(out, combined)
		} else if len(pat) == 6 {
			for i := 0; i < 3; i++ {
				p := make([][]bool, 2)
				p[0] = pat[i*2][:2]
				p[1] = pat[i*2+1][:2]
				out = append(out, m[encode(p)])

				p = make([][]bool, 2)
				p[0] = pat[i*2][2:4]
				p[1] = pat[i*2+1][2:4]
				out = append(out, m[encode(p)])

				p = make([][]bool, 2)
				p[0] = pat[i*2][4:]
				p[1] = pat[i*2+1][4:]
				out = append(out, m[encode(p)])
			}
		} else {
			out = append(out, m[encode(pat)])
		}
	}
	return
}

func newMap(rules []rule) map[uint16][][]bool {
	m := make(map[uint16][][]bool)
	for _, r := range rules {
		for i := 0; i < 4; i++ {
			m[encode(rotate(r.input, i))] = r.output
			m[encode(flip(rotate(r.input, i)))] = r.output
		}
	}
	return m
}

func encode(in [][]bool) (out uint16) {
	if len(in) > 2 {
		out |= (1 << 15)
	}
	var i uint
	for _, row := range in {
		for _, b := range row {
			if b {
				out |= (1 << i)
			}
			i++
		}
	}
	return
}

func rotate(in [][]bool, n int) (out [][]bool) {
	if n == 0 {
		out = in
	} else if n == 1 {
		out = make([][]bool, len(in))
		for i := range in {
			out[i] = make([]bool, len(in))
		}
		if len(in) == 2 {
			out[0][0] = in[1][0]
			out[0][1] = in[0][0]
			out[1][0] = in[1][1]
			out[1][1] = in[0][1]
		} else {
			out[0][0] = in[2][0]
			out[0][1] = in[1][0]
			out[0][2] = in[0][0]
			out[1][0] = in[2][1]
			out[1][1] = in[1][1]
			out[1][2] = in[0][1]
			out[2][0] = in[2][2]
			out[2][1] = in[1][2]
			out[2][2] = in[0][2]
		}
	} else {
		out = rotate(rotate(in, n-1), 1)
	}
	return
}

func flip(in [][]bool) (out [][]bool) {
	out = make([][]bool, len(in))
	for i := range in {
		out[i] = make([]bool, len(in))
		copy(out[i], in[i])
	}
	for i := range in {
		out[i][0], out[i][len(out)-1] = out[i][len(out)-1], out[i][0]
	}
	return
}

func (p pattern) String() string {
	var s []byte
	for _, row := range p {
		for _, b := range row {
			if b {
				s = append(s, '#')
			} else {
				s = append(s, '.')
			}
		}
		s = append(s, '\n')
	}
	return string(s)
}

func (r rule) String() string {
	return r.input.String() + r.output.String()
}

func parse(lines []string) (rules []rule) {
	rules = make([]rule, len(lines))
	for idx, line := range lines {
		if len(line) > 20 {
			rules[idx] = newRule(3, 4)
			parseRule(&rules[idx].input, line[:11])
			parseRule(&rules[idx].output, line[15:])
		} else {
			rules[idx] = newRule(2, 3)
			parseRule(&rules[idx].input, line[:5])
			parseRule(&rules[idx].output, line[9:])
		}
	}
	return
}

func parseRule(r *pattern, s string) {
	i := 0
	for row := range *r {
		for col := range (*r)[row] {
			if s[i] == '#' {
				(*r)[row][col] = true
			}
			i++
		}
		i++
	}
}

func newRule(in, out int) rule {
	var r rule
	r.input = make([][]bool, in)
	for j := range r.input {
		r.input[j] = make([]bool, in)
	}
	r.output = make([][]bool, out)

	for j := range r.output {
		r.output[j] = make([]bool, out)
	}
	return r
}

func getInput() []string {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(bytes), "\n")
	return strs
}
