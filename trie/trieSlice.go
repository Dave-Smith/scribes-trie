package trie

import (
	"bytes"
)

func (t *Trie) FastFromPrefix(prefix []byte) []string {
	prefix = cleanWord(prefix)
	start := t.root.FastFindNode(prefix)
	if start == nil {
		return []string{}
	}
	return fastGetWords(start, prefix[:len(prefix)-1], t.maxDepth-len(prefix))
}

func (t *Trie) Suggestions(prefix []byte) []string {
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

	return nil
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
	children := n.children
	for i := 0; i < len(children); i++ {
		if children[i].c == c {
			children[i].fastInsert(word[1:])
			return
		}
	}

	next := NewNode(c)
	n.children = append(n.children, &next)

	(&next).fastInsert(word[1:])
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
	trimmed := bytes.Trim(word, " \n\r")

	// lowercasing
	// for i := range trimmed {
	// 	if trimmed[i] >= 'A' && trimmed[i] <= 'Z' {
	// 		trimmed[i] += 32
	// 	}
	// }
	// return trimmed
	return bytes.ToLower(trimmed)
}
