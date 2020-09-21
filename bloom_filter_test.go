package main

import (
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	bf := new(BloomFilter)
	bf.newFilter(100, 10)
	if len(bf.bitArray) != 100 {
		t.Errorf("Initialization of filter failed")
	}
}

func TestAdd(t *testing.T) {
	bf := new(BloomFilter)
	bf.newFilter(100, 10)
	bf.add("First")
	test := make([]byte, 100)
	match := true
	for i := 0; i < len(bf.bitArray); i++ {
		if test[i] != bf.bitArray[i] {
			match = false
		}
	}
	if match {
		t.Errorf("Initialization of filter failed")
	}
}

func TestContains(t *testing.T) {
	bf := new(BloomFilter)
	bf.newFilter(100, 10)
	bf.add("First")

	contains := bf.contains("First")
	if !contains {
		t.Errorf("")
	}
}

func TestFilter(t *testing.T) {
	bf := new(BloomFilter)
	bf.newFilter(100, 10)

	animals := []string{"dog", "cat", "giraffe", "fly", "mosquito", "horse", "eagle", "bird", "bison", "boar", "butterfly", "ant", "anaconda", "bear", "chicken", "dolphin", "donkey", "crow", "crocodile"}
	for i := range animals {
		bf.add(animals[i])
	}

	for i := range animals {
		if !bf.contains(animals[i]) {
			t.Errorf("False Negative %s", animals[i])
		}
	}
}

func TestFalsePositive(t *testing.T) {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()
	bf := new(BloomFilter)
	bf.newFilter(100, 10)

	animals := []string{"dog", "cat", "giraffe", "fly", "mosquito", "horse", "eagle", "bird", "bison", "boar", "butterfly", "ant", "anaconda", "bear", "chicken", "dolphin", "donkey", "crow", "crocodile"}
	for i := range animals {
		bf.add(animals[i])
	}

	otherAnimals := []string{"badger", "cow", "pig", "sheep", "bee", "wolf", "fox",
		"whale", "shark", "fish", "turkey", "duck", "dove",
		"deer", "elephant", "frog", "falcon", "goat", "gorilla",
		"hawk"}

	noFalsePositives := true
	for i := range otherAnimals {
		if bf.contains(otherAnimals[i]) {
			noFalsePositives = false
		}
	}

	if noFalsePositives {
		t.Errorf("No false positives found")
	}

}

func TestAddComb(t *testing.T) {

	bf := new(BloomFilter)
	bf.newFilter(100, 10)
	bf.addComb("First")
	test := make([]byte, 100)
	match := true
	for i := 0; i < len(bf.bitArray); i++ {
		if test[i] != bf.bitArray[i] {
			match = false
		}
	}
	if match {
		t.Errorf("Initialization of filter failed")
	}
}

func TestContainsComb(t *testing.T) {
	bf := new(BloomFilter)
	bf.newFilter(100, 10)
	bf.addComb("First")

	contains := bf.containsComb("First")
	if !contains {
		t.Errorf("")
	}
}

func TestFalsePositiveComb(t *testing.T) {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()
	bf := new(BloomFilter)
	bf.newFilter(100, 10)

	animals := []string{"dog", "cat", "giraffe", "fly", "mosquito", "horse", "eagle", "bird", "bison", "boar", "butterfly", "ant", "anaconda", "bear", "chicken", "dolphin", "donkey", "crow", "crocodile"}
	for i := range animals {
		bf.addComb(animals[i])
	}

	otherAnimals := []string{"badger", "cow", "pig", "sheep", "bee", "wolf", "fox",
		"whale", "shark", "fish", "turkey", "duck", "dove",
		"deer", "elephant", "frog", "falcon", "goat", "gorilla",
		"hawk"}

	noFalsePositives := true
	for i := range otherAnimals {
		if bf.containsComb(otherAnimals[i]) {
			noFalsePositives = false
		}
	}

	if noFalsePositives {
		t.Errorf("No false positives found")
	}

}
