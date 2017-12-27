package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func main() {
	program := getInput()
	fmt.Println("Star 1: ", debug(program))
	fmt.Println("Star 2: ", optimized())
}

func debug(program []string) int {
	registers := make(map[byte]int)
	registers['1'] = 1
	count := 0
	for i := 0; i > -1 && i < len(program); i++ {
		line := program[i]
		cmd := line[0:3]
		r := line[4]
		switch cmd {
		case "set":
			registers[r] = getY(line, registers)
		case "sub":
			registers[r] -= getY(line, registers)
		case "mul":
			registers[r] *= getY(line, registers)
			count++
		case "jnz":
			if registers[r] != 0 {
				i += int(getY(line, registers) - 1)
			}
		default:
			panic("Oh no!")
		}
	}
	return count
}

func optimized() int {
	composite_count := 0
	b := 79*100 + 100000

	for i := b; i <= b+17000; i += 17 {
		if isComposite(i) {
			composite_count++
		}
	}
	return composite_count
}

func isComposite(n int) bool {
	if n%2 == 0 {
		return true
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return true
		}
	}
	return false
}

func getY(line string, rr map[byte]int) int {
	b := line[6]
	if b >= 'a' && b <= 'z' {
		return rr[b]
	} else {
		var i int
		fmt.Sscanf(line[5:], "%d", &i)
		return i
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
