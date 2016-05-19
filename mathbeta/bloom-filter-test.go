package main

import (
	"fmt"
	"mathbeta"
)

func main() {
	bits := make([]byte, 1024)
	seeds := []int{5, 7, 11, 13, 31, 37, 61}
	cap := 8192
	hashes := make([]func(string) int, 7)
	for i := 0; i < len(seeds); i++ {
		hashes[i] = mathbeta.SimpleHash(cap, seeds[i])
	}

	bf := mathbeta.NewBloomFilter(bits, hashes)
	bf.Add("abc", "def", "ghi")
	fmt.Println(bf.Contains("ghi"))
}
