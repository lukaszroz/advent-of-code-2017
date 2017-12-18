package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input := getInput()
	array := newCircular(16)
	for _, in := range input {
		dance(array, in)
	}
	fmt.Println("Star 1: ", array)

	cycle := findCycle(input)
	n := 1000000000 % cycle
	array = newCircular(16)
	for i := 0; i < n; i++ {
		for _, in := range input {
			dance(array, in)
		}
	}
	fmt.Println("Star 2: ", array)
}

func findCycle(input []string) int {
	array := newCircular(16)
	initial := array.String()
	for _, in := range input {
		dance(array, in)
	}
	for i := 1; i < 1000000000; i++ {
		for _, in := range input {
			if initial == array.String() {
				return i
			}
			dance(array, in)
		}
	}
	panic("Not found!")
}

func dance(c *circural, in string) {
	move := in[0]
	switch move {
	case 's':
		var sp int
		fmt.Sscanf(in, "s%d", &sp)
		c.spin(sp)
	case 'x':
		var a, b int
		fmt.Sscanf(in, "x%d/%d", &a, &b)
		c.exchange(a, b)
	case 'p':
		c.swap(in[1], in[3])
	}
}

type circural struct {
	programs []byte
	first    int
}

func (c *circural) spin(spin int) {
	c.first += len(c.programs) - spin
	c.first %= len(c.programs)
}

func (c *circural) exchange(a, b int) {
	a_idx := (a + c.first) % len(c.programs)
	b_idx := (b + c.first) % len(c.programs)
	c._swap(a_idx, b_idx)
}

func (c *circural) _swap(a, b int) {
	c.programs[a], c.programs[b] = c.programs[b], c.programs[a]
}

func (c *circural) swap(a, b byte) {
	c._swap(c._find(a), c._find(b))
}

func (c *circural) _find(a byte) int {
	for idx, r := range c.programs {
		if r == a {
			return idx
		}
	}
	panic(errors.New("Not Found!"))
}

func newCircular(size int) *circural {
	c := new(circural)
	c.programs = make([]byte, size)
	for idx := range c.programs {
		c.programs[idx] = 'a' + byte(idx)
	}
	return c
}

func (c circural) String() string {
	bs := make([]byte, len(c.programs))
	for idx := range bs {
		bs[idx] = c.programs[(c.first+idx)%len(c.programs)]
	}
	return string(bs)
}

func getInput() []string {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(bytes), ",")
	return strs
}
