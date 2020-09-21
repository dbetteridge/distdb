package main

import (
	"fmt"
)

/**
* Leaf has no children and is black
* Red node must have black children
* Every simple path from a node to a descendant leaf contains the same number of black nodes
 */

// Node : Single Node element in the Tree
type Node struct {
	Parent     *Node
	LeftChild  *Node
	RightChild *Node
	Key        string
	Value      string
	Black      bool
}

// RBTree : A Red-Black tree
type RBTree struct {
	Root *Node
	Size int
}

func (t *RBTree) drawTree() {
	root := t.Root
	level := 0
	var nodes []*Node
	nodes = append(nodes, root)
	for {
		if len(nodes) == 0 {
			break
		}
		var nodesTemp []*Node
		for _, node := range nodes {
			if node.Parent != nil {
				fmt.Print("L", level, "P", node.Parent.Value, "N", node.Value, " ")
			} else {
				fmt.Print("L", level, "N", node.Value, " ")
			}
			if node.LeftChild != nil {
				nodesTemp = append(nodesTemp, node.LeftChild)
			}
			if node.RightChild != nil {
				nodesTemp = append(nodesTemp, node.RightChild)
			}
		}
		fmt.Println("")
		nodes = nodesTemp
		level++
	}
}

func (t *RBTree) findParent(key string, root *Node) (parent *Node, right bool) {

	if key < root.Key {
		if root.LeftChild != nil {
			left := root.LeftChild
			return t.findParent(key, left)
		}
		// Insert as LeftChild of current root
		return root, false
	}

	if key > root.Key {
		if root.RightChild != nil {
			// fmt.Println(key, root.Key, root.RightChild.Key)
			right := root.RightChild
			return t.findParent(key, right)
		}
		// Insert as RightChild of current root
		return root, true
	}

	// current root == key
	if root.Parent.LeftChild.Key == key {
		return root.Parent, false
	}
	return root.Parent, true
}

func (t *RBTree) deleteFix(node *Node) {

	for {
		isRoot := node.Parent == nil
		isBlack := node.Black

		if isBlack && isRoot {
			break
		}

		if node.Parent != nil && node.Parent.LeftChild != nil && node.Parent.LeftChild.Key == node.Key {
			// node is left child of its parent
			sibling := node.Parent.RightChild
			if !sibling.Black {
				sibling.Black = true
				node.Parent.Black = false

				t.rotateLeft(node.Parent)

				sibling = node.RightChild
			}
			if sibling.RightChild.Black && sibling.LeftChild.Black {
				sibling.Black = false
				node = node.Parent
			} else {
				if sibling.RightChild.Black {
					sibling.LeftChild.Black = true
					sibling.Black = false
					t.rotateRight(sibling)

					sibling = node.Parent.RightChild
				}

				sibling.Black = node.Parent.Black
				node.Parent.Black = true
				sibling.RightChild.Black = true
				t.rotateLeft(node.Parent)
				node = t.Root
			}

		} else {
			// node is right child of its parent
			var sibling *Node
			if node.Parent != nil {
				sibling = node.Parent.LeftChild
			}

			if sibling != nil && !sibling.Black {
				sibling.Black = true
				node.Parent.Black = false

				t.rotateRight(node.Parent)

				sibling = node.LeftChild
			}
			if sibling != nil && sibling.LeftChild != nil && sibling.RightChild != nil && sibling.LeftChild.Black && sibling.RightChild.Black {
				sibling.Black = false
				node = node.Parent
			} else if sibling != nil && sibling.LeftChild != nil && sibling.LeftChild.Black {
				sibling.RightChild.Black = true
				sibling.Black = false
				t.rotateLeft(sibling)

				sibling = node.Parent.LeftChild
			} else {
				if sibling != nil {
					sibling.Black = node.Parent.Black
				}
				if node.Parent != nil {
					node.Parent.Black = true
				}
				if sibling != nil && sibling.LeftChild != nil {
					sibling.LeftChild.Black = true
				}
				if node.Parent != nil {
					t.rotateRight(node.Parent)
				}
				node = t.Root
			}
		}
	}
	node.Black = true
}

func (t *RBTree) min(node *Node) *Node {
	var subTree *Node
	subTree = node
	for {
		if subTree.LeftChild != nil {
			subTree = subTree.LeftChild
		} else {
			break
		}
	}

	return subTree
}

func (t *RBTree) rbTransplant(a *Node, b *Node) {
	if a.Parent == nil {
		t.Root = b
	} else if a.Key == a.Parent.LeftChild.Key {
		a.Parent.LeftChild = b
	} else {
		a.Parent.RightChild = b
	}

	b.Parent = a.Parent
}

func (t *RBTree) delete(key string) {
	root := t.Root
	parent, right := t.findParent(key, root)

	var nodeToDelete *Node
	var originalColor bool
	var tempNode *Node

	if right {
		nodeToDelete = parent.RightChild
	} else {
		nodeToDelete = parent.LeftChild
	}
	originalColor = nodeToDelete.Black

	if nodeToDelete.LeftChild == nil {
		tempNode = nodeToDelete.RightChild
		// Assign nodes right child to replace the deleted node
		tempNode.Parent = parent
		if right {
			parent.RightChild = tempNode
		} else {
			parent.LeftChild = tempNode
		}
	}

	if nodeToDelete.RightChild == nil {
		tempNode = nodeToDelete.LeftChild
		// Assign nodes left child to replace the deleted node
		tempNode.Parent = parent
		if right {
			parent.RightChild = tempNode
		} else {
			parent.LeftChild = tempNode
		}
	}

	if nodeToDelete.LeftChild != nil && nodeToDelete.RightChild != nil {
		// Deleted node has two children
		minST := t.min(nodeToDelete.RightChild)

		originalColor = minST.Black
		tempNode = minST.RightChild

		if minST.Parent.Key == nodeToDelete.Key {
			tempNode.Parent = minST
		} else {
			t.rbTransplant(minST, minST.RightChild)
			minST.RightChild = nodeToDelete.RightChild
			minST.RightChild.Parent = minST
		}

		t.rbTransplant(nodeToDelete, minST)
		minST.LeftChild = nodeToDelete.LeftChild
		minST.LeftChild.Parent = minST
		minST.Black = nodeToDelete.Black
	}

	if originalColor {
		t.deleteFix(tempNode)
	}
}

func (t *RBTree) rotateLeft(node *Node) {

	y := node.RightChild
	node.RightChild = y.LeftChild

	if y.LeftChild != nil && y.LeftChild.Key != t.Root.Key {
		y.LeftChild.Parent = node
	}

	y.Parent = node.Parent

	if node.Parent == nil {
		t.Root = y
	} else if node.Parent.LeftChild != nil && node.Key == node.Parent.LeftChild.Key {
		node.Parent.LeftChild = y
	} else {
		node.Parent.RightChild = y
	}

	y.LeftChild = node
	node.Parent = y

}

func (t *RBTree) rotateRight(node *Node) {
	if node == nil {
		return
	}
	y := node.LeftChild
	if y != nil {
		node.LeftChild = y.RightChild

		if y.RightChild != nil && y.RightChild.Key != t.Root.Key {
			y.RightChild.Parent = node
		}

		y.Parent = node.Parent

		if node.Parent == nil {
			t.Root = y
		} else if node.Key == node.Parent.RightChild.Key {
			node.Parent.RightChild = y
		} else {
			node.Parent.LeftChild = y
		}

		y.RightChild = node
		node.Parent = y
	}

}

func (t *RBTree) reorder(target *Node) {
	for {

		p := target.Parent

		if p.Black {
			break
		}
		var gP *Node
		if p != nil {
			gP = p.Parent
		}

		if gP != nil && gP.LeftChild != nil && p.Key == gP.LeftChild.Key {
			uncle := gP.RightChild
			if uncle != nil && !uncle.Black {
				gP.LeftChild.Black = true
				gP.RightChild.Black = true
				gP.Black = false
				target = gP
			} else {
				if p.RightChild != nil && target.Key == p.RightChild.Key {
					target = p
					t.rotateLeft(target)
				}
				p.Black = true
				gP.Black = false
				t.rotateRight(gP)
			}
		} else {
			if gP != nil && gP.LeftChild != nil && !gP.LeftChild.Black {
				gP.LeftChild.Black = true
				gP.RightChild.Black = true
				gP.Black = false
				target = gP
			} else {
				if p.LeftChild != nil && target.Key == p.LeftChild.Key {
					target = p
					t.rotateRight(target)
				}
				p.Black = true
				if gP != nil {
					gP.Black = false
					t.rotateLeft(gP)
				}
			}
		}

		if target.Key == t.Root.Key {
			break
		}

	}
	t.Root.Black = true
}

func (t *RBTree) insert(key string, value string) {
	if t.Root == nil {
		t.Root = &Node{Parent: nil, LeftChild: nil, RightChild: nil, Key: key, Value: value, Black: true}
		t.Size = len(value)
		return
	}
	root := t.Root

	parent, right := t.findParent(key, root)

	var target *Node
	if right {
		parent.RightChild = &Node{Parent: parent, LeftChild: nil, RightChild: nil, Key: key, Value: value, Black: false}
		// Used for balancing
		target = parent.RightChild
	} else {
		parent.LeftChild = &Node{Parent: parent, LeftChild: nil, RightChild: nil, Key: key, Value: value, Black: false}
		target = parent.LeftChild
	}

	t.Size += len(value)

	// Need to walk the tree, update node colours and rotate
	t.reorder(target)
}

func (t *RBTree) get(key string) (string, error) {
	root := t.Root
	for {
		if key < root.Key {
			if root.LeftChild != nil {
				root = root.LeftChild
			} else {
				return "", fmt.Errorf(`Key: %s not found`, key)
			}
		}

		if key > root.Key {
			if root.RightChild == nil {
				return "", fmt.Errorf(`Key: %s not found`, key)
			}
			root = root.RightChild
		}

		if key == root.Key {
			return root.Value, nil
		}
	}
}

// Row :
type Row struct {
	Key   string
	Value string
}

// func (t *RBTree) flush(root *Node, rows []Row) []Row {
// 	for {
// 		if root.LeftChild != nil {
// 			rows = append(rows, Row{ Key: , Value: })
// 		}
// 	}
// }
