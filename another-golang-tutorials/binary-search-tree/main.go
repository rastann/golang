package main

import "fmt"

type node struct {
	key int
	left *node
	right *node
}

func (n *node) insert(k int) {
	if n.key < k {
		if n.right == nil {
			n.right = &node{key: k}
		} else {
			n.right.insert(k)
		}
	} else if n.key > k {
		if n.left == nil {
			n.left = &node{key: k}
		} else {
			n.left.insert(k)
		}
	}
}

func (n *node) search(k int) bool {
	if n == nil {
		return false
	}
	if n.key < k {
		return n.right.search(k)
	} else if n.key > k {
		return n.left.search(k)
	}
	return true
}

func main() {
	ttree := &node{key: 1}
	ttree.insert(50)
	fmt.Print(ttree)
}
