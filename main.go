package main

import (
	"fmt"
)

func main() {
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

	tree.delete("K")

	tree.drawTree()
	fmt.Println("SIZE", tree.Size)

	w := WAL{}
	w.new()

	w.write("Hello")
	// fmt.Println(w.flush())
}
