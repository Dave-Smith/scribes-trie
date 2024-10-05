package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

type Trie struct {
	root  *Node
	count int
}
type Node struct {
	c        byte
	children []*Node
	isWord   bool
}

func main() {
	fmt.Println("Type ':q' to exit")
	t2, _ := constructTrie()
	t, _ := fastConstructTrie()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Search for a word: ")
		text, _ := reader.ReadBytes(10)
		text = cleanWord(text)
		if bytes.Compare(text, []byte(":q")) == 0 {
			break
		}
		if len(text) == 0 {
			fmt.Println("Try again.")
			continue
		}

		fmt.Printf("searching for words starting with %s . . . \n", text)
		words2 := t2.FromPrefix(string(text))
		words2 = t2.FromPrefix(string(text))
		words := t.FastFromPrefix(text)
		words = t.FastFromPrefix(text)
		fmt.Printf("found %d in the trie\n", len(words))
		count := len(words2)
		if count >= 10 {
			fmt.Printf("%d words found, truncating results to 10\n", count)
			words2 = words2[:10]
		}
		fmt.Printf("words: %v\n", words2)
	}
}

func fastConstructTrie() (Trie, error) {
	defer duration(track("fast construct trie"))
	file, err := os.Open("words.txt")
	words := 0
	var t Trie
	if err != nil {
		return t, err
	}
	defer file.Close()
	t = NewTrie()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	word := []byte{}
	for scanner.Scan() {
		b := scanner.Bytes()
		if b[0] != 10 {
			word = append(word, b[0])
			continue
		}

		t.FastInsertWord(word)
		words++
		word = []byte{}
	}

	fmt.Printf("Constructed the prefix trie with %d words\n", words)

	return t, nil
}

func constructTrie() (Trie, error) {
	defer duration(track("regular construct trie"))
	file, err := os.Open("words.txt")
	words := 0
	var t Trie
	if err != nil {
		return t, err
	}
	defer file.Close()
	t = NewTrie()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		t.InsertWord(scanner.Text())
		words++
	}

	fmt.Printf("Constructed the prefix trie with %d words\n", words)

	return t, nil
}

func NewTrie() Trie {
	root := NewNode(0)
	return Trie{root, 0}
}

func NewNode(c byte) *Node {
	var children []*Node = make([]*Node, 0, 50)
	return &Node{c, children, false}
}

func (t *Trie) FromPrefix(prefix string) []string {
	defer duration(track("prefix"))
	prefix = cleanInput(prefix)
	start := t.root.FindNode(prefix)
	// fmt.Printf("starting at node %v\n", start.c)
	return getWords(start, prefix[:len(prefix)-1])
}

func (t *Trie) FastFromPrefix(prefix []byte) []string {
	defer duration(track("fast prefix"))
	prefix = cleanWord(prefix)
	start := t.root.FastFindNode(prefix)
	// fmt.Printf("starting at node %v\n", start.c)
	return fastGetWords(start, prefix[:len(prefix)-1])
}

func (trie *Trie) InsertWord(word string) {
	word = cleanInput(word)
	// fmt.Printf("Adding word %s\n", word)
	trie.root.insert(word)
	trie.count++
}

func cleanWord(word []byte) []byte {
	return bytes.ToLower(bytes.Trim(word, " \n\r"))
}

func (trie *Trie) FastInsertWord(word []byte) {
	word = cleanWord(word)
	trie.root.fastInsert(word)
	trie.count++
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
		next = NewNode(c)
		n.children = append(n.children, next)
	}

	next.fastInsert(word[1:])
}

func (trie Trie) Count() int {
	return trie.count
}

func (t *Trie) DeleteWord(word string) {
	word = cleanInput(word)
	node := t.root.FindNode(word)
	node.isWord = false
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

func fastGetWords(node *Node, prefix []byte) []string {
	words := make([]string, 0)
	word := append(prefix, node.c)

	if node.isWord {
		// fmt.Printf("found word %s\n", current.String())
		words = append(words, string(word))
	}

	for i := 0; i < len(node.children); i++ {
		words = append(words, fastGetWords(node.children[i], word)...)
	}

	return words
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

	return NewNode(r)
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

	return NewNode(r)
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
		next = NewNode(c)
		n.children = append(n.children, next)
	}

	next.insert(word[1:])
}

func cleanInput(input string) string {
	return strings.ToLower(strings.Trim(input, " \n\r"))
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}
