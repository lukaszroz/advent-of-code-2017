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

	in := make([]int, len(lines))
	for i, ns := range lines {
		in[i], err = strconv.Atoi(ns)
		if err != nil {
			panic(err)
		}
	}

	
	fmt.Println("1st star: ", star(in, false))
	fmt.Println("2nd star: ", star(in, true))
}

func star(in []int, second bool) int {
	list := make([]int, len(in))
	copy(list, in)
	count := 0
	for j := 0; j > -1 && j < len(list); count++ {
		k := list[j]
		if second && k > 2 {
			list[j]--
		} else {			
			list[j]++
		}
		j += k
	}
	return count
}