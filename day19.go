package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var steps int = 1

func main() {
	diagram := getInput()
	path := make([]byte, 10)
	goSouth(1, strings.IndexRune(diagram[0], '|'), diagram, &path)
	fmt.Println("Star 1:", string(path))
	fmt.Println("Star 2:", steps)
}

func goCommon(i, j *int, id, jd, is, js int, diagram []string, path *[]byte) {
	var done bool
	for ; *i >= 0 && *i < len(diagram) && *j >= 0 && *j < len(diagram[*i]) && !done; *i, *j = *i+id, *j+jd {
		steps++
		b := diagram[*i][*j]
		switch b {
		case ' ':
			*i += is
			*j += js
			done = true
		case '|':
		case '+':
		case '-':
		default:
			*path = append(*path, b)
		}
	}
	steps--
}

func goSouth(i, j int, diagram []string, path *[]byte) {
	goCommon(&i, &j, 1, 0, -2, 0, diagram, path)
	goWestEast(i, j, diagram, path)
}

func goNorth(i, j int, diagram []string, path *[]byte) {
	goCommon(&i, &j, -1, 0, 2, 0, diagram, path)
	goWestEast(i, j, diagram, path)
}
func goWest(i, j int, diagram []string, path *[]byte) {
	goCommon(&i, &j, 0, 1, 0, -2, diagram, path)
	goSouthNorth(i, j, diagram, path)
}
func goEast(i, j int, diagram []string, path *[]byte) {
	goCommon(&i, &j, 0, -1, 0, 2, diagram, path)
	goSouthNorth(i, j, diagram, path)
}

func goWestEast(i, j int, diagram []string, path *[]byte) {
	if j+1 < len(diagram[i]) && diagram[i][j+1] != ' ' {
		goWest(i, j+1, diagram, path)
	} else if j > 0 && diagram[i][j-1] != ' ' {
		goEast(i, j-1, diagram, path)
	}
}

func goSouthNorth(i, j int, diagram []string, path *[]byte) {
	if i+1 < len(diagram) && diagram[i+1][j] != ' ' {
		goSouth(i+1, j, diagram, path)
	} else if i > 0 && diagram[i-1][j] != ' ' {
		goNorth(i-1, j, diagram, path)
	}
}

func getInput() []string {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(bytes), "\n")
	return strs
}
