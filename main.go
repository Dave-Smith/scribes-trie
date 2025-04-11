package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/dave-smith/trie-prefix/timing"
	"github.com/dave-smith/trie-prefix/trie"
)

func main() {
	fmt.Println("Type ':q' to exit")
	t, _ := fastConstructTrie()
	// t, _ := fastConstructTrie()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Search for a word: ")
		text, _ := reader.ReadBytes(10)
		text = cleanPrompt(text)
		if bytes.Compare(text, []byte(":q")) == 0 {
			break
		}
		if len(text) == 0 {
			fmt.Println("Try again.")
			continue
		}

		fmt.Printf("searching for words starting with %s . . . \n", text)
		// words2 := t2.FromPrefix(string(text))
		// words2 = t2.FromPrefix(string(text))
		words := []string{}
		if text[len(text)-1] == '?' {
			func() {
				defer timing.Duration(timing.Track("Suggestion"))
				words = t.Suggestions(text[:len(text)-1])
			}()
		} else {
			func() {
				defer timing.Duration(timing.Track("Find All"))
				words = t.FastFromPrefix(text)
			}()
		}
		// words := t.FastFromPrefix(text)
		// words = t.FastFromPrefix(text)
		fmt.Printf("found %d in the trie\n", len(words))
		count := len(words)
		if count >= 10 {
			fmt.Printf("%d words found, truncating results to 10\n", count)
			words = words[:10]
		}
		fmt.Printf("words: %v\n", words)
	}
}

func fastConstructTrie() (trie.Trie, error) {
	defer timing.Duration(timing.Track("fast construct trie"))
	file, err := os.Open("words.txt")
	words := 0
	var t trie.Trie
	if err != nil {
		return t, err
	}
	defer file.Close()
	t = trie.NewTrie()
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

func cleanPrompt(prompt []byte) []byte {
	return bytes.ToLower(bytes.Trim(prompt, " \n\r"))
}
