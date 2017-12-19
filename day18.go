package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	program := getInput()
	fmt.Println("Star 1: ", makeSomeNoise(program))

	p0, p1 := make(chan int, 100), make(chan int, 100)
	ret0, ret1 := make(chan int), make(chan int)
	go duet(program, p0, p1, ret0, 0)
	go duet(program, p1, p0, ret1, 1)
	fmt.Println("Start 2: ", <-ret1)
}

func makeSomeNoise(program []string) int {
	registers := make(map[byte]int)
	registers['1'] = 1
	sound := int(0)
	for i := 0; i > -1 && i < len(program); i++ {
		line := program[i]
		cmd := line[0:3]
		r := line[4]
		switch cmd {
		case "snd":
			sound = registers[r]
		case "rcv":
			if registers[r] != 0 {
				return sound
			}
		default:
			common(line, &i, registers)
		}
	}
	panic("Oh no!")
}

func duet(program []string, this, other, ret chan int, id int) {
	registers := make(map[byte]int)
	registers['p'] = int(id)
	registers['1'] = 1
	counter := 0
	for i := 0; i > -1 && i < len(program); i++ {
		line := program[i]
		cmd := line[0:3]
		r := line[4]
		switch cmd {
		case "snd":
			counter++
			other <- registers[r]
		case "rcv":
			timeout := make(chan bool, 1)
			go func() {
				time.Sleep(time.Millisecond * 10)
				timeout <- true
			}()
			select {
			case registers[r] = <-this:
			case <-timeout:
				ret <- counter
				return
			}
		default:
			common(line, &i, registers)
		}
	}
}

func common(line string, current *int, registers map[byte]int) {
	cmd := line[0:3]
	r := line[4]
	switch cmd {
	case "set":
		registers[r] = getY(line, registers)
	case "add":
		registers[r] += getY(line, registers)
	case "mul":
		registers[r] *= getY(line, registers)
	case "mod":
		registers[r] = registers[r] % getY(line, registers)
	case "jgz":
		if registers[r] > 0 {
			*current += int(getY(line, registers) - 1)
		}
	}
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
