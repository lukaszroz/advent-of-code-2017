package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes := getInput()
	hash := make([]byte, 256)
	for i := range hash {
		hash[i] = byte(i)
	}

	start := 0
	for skip, length := range star1_lengths(bytes) {
		reverse(hash, start, length-1)
		start += length + skip
		start %= len(hash)
	}
	fmt.Println("Star 1: ", int(hash[0])*int(hash[1]))

	bytes = append(bytes, 17, 31, 73, 47, 23)
	for i := range hash {
		hash[i] = byte(i)
	}
	for start, skip := 0, 0; skip < len(bytes)*64; skip++ {
		length := int(bytes[skip%len(bytes)])
		reverse(hash, start, length-1)
		start += length + skip
		start %= len(hash)
	}
	dense := make([]byte, 16)
	for idx := range dense {
		dense[idx] = hash[idx*16]
		for i := 1; i < 16; i++ {
			dense[idx] ^= hash[idx*16+i]
		}
	}
	fmt.Println("Star 2: ", hex.EncodeToString(dense))
}

func reverse(hash []byte, start, length int) {
	for i := 0; i <= length/2; i++ {
		j := (start + i) % len(hash)
		k := (start + length - i) % len(hash)
		hash[j], hash[k] = hash[k], hash[j]
	}
}

func getInput() []byte {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	return bytes
}

func star1_lengths(bytes []byte) []int {
	strs := strings.Split(string(bytes), ",")
	ints := make([]int, len(strs))
	for idx, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints[idx] = n
	}
	return ints
}
