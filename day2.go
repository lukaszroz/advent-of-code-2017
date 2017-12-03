package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	sum := 0
	for _, line := range lines {
		nums := strings.Split(string(line), "\t")
		min, _ := strconv.Atoi(nums[0])
		max := min
		for _, ns := range nums {
			n, _ := strconv.Atoi(ns)
			max = maxF(max, n)
			min = minF(min, n)
		}
		sum += max - min
	}

	fmt.Println(sum)
}

func maxF(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minF(a, b int) int {
	if a < b {
		return a
	}
	return b
}
