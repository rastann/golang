package main

import "fmt"

const alphabetSize = 26

type node struct {
	children [26]*node
	isEnd bool
}

type trie struct {
	root *node
}

func createTrie() *trie {
	result := &trie{root: &node{}}
	return result
}

func (t *trie) insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i :=0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *trie) search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i :=0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	ttrie := createTrie()
	ttrie.insert("aragorn")

	fmt.Print(ttrie.search("aragorn"))
}
