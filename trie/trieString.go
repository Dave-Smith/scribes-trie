package trie

import (
	"strings"
	// "github.com/dave-smith/trie-prefix/timing"
)

func (t *Trie) FromPrefix(prefix string) []string {
	// defer timing.Duration(timing.Track("prefix"))
	prefix = cleanInput(prefix)
	start := t.root.FindNode(prefix)
	return getWords(start, prefix[:len(prefix)-1])
}

func (trie *Trie) InsertWords(words []string) {
	for i := 0; i < len(words); i++ {
		trie.InsertWord(words[i])
	}
}

func (trie *Trie) InsertWord(word string) {
	word = cleanInput(word)
	// fmt.Printf("Adding word %s\n", word)
	trie.root.insert(word)
	trie.count++
}

func (t *Trie) DeleteWord(word string) {
	word = cleanInput(word)
	node := t.root.FindNode(word)
	node.isWord = false
}

func (n *Node) FindNode(prefix string) *Node {
	if len(prefix) == 0 {
		return n
	}

	r := prefix[0]

	for i := 0; i < len(n.children); i++ {
		if r == n.children[i].c {
			return n.children[i].FindNode(prefix[1:])
		}
	}

	node := NewNode(r)
	return &node
}

func getWords(node *Node, prefix string) []string {
	words := make([]string, 0)
	var current strings.Builder
	current.WriteString(prefix)
	current.WriteByte(node.c)

	if node.isWord {
		// fmt.Printf("found word %s\n", current.String())
		words = append(words, current.String())
	}

	for i := 0; i < len(node.children); i++ {
		words = append(words, getWords(node.children[i], current.String())...)
	}

	return words
}

func (n *Node) insert(word string) {
	if len(word) == 0 {
		n.isWord = true
		return
	}

	c := word[0]
	var next *Node
	children := n.children
	for i := 0; i < len(children); i++ {
		if children[i].c == c {
			next = children[i]
		}
	}

	if next == nil {
		nn := NewNode(c)
		next = &nn
		n.children = append(n.children, next)
	}

	next.insert(word[1:])
}

func cleanInput(input string) string {
	return strings.ToLower(strings.Trim(input, " \n\r"))
}
