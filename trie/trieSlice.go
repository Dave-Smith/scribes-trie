package trie

import (
	"bytes"
	// "github.com/dave-smith/trie-prefix/timing"
)

func (t *Trie) FastFromPrefix(prefix []byte) []string {
	// defer timing.Duration(timing.Track("fast prefix"))
	prefix = cleanWord(prefix)
	start := t.root.FastFindNode(prefix)
	return fastGetWords(start, prefix[:len(prefix)-1], t.maxDepth-len(prefix))
}

func (t *Trie) Suggestions(prefix []byte) []string {
	// defer timing.Duration(timing.Track("Suggestions"))
	// fmt.Printf("Getting suggestions for %s\n", prefix)
	prefix = cleanWord(prefix)
	start := t.root.FastFindNode(prefix)
	return fastGetWords(start, prefix[:len(prefix)-1], 2)
}

func (n *Node) FastFindNode(prefix []byte) *Node {
	if len(prefix) == 0 {
		return n
	}

	r := prefix[0]

	for i := 0; i < len(n.children); i++ {
		if r == n.children[i].c {
			return n.children[i].FastFindNode(prefix[1:])
		}
	}

	node := NewNode(r)
	return &node
}

func (trie *Trie) FastInsertWord(word []byte) {
	word = cleanWord(word)
	if len(word) > trie.maxDepth {
		trie.maxDepth = len(word)
	}

	trie.root.fastInsert(word)
	trie.count++
}

func (trie *Trie) FastInsertWords(words [][]byte) {
	for i := 0; i < len(words); i++ {
		trie.FastInsertWord(words[i])
	}
}

func (n *Node) fastInsert(word []byte) {
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
			break
		}
	}

	if next == nil {
		nn := NewNode(c)
		next = &nn
		n.children = append(n.children, next)
	}

	next.fastInsert(word[1:])
}
func getSuggestions(node *Node, prefix []byte, depth int) []string {
	word := append(prefix, node.c)
	words := make([]string, 0, 0)

	if node.isWord {
		words = append(words, string(word))
	}
	if depth == 0 {
		return []string{string(word)}
	}

	for i := 0; i < len(node.children); i++ {
		words = append(words, getSuggestions(node.children[i], word, depth-1)...)
	}
	return words
}

func fastGetWords(node *Node, prefix []byte, depth int) []string {
	words := make([]string, 0)
	word := append(prefix, node.c)

	if node.isWord {
		words = append(words, string(word))
	}

	if depth == 0 {
		return words
	}

	for i := 0; i < len(node.children); i++ {
		words = append(words, fastGetWords(node.children[i], word, depth-1)...)
	}

	return words
}

func cleanWord(word []byte) []byte {
	return bytes.ToLower(bytes.Trim(word, " \n\r"))
}
