package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	startA := parseUint(os.Args[1])
	startB := parseUint(os.Args[2])

	fmt.Println("Star 1: ", count(startA, startB, 1, 1, 40000000, false))
	fmt.Println("Star 2: ", count(startA, startB, 4, 8, 5000000, false))
}

func count(startA, startB, filterA, filterB uint64, times int, channel bool) int {
	var factory newGenerator
	if channel {
		factory = newGeneratorChannel
	} else {
		factory = newGeneratorClosure
	}
	genA := factory(startA, 16807, filterA)
	genB := factory(startB, 48271, filterB)
	sum := 0
	for i := 0; i < times; i++ {
		if (genA() % 65536) == (genB() % 65536) {
			sum++
		}
	}
	return sum
}

func next(n, factor, filter uint64) uint64 {
	n *= factor
	n %= 2147483647
	for n%filter != 0 {
		n *= factor
		n %= 2147483647
	}
	return n
}

type generator func() uint64
type newGenerator func(uint64, uint64, uint64) generator

func newGeneratorClosure(n, factor, filter uint64) generator {
	return func() uint64 {
		n = next(n, factor, filter)
		return n
	}
}

func newGeneratorChannel(start, factor, filter uint64) generator {
	c := make(chan uint64, 64)
	go generatorRoutine(start, factor, filter, c)
	return func() uint64 {
		return <-c
	}
}

func generatorRoutine(n, factor, filter uint64, c chan uint64) {
	for {
		n = next(n, factor, filter)
		c <- n
	}
}

func parseUint(s string) (u uint64) {
	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return
}
