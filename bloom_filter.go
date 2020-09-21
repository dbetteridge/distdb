package main

import (
	"encoding/binary"
	"fmt"

	"github.com/roberson-io/mmh3"
)

// https://github.com/Claudenw/BloomFilter/wiki/Bloom-Filters----An-overview

// BloomFilter : Filter using a number of hashes
type BloomFilter struct {
	numHashes int
	size      int
	bitArray  []byte
}

// New : Create a new BloomFilter
func (b *BloomFilter) newFilter(size int, numHashes int) {
	arr := make([]byte, size)
	b.numHashes = numHashes
	b.size = size
	b.bitArray = arr
}

func (b *BloomFilter) add(item string) {
	for i := 0; i <= b.numHashes; i++ {
		hash := mmh3.Hashx86_128([]byte(item), uint32(i))
		index := binary.LittleEndian.Uint32(hash) % uint32(b.size)
		b.bitArray[index] = 1
	}
}

func (b *BloomFilter) contains(item string) bool {
	out := true
	for i := 0; i <= b.numHashes; i++ {
		hash := mmh3.Hashx86_128([]byte(item), uint32(i))
		index := binary.LittleEndian.Uint32(hash) % uint32(b.size)
		if b.bitArray[index] == 0 {
			out = false
		}
	}
	return out
}

// addComb : Generate a single hash and split in two, use the left half as a base and add hash * numHashes to get the index for each hash
// so for hash 0 = h1 + (10 * h2)
func (b *BloomFilter) addComb(item string) {
	hashLong := binary.LittleEndian.Uint64(mmh3.Hashx64_128([]byte(item), uint32(0)))
	h1 := uint32(hashLong >> 32)                // First half of the hash
	h2 := uint32(hashLong & 0xffffffffffffffff) // Second half of the hash
	hash := h1
	for i := 0; i <= b.numHashes; i++ {
		add := uint32(b.numHashes) * h2
		hash += add
		index := hash % uint32(b.size) // Confine result to between 0 and b.size for indexing the bitarray
		b.bitArray[index] = 1
	}
}

func (b *BloomFilter) containsComb(item string) bool {
	out := true
	hashLong := binary.LittleEndian.Uint64(mmh3.Hashx64_128([]byte(item), uint32(0)))
	h1 := uint32(hashLong >> 32)
	h2 := uint32(hashLong & 0xffffffffffffffff)
	fmt.Println(hashLong, h1, h2)
	hash := h1
	for i := 0; i <= b.numHashes; i++ {
		add := uint32(b.numHashes) * h2
		hash += add
		index := hash % uint32(b.size)
		if b.bitArray[index] == 0 {
			out = false
		}
	}
	return out
}
