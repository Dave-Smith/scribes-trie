package trie

type Trie struct {
	root     *Node
	count    int
	maxDepth int
}
type Node struct {
	c        byte
	children []*Node
	isWord   bool
}

func NewTrie() Trie {
	root := NewNode(0)
	return Trie{&root, 0, 0}
}

func NewNode(c byte) Node {
	var children []*Node = []*Node{}
	return Node{c, children, false}
}

func (trie Trie) Count() int {
	return trie.count
}
