package jiebago

import (
	"bufio"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Trie struct {
	Nodes  map[rune]*Trie
	IsLeaf bool
}

func NewTrie() *Trie {
	return &Trie{make(map[rune]*Trie), false}
}

type TopTrie struct {
	T       *Trie
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

	fi, err := os.Stat(file_path)
	if err != nil {
		return nil, err
	}
	log.Printf("Building Trie..., from %s\n", file_path)
	h := fmt.Sprintf("%x", md5.Sum([]byte(file_path)))
	cache_file_name := fmt.Sprintf("jieba.%s.cache", h)
	cache_path := filepath.Join(os.TempDir(), cache_file_name)
	isDictCached := true
	cache_fi, err := os.Stat(cache_path)

	if err != nil {
		isDictCached = false
	}

	if isDictCached {
		isDictCached = cache_fi.ModTime().After(fi.ModTime())
	}

	var cacheFile *os.File
	if isDictCached {
		cacheFile, err = os.Open(cache_path)
		if err != nil {
			isDictCached = false
		}
		defer cacheFile.Close()
	}
	if isDictCached {
		dec := gob.NewDecoder(cacheFile)
		err = dec.Decode(&topTrie)
		if err != nil {
			isDictCached = false
		} else {
			log.Printf("loaded model from cache %s\n", cache_path)
		}
	}

	if !isDictCached {
		topTrie = &TopTrie{T: NewTrie(), MinFreq: 100.0, Total: 0.0, Freq: make(map[string]float64)}
		file, openError := os.Open(file_path)
		if openError != nil {
			return nil, openError
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			words := strings.Split(line, " ")
			word, freqStr := words[0], words[1]
			freq, _ := strconv.ParseFloat(freqStr, 64)
			topTrie.Total += freq
			topTrie.addWord(word, freq)
		}
		if scanErr := scanner.Err(); scanErr != nil {
			return nil, scanErr
		}

		var val float64
		for key := range topTrie.Freq {
			val = math.Log(topTrie.Freq[key] / topTrie.Total)
			if val < topTrie.MinFreq {
				topTrie.MinFreq = val
			}
			topTrie.Freq[key] = val
		}

		// dump topTrie
		cacheFile, err = os.OpenFile(cache_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return topTrie, err
		}
		defer cacheFile.Close()
		enc := gob.NewEncoder(cacheFile)
		err := enc.Encode(topTrie)
		if err != nil {
			return topTrie, err
		} else {
			log.Printf("dumped model from cache %s\n", cache_path)
		}
	}
	return topTrie, nil
}

func (tt *TopTrie) addWord(word string, freq float64) {
	tt.Freq[word] = freq
	var p *Trie
	runes := []rune(word)
	count := len(runes)
	for index, key := range runes {
		if index == 0 {
			p = tt.T
		}
		if _, ok := p.Nodes[key]; !ok {
			p.Nodes[key] = NewTrie()
		}
		if index == count-1 {
			p.Nodes[key].IsLeaf = true
		}
		p = p.Nodes[key]
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
