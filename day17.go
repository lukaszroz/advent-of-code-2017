package main

import (
	"fmt"
	"os"
)

func main() {
	var step int
	fmt.Sscanf(os.Args[1], "%d", &step)
	root := &node{}
	root.next = root
	root.prev = root
	current := root

	for i := 1; i < 2018; i++ {
		for j := 0; j < step; j++ {
			current = current.next
		}
		n := &node{}
		n.val = i
		n.next = current.next
		n.prev = current

		current.next.prev = n
		current.next = n
		current = n
	}

	fmt.Println("Star 1: ", current.next.val)

	length := 2
	position := 1
	after_zero := 1
	for i := 2; i <= 50000000; i++ {
		position = (position + step) % length
		if position == 0 {
			after_zero = i
		}
		length++
		position++
	}

	fmt.Println("Star 2: ", after_zero)
}

type node struct {
	next, prev *node
	val        int
}

func print(root *node) {
	fmt.Print(root.val, " ")
	for n := root.next; n != root; n = n.next {
		fmt.Print(n.val, " ")
	}
	fmt.Println("\n")
}
