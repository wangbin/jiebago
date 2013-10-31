package jiebago

import (
	"bufio"
	"crypto/sha1"
	"encoding/gob"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	CACHE_NAME        = "jieba.gob"
	USER_CACHE_PREFIX = "jieba.user."
	USER_CACHE_SUFFIX = ".gob"
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

func hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func getUserCacheName(prefix string, path string, suffix string) string {
	return fmt.Sprintf("%s%s%s", prefix, hash(path), suffix)
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

	_, curFileName, _, _ := runtime.Caller(1)
	_curpath := filepath.Dir(curFileName)
	abs_path := filepath.Join(_curpath, Dictionary)
	var cache_file string
	if file_path == abs_path {
		cache_file = filepath.Join(os.TempDir(), CACHE_NAME)
	} else {
		cache_file = filepath.Join(os.TempDir(),
			getUserCacheName(USER_CACHE_PREFIX, abs_path, USER_CACHE_SUFFIX))
	}

	cacheFileStat, cacheErr := os.Stat(cache_file)
	dictFileStat, _ := os.Stat(abs_path)
	if cacheErr == nil {
		if cacheFileStat.ModTime().After(dictFileStat.ModTime()) {
			cacheFile, openError := os.Open(cache_file)
			if openError == nil {
				dec := gob.NewDecoder(cacheFile)
				err := dec.Decode(&topTrie)
				if err == nil {
					return topTrie, nil
				}
			}
		}
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

	cacheFile_, _ := os.OpenFile(cache_file, os.O_CREATE|os.O_WRONLY, 0644)
	defer cacheFile_.Close()
	enc := gob.NewEncoder(cacheFile_)
	enc.Encode(topTrie)

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

	reader := bufio.NewReader(file)
	for {
		line, readError := reader.ReadString('\n')
		if readError != nil && len(line) == 0 {
			break
		}
		words := strings.Split(line, " ")
		word, freqStr := words[0], words[1]
		freq, _ := strconv.ParseFloat(freqStr, 64)
		TT.addWord(word, freq)
	}
	return nil
}
