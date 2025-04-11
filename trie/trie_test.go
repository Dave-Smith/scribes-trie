package trie

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

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

func BenchmarkFindBytes(b *testing.B) {
	t := NewTrie()
	t.FastInsertWords(wordBytes)
	word := []byte("help")
	for i := 0; i < b.N; i++ {
		t.FastFromPrefix(word)
	}
}

func BenchmarkToLower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, w := range bytesLower {
			bytes.ToUpper(w)
		}
	}
}

var words []string
var wordBytes [][]byte
var bytesLower [][]byte

func init() {
	file, _ := os.Open("words-upper.txt")
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

	for _, w := range wordsLower {
		bytesLower = append(bytesLower, []byte(w))
	}
}

var wordsLower []string = []string{"append", "bytes", "words", "word", "scanner", "open", "defer", "string", "if", "split", "file", "init", "continue", "scan"}
