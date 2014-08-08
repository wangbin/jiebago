package jiebago

import (
	"bufio"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Node struct {
	Name     string
	SubNodes Trie
	IsLeaf   bool
}

type Trie map[string]*Node

type TopTrie struct {
	T       Trie
	MinFreq float64
	Total   float64
	Freq    map[string]float64
}

func newTopTrie(filename string) (*TopTrie, error) {
	var file_path string
	var topTrie *TopTrie
	if filepath.IsAbs(filename) {
		file_path = filename
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		file_path = filepath.Clean(filepath.Join(pwd, filename))
	}

	topTrie = &TopTrie{T: make(Trie), MinFreq: 100.0, Total: 0.0, Freq: make(map[string]float64)}
	file, openError := os.Open(file_path)
	if openError != nil {
		return nil, openError
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, readError := reader.ReadString('\n')
		if readError != nil && len(line) == 0 {
			break
		}
		words := strings.Split(line, " ")
		word, freqStr := words[0], words[1]
		freq, _ := strconv.ParseFloat(freqStr, 64)
		topTrie.Total += freq
		topTrie.addWord(word, freq)
	}
	var val float64
	for key := range topTrie.Freq {
		val = math.Log(topTrie.Freq[key] / topTrie.Total)
		if val < topTrie.MinFreq {
			topTrie.MinFreq = val
		}
		topTrie.Freq[key] = val
	}

	return topTrie, nil
}

func (tt *TopTrie) addWord(word string, freq float64) {
	tt.Freq[word] = freq
	var p Trie
	var node *Node
	var key string
	count := utf8.RuneCountInString(word)
	for index, c := range []rune(word) {
		if index == 0 {
			p = tt.T
		}
		key = string(c)
		if _, ok := p[key]; ok {
			node = p[key]
		} else {
			node = &Node{Name: key, IsLeaf: false}
			p[key] = node
			node.SubNodes = make(Trie)
		}
		if index == count-1 {
			p[key].IsLeaf = true
		}
		p = node.SubNodes
	}
}

func addWord(word string, freq float64, tag string) {
	if len(tag) > 0 {
		UserWordTagTab[word] = strings.TrimSpace(tag)
	}
	TT.addWord(word, freq)
}

func LoadUserDict(file_path string) error {
	file, openError := os.Open(file_path)
	if openError != nil {
		return openError
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		word, freqStr := words[0], words[1]
		word = strings.Replace(word, "\ufeff", "", 1)
		freq, freqErr := strconv.ParseFloat(freqStr, 64)
		if freqErr != nil {
			continue // TODO: how to handle wrong type of frequency?
		}
		tag := ""
		if len(words) == 3 {
			tag = words[2]
		}
		addWord(word, freq, tag)
	}

	return scanner.Err()
}
