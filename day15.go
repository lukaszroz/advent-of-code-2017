package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	startA := parseUint(os.Args[1])
	startB := parseUint(os.Args[2])

	//Closures are faster for both stars on my system
	fmt.Println("Star 1: ", count(startA, startB, 1, 1, 40000000, false))
	fmt.Println("Star 2: ", count(startA, startB, 4, 8, 5000000, false))
}

func count(startA, startB, filterA, filterB int, times int, channel bool) int {
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
		if genA() == genB() {
			sum++
		}
	}
	return sum
}

func next(n, factor, filter int) int {
	n *= factor
	n %= 2147483647
	for n&(filter-1) != 0 {
		n *= factor
		n %= 2147483647
	}
	return n
}

type generator func() int
type newGenerator func(int, int, int) generator

func newGeneratorClosure(n, factor, filter int) generator {
	return func() int {
		n = next(n, factor, filter)
		return n & 0xffff
	}
}

func newGeneratorChannel(start, factor, filter int) generator {
	c := make(chan int, 64)
	go generatorRoutine(start, factor, filter, c)
	return func() int {
		return <-c
	}
}

func generatorRoutine(n, factor, filter int, c chan int) {
	for {
		n = next(n, factor, filter)
		c <- n & 0xffff
	}
}

func parseUint(s string) (u int) {
	u, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}
