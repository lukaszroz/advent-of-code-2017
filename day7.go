package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	label       string
	children    []*node
	parent      *node
	weight      int
	totalWeight int
}

func main() {
	var in map[string]*node = getInput()
	fmt.Println("Star 1: ", getRoot(in).label)

	var fixedWeight int
	for n := getRoot(in); n != nil && len(n.children) > 2; {
		w := n.children[0].totalWeight
		w2 := n.children[1].totalWeight
		if n.children[2].totalWeight == w2 {
			w = w2
		}
		var next *node
		for _, n := range n.children {
			if n.totalWeight != w {
				next = n
				fixedWeight = n.weight - (n.totalWeight - w)
			}
		}
		n = next
	}
	fmt.Println("Star 2: ", fixedWeight)

}

func getRoot(nodes map[string]*node) *node {
	for _, n := range nodes {
		for n.parent != nil {
			n = n.parent
		}
		return n
	}
	panic("empty map!")
}

func getInput() map[string]*node {
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	in := make(map[string]*node)
	children := make(map[*node][]string)

	reLabels := regexp.MustCompile("[a-z]+")
	reWeight := regexp.MustCompile("[0-9]+")

	for _, line := range lines {
		labels := reLabels.FindAllString(line, -1)
		weight, err := strconv.Atoi(reWeight.FindString(line))
		if err != nil {
			panic(err)
		}
		n := &node{label: labels[0], weight: weight, totalWeight: weight}
		in[labels[0]] = n
		children[n] = labels[1:]
	}

	//assign children
	for n, list := range children {
		n.children = make([]*node, len(list))
		for i, child := range list {
			n.children[i] = in[child]
		}
	}

	//assign parents
	for _, n := range in {
		for _, child := range n.children {
			child.parent = n
		}
	}

	//calculate total weights
	for _, n := range in {
		for k := n.parent; k != nil; k = k.parent {
			k.totalWeight += n.weight
		}
	}
	return in
}
