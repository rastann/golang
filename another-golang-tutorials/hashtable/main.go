package main

import (
	"fmt"
)

const arraySize = 100

type hashTable struct {
	array [arraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func createHashTable() *hashTable {
	result := &hashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func (h *hashTable) insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *hashTable) search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *hashTable) delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func (h *hashTable) print() {
	for i := range h.array {
		currentNode := h.array[i].head
		for currentNode != nil {
			fmt.Println("index: ", i , " value: ", currentNode.key)
			currentNode = currentNode.next
		}
	}
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % arraySize
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, " already exists")
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

func main() {
	hashTable := createHashTable()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
		"TOEKN",
	}

	for _, v := range list {
		hashTable.insert(v)
	}

	hashTable.print()
	//fmt.Println(hashTable.search("STAN"))
}
