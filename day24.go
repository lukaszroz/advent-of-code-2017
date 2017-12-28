package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type component struct {
	a, b uint
}

type data struct {
	components                map[uint][]component
	isUsed                    map[component]bool
	best, current             uint
	bestLength, currentLength uint
}

func main() {
	in := input()
	fmt.Println("Star 1: ", findBest(&in, 0))
	in.best, in.current = 0, 0
	fmt.Println("Star 1: ", findLongest(&in, 0))
}

func findLongest(d *data, n uint) uint {
	for _, c := range d.components[n] {
		if d.isUsed[c] {
			continue
		}
		d.isUsed[c] = true
		d.current += c.a + c.b
		d.currentLength++
		if d.currentLength > d.bestLength || (d.currentLength == d.bestLength && d.current > d.best) {
			d.bestLength = d.currentLength
			d.best = d.current
		}
		if n == c.a {
			findLongest(d, c.b)
		} else {
			findLongest(d, c.a)
		}
		d.currentLength--
		d.current -= c.a + c.b
		delete(d.isUsed, c)

	}
	return d.best
}

func findBest(d *data, n uint) uint {
	for _, c := range d.components[n] {
		if d.isUsed[c] {
			continue
		}
		d.isUsed[c] = true
		d.current += c.a + c.b
		if d.current > d.best {
			d.best = d.current
		}
		if n == c.a {
			findBest(d, c.b)
		} else {
			findBest(d, c.a)
		}
		d.current -= c.a + c.b
		delete(d.isUsed, c)

	}
	return d.best
}

func input() data {
	var d data
	d.components = make(map[uint][]component)
	d.isUsed = make(map[component]bool)
	lines := getLines()
	for _, line := range lines {
		var c component
		fmt.Sscanf(line, "%d/%d", &c.a, &c.b)
		d.components[c.a] = append(d.components[c.a], c)
		d.components[c.b] = append(d.components[c.b], c)
	}
	return d
}

func getLines() []string {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(bytes), "\n")
	return strs
}
