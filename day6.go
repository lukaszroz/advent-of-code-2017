package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := getInput()

	checksums := make(map[uint64]int)
	i := 1
	checksums[checksum(in)] = i
	i++
	maxIdx := findMaxIdx(in)

	for {
		n := in[maxIdx]
		in[maxIdx] = 0
		for idx := maxIdx + 1; n > 0; idx, n = idx+1, n-1 {
			in[idx%len(in)]++
		}

		sum := checksum(in)
		if checksums[sum] > 0 {
			fmt.Println("1st star: ", len(checksums))
			fmt.Println("2nd star: ", i-checksums[sum])
			return
		} else {
			checksums[sum] = i
			i++
		}

		maxIdx = findMaxIdx(in)
	}
}

func checksum(list []int) uint64 {
	var sum, base uint64
	sum = 0
	base = 17
	for _, i := range list {
		sum *= base
		sum += uint64(i)
	}
	return sum
}

func findMaxIdx(list []int) int {
	maxIdx := 0
	for idx, item := range list {
		if item > list[maxIdx] {
			maxIdx = idx
		}
	}
	return maxIdx
}

func getInput() []int {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\t")
	in := make([]int, len(lines))
	for i, ns := range lines {
		in[i], err = strconv.Atoi(ns)
		if err != nil {
			panic(err)
		}
	}
	return in
}
