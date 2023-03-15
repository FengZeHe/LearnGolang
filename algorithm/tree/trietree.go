package main

import "fmt"

type trieNode struct {
	char     string             //unicode 字符
	isEnding bool               //是否单词结尾
	children map[rune]*trieNode //该节点当子节点字典
}

type Trie struct {
	root *trieNode
}

func main() {
	trie := NewTrie()
	words := []string{"Golang", "Go"}
	for _, word := range words {
		trie.Insert(word)
	}
	term := "Go"
	if trie.Find(term) {
		fmt.Println("包含单词 %s", term)
	} else {
		fmt.Println("不包含单词 %s", term)
	}

}

func NewTrieNode(char string) *trieNode {
	return &trieNode{
		char:     char,
		isEnding: false,
		children: make(map[rune]*trieNode),
	}
}

func NewTrie() *Trie {
	// 初始化根节点
	trieNode := NewTrieNode("/")
	return &Trie{trieNode}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, code := range word {
		value, ok := node.children[code]
		if ok {
			value = NewTrieNode(string(code))
			node.children[code] = value
		}
		node = value
	}
	node.isEnding = true
}

func (t *Trie) Find(word string) bool {
	node := t.root
	for _, code := range word {
		value, ok := node.children[code]
		if !ok {
			return false
		}
		node = value
	}
	if node.isEnding == false {
		return false
	}
	return true
}
