package mathbeta

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	bits := make([]byte, 1024)
	seeds := []int{5, 7, 11, 13, 31, 37, 61}
	cap := 8192
	hashes := make([]func(string) int, 7)
	for i := 0; i < len(seeds); i++ {
		hashes[i] = SimpleHash(cap, seeds[i])
	}

	bf := NewBloomFilter(bits, hashes)
	bf.Add("abc", "def", "ghi")
	c := bf.Contains("ghi")
	fmt.Println(c)
	if !c {
		t.Error("false negative, 'ghi' has already been added")
	}
	c = bf.Contains("jkl")
	if c {
		t.Error("false positive, 'jkl' has not been added")
	}
}
