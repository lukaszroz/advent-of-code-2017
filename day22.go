package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type state byte

const (
	Clean state = iota
	Weakened
	Infected
	Flagged
)

type position [2]int
type direction position

type virus struct {
	p position
	d direction
}

func main() {
	in := getInput()

	fmt.Println("Star 1:", countStage1(in, 10000))
	fmt.Println("Star 2:", countStage2(in, 10000000))
}

func countStage2(in []string, iterations int) int {
	n := len(in)
	infectionMap := parse(in)
	infected := 0
	virus := virus{[2]int{n / 2, n / 2}, [2]int{-1, 0}}
	for i := 0; i < iterations; i++ {
		switch infectionMap[virus.p] {
		case Clean:
			virus.turnLeft()
			infectionMap[virus.p] = Weakened
		case Weakened:
			infectionMap[virus.p] = Infected
			infected++
		case Infected:
			virus.turnRight()
			infectionMap[virus.p] = Flagged
		case Flagged:
			virus.turnLeft()
			virus.turnLeft()
			delete(infectionMap, virus.p)
		}
		virus.move()
	}
	return infected
}

func countStage1(in []string, iterations int) int {
	n := len(in)
	infectionMap := parse(in)
	infected := 0
	virus := virus{[2]int{n / 2, n / 2}, [2]int{-1, 0}}
	for i := 0; i < iterations; i++ {
		if infectionMap[virus.p] == Infected {
			virus.turnRight()
			delete(infectionMap, virus.p)
		} else {
			virus.turnLeft()
			infectionMap[virus.p] = Infected
			infected++
		}
		virus.move()
	}
	return infected
}

var rightMap map[[2]int][2]int = map[[2]int][2]int{
	[2]int{-1, 0}: [2]int{0, 1},
	[2]int{0, 1}:  [2]int{1, 0},
	[2]int{1, 0}:  [2]int{0, -1},
	[2]int{0, -1}: [2]int{-1, 0},
}

func (v *virus) turnRight() {
	v.d = rightMap[v.d]
}

var leftMap map[[2]int][2]int = map[[2]int][2]int{
	[2]int{0, 1}:  [2]int{-1, 0},
	[2]int{1, 0}:  [2]int{0, 1},
	[2]int{0, -1}: [2]int{1, 0},
	[2]int{-1, 0}: [2]int{0, -1},
}

func (v *virus) turnLeft() {
	v.d = leftMap[v.d]
}

func (v *virus) move() {
	v.p[0] += v.d[0]
	v.p[1] += v.d[1]
}

func parse(lines []string) map[position]state {
	m := make(map[position]state)
	for i := range lines {
		for j := range lines {
			if lines[i][j] == '#' {
				m[[2]int{i, j}] = Infected
			}
		}
	}
	return m
}

func getInput() []string {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(bytes), "\n")
	return strs
}
