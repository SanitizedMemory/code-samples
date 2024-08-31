package main

import (
	"fmt"
	"os"
)

type CacheEntry struct {
	Key        []byte
	Occurances int
}

type PrefixTrie struct {
	Value    []byte
	Children []PrefixTrie
}

func makeBasicVocabulary() PrefixTrie {
	// Makes the base vocabulary to work with
	vocabulary := PrefixTrie{}
	var i byte
	for i = 1; i < 255; i++ {
		vocabulary.Children = append(
			vocabulary.Children, PrefixTrie{
				Value:    []byte{i},
				Children: []PrefixTrie{},
			},
		)
	}
	return vocabulary
}

func findLastMatch(parentNode *PrefixTrie, chars []byte, fullToken *[]byte) *PrefixTrie {
	if len(chars) == 0 {
		return nil
	}

	for idx, child := range parentNode.Children {
		value := make([]byte, len(child.Value))
		for i := 0; i < len(child.Value); i++ {
			value[i] = chars[i]
		}

		match := 0
		for i := 0; i < len(child.Value); i++ {
			if child.Value[i] == chars[i] {
				match += 1
			}
		}
		if match == len(child.Value) { // Current child is a match
			*fullToken = append(*fullToken, child.Value...)
			return findLastMatch(&parentNode.Children[idx], chars[len(child.Value):], fullToken)
		}
	}
	// No children matched
	return parentNode
}

func addToCache(cache map[*PrefixTrie][]CacheEntry, A *PrefixTrie, pathWalked []byte) {
	entries, ok := cache[A]
	if !ok {
		pathEntry := CacheEntry{
			Key:        pathWalked,
			Occurances: 1,
		}
		cache[A] = []CacheEntry{pathEntry}
		return
	}

	for idx, entry := range entries {
		if len(entry.Key) != len(pathWalked) {
			continue
		}

		counter := 0
		for i := 0; i < len(entry.Key); i++ {
			if entry.Key[i] != pathWalked[i] {
				break
			}
			counter += 1
		}
		if counter == len(entry.Key) {
			cache[A][idx].Occurances += 1
			return
		}
	}
}

func bytePairEncoding(parentNode *PrefixTrie, chars []byte) {
	cache := make(map[*PrefixTrie][]CacheEntry)
	slider := 0
	for {
		pathWalkedA := []byte{}
		pathWalkedB := []byte{}

		A := findLastMatch(parentNode, chars[slider:], &pathWalkedA)
		if A == nil {
			break
		}

		B := findLastMatch(parentNode, chars[slider+len(pathWalkedA):], &pathWalkedB)
		if B == nil {
			break
		}

		slider += len(pathWalkedA)
		addToCache(cache, A, pathWalkedB)
	}
	maxTrie, maxEntry := findMaxOccurance(cache)
	maxTrie.Children = append(maxTrie.Children, PrefixTrie{
		Value:    maxEntry.Key,
		Children: []PrefixTrie{},
	})
}

func printBPETrie(trie *PrefixTrie, path *[]byte) {
	*path = append(*path, trie.Value...)
	fmt.Println("`", string(*path), "`")

	for _, child := range trie.Children {
		printBPETrie(&child, path)
	}
	*path = (*path)[:len(*path)-len(trie.Value)]
}

func findMaxOccurance(cache map[*PrefixTrie][]CacheEntry) (*PrefixTrie, CacheEntry) {
	maxOccuranceTrie := &PrefixTrie{}
	maxOccuranceEntry := CacheEntry{}

	for trie, entries := range cache {
		for _, entry := range entries {
			if entry.Occurances > maxOccuranceEntry.Occurances ||
				(entry.Occurances == maxOccuranceEntry.Occurances && len(entry.Key) < len(maxOccuranceEntry.Key)) {
				maxOccuranceTrie = trie
				maxOccuranceEntry = entry
			}
		}
	}
	return maxOccuranceTrie, maxOccuranceEntry
}

func readFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func main() {
	file := readFile("data.txt")
	vocabulary := makeBasicVocabulary()
	vocabSize := 100
	for i := 0; i < vocabSize; i++ {
		bytePairEncoding(&vocabulary, file)
	}
	printBPETrie(&vocabulary, &[]byte{})
}
