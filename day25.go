package main

import (
	"fmt"
)

type state byte

const (
	A state = iota + 1
	B
	C
	D
	E
	F
)

type turing struct {
	current  state
	tape     map[int]bool
	i        int
	checksum uint
}

func main() {
	var t turing
	t.tape = make(map[int]bool)
	t.current = A
	for i := 0; i < 12794428; i++ {
		t.step()
	}
	fmt.Println("Star 1: ", t.checksum)
}

func (t *turing) step() {
	switch t.current {
	case A:
		if t.tape[t.i] {
			t.tape[t.i] = false
			t.checksum--
			t.i--
			t.current = F
		} else {
			t.tape[t.i] = true
			t.checksum++
			t.i++
			t.current = B
		}
	case B:
		if t.tape[t.i] {
			t.tape[t.i] = false
			t.checksum--
			t.i++
			t.current = D
		} else {
			t.i++
			t.current = C
		}
	case C:
		if t.tape[t.i] {
			t.i++
			t.current = E
		} else {
			t.tape[t.i] = true
			t.checksum++
			t.i--
			t.current = D
		}
	case D:
		if t.tape[t.i] {
			t.tape[t.i] = false
			t.checksum--
			t.i--
			t.current = D
		} else {
			t.i--
			t.current = E
		}
	case E:
		if t.tape[t.i] {
			t.i++
			t.current = C
		} else {
			t.i++
			t.current = A
		}
	case F:
		if t.tape[t.i] {
			t.i++
			t.current = F
		} else {
			t.tape[t.i] = true
			t.checksum++
			t.i--
			t.current = A
		}
	}
}
