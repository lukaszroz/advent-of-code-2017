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

	fmt.Println("1st star: ", sum)

	sum = 0
	for _, line := range lines {
		nums := strings.Split(string(line), "\t")
		fmt.Println(nums)
		
		sum += noCoprime(nums)
	}
	fmt.Println("2nd star: ", sum)
}

func noCoprime(nums []string) int {
	for i, nis := range nums {
		for j, njs := range nums {
			if i != j {
				ni, _ := strconv.Atoi(nis)
				nj, _ := strconv.Atoi(njs)
				if(ni % nj == 0) {
					return ni / nj
				} else if(nj % ni == 0) {
					return nj /ni
				}
			}
		}
	}
	panic(fmt.Sprintf("All numbers in a row coprime! %v", nums))
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
