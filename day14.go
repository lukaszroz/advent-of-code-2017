package main

import (
	"fmt"
	"math/bits"
	"os"
	"sync"
	"sync/atomic"
)

func main() {
	input := os.Args[1]
	grid := make([][]bool, 128)
	var usedSum int32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 128; i++ {
		grid[i] = make([]bool, 128)
		wg.Add(1)
		go processRow(grid[i], &usedSum, fmt.Sprintf("%s-%d", input, i), &wg)
	}

	wg.Wait()
	fmt.Println("Star 1: ", usedSum)
	fmt.Println("Star 2: ", countRegions(grid))
}

func processRow(row []bool, onesCount *int32, hashInput string, wg *sync.WaitGroup) {
	defer wg.Done()
	j := 0
	for _, b := range knotHash(hashInput) {
		atomic.AddInt32(onesCount, int32(bits.OnesCount8(b)))
		for _, bit := range fmt.Sprintf("%08b", b) {
			if bit == '1' {
				row[j] = true
			}
			j++
		}
	}
}

func countRegions(grid [][]bool) int {
	count := 0
	for i, row := range grid {
		for j, used := range row {
			if used {
				visit(i, j, grid)
				count++
			}
		}
	}
	return count
}

func visit(i, j int, grid [][]bool) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || !grid[i][j] {
		return
	}
	grid[i][j] = false
	visit(i+1, j, grid)
	visit(i-1, j, grid)
	visit(i, j+1, grid)
	visit(i, j-1, grid)
}

func knotHash(s string) []byte {
	bytes := []byte(s)
	bytes = append(bytes, 17, 31, 73, 47, 23)

	sparseHash := make([]byte, 256)
	for i := range sparseHash {
		sparseHash[i] = byte(i)
	}
	for start, skip := 0, 0; skip < len(bytes)*64; skip++ {
		length := int(bytes[skip%len(bytes)])
		reverse(sparseHash, start, length-1)
		start += length + skip
		start %= len(sparseHash)
	}

	denseHash := make([]byte, 16)
	for idx := range denseHash {
		denseHash[idx] = sparseHash[idx*16]
		for i := 1; i < 16; i++ {
			denseHash[idx] ^= sparseHash[idx*16+i]
		}
	}
	return denseHash
}

func reverse(hash []byte, start, length int) {
	for i := 0; i <= length/2; i++ {
		j := (start + i) % len(hash)
		k := (start + length - i) % len(hash)
		hash[j], hash[k] = hash[k], hash[j]
	}
}
