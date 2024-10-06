package trie

import (
	"bufio"
	"os"
	"testing"
)

var wordsAsString = []string{
	"help",
	"hi",
	"hip",
	"hint",
	"hints",
	"helpful",
	"helps",
	"and",
	"an",
	"ant",
	"ants",
	"aunt",
	"aunts",
	"aunty",
	"get",
	"gets",
	"go",
	"got",
	"goth",
}
var wordsAsByteSlice = fromString(wordsAsString)

func fromString(words []string) [][]byte {
	b := [][]byte{}
	for i := 0; i < len(words); i++ {
		b = append(b, []byte(words[i]))
	}
	return b
}

func BenchmarkInsertBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := NewTrie()
		t.FastInsertWords(wordBytes)
	}
}

func BenchmarkInsertStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := NewTrie()
		t.InsertWords(words)
	}
}

func BenchmarkFindBytes(b *testing.B) {
	t := NewTrie()
	t.InsertWords(words)
	word := []byte("help")
	for i := 0; i < b.N; i++ {
		t.FastFromPrefix(word)
	}
}

func BenchmarkFindString(b *testing.B) {
	t := NewTrie()
	t.InsertWords(words)
	word := "help"
	for i := 0; i < b.N; i++ {
		t.FromPrefix(word)
	}
}

var words []string
var wordBytes [][]byte

func init() {
	file, _ := os.Open("words.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	word := []byte{}
	for scanner.Scan() {
		b := scanner.Bytes()
		if b[0] != 10 {
			word = append(word, b[0])
			continue
		}
		wordBytes = append(wordBytes, word)
		words = append(words, string(word))
		word = []byte{}
	}
}
