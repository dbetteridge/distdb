package main

import "testing"

func TestInsertRoot(t *testing.T) {
	tree := RBTree{}
	tree.insert("A", "1")
	if tree.Root.Value != "1" {
		t.Errorf("Insert on empty tree failed")
	}
}

func TestInsertSmaller(t *testing.T) {
	tree := RBTree{}
	tree.insert("K", "11")
	tree.insert("B", "2")
	if tree.Root.Value != "11" {
		t.Errorf("Insert on empty tree failed")
	}
	if tree.Root.LeftChild.Value != "2" {
		t.Error("Smaller element not placed correctly")
	}
}

func TestBlackBalance(t *testing.T) {
	// If balanced all routes from root to leaf should have the same +/-1 number of black nodes
	tree := RBTree{}

	tree.insert("K", "11")
	tree.insert("B", "2")
	tree.insert("A", "1")
	tree.insert("G", "7")
	tree.insert("E", "5")
	tree.insert("H", "8")

	tree.insert("N", "14")
	tree.insert("O", "15")

	tree.insert("D", "4")

	countA := 0
	countB := 0
	root := tree.Root
	for {
		if root.Black {
			countA++
		}
		if root.LeftChild != nil {
			root = root.LeftChild

		} else {
			break
		}
	}

	for {
		if root.Black {
			countB++
		}
		if root.RightChild != nil {
			root = root.RightChild

		} else {
			break
		}
	}

	if !((countA-countB) == 1 || (countA-countB) == -1) {
		t.Error("Not balanced", countA, countB)
	}

}
