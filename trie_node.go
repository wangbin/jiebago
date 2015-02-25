package jiebago

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var T *Trie

type Trie struct {
	Nodes mapset.Set
	Total float64
	Freq  map[string]float64
}

func (t Trie) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(t.Nodes.ToSlice())
	if err != nil {
		return nil, err
	}
	err = enc.Encode(t.Total)
	if err != nil {
		return nil, err
	}
	err = enc.Encode(t.Freq)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (t *Trie) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	dec := gob.NewDecoder(b)
	var nodes []interface{}
	err := dec.Decode(&nodes)
	if err != nil {
		return err
	}
	t.Nodes = mapset.NewSetFromSlice(nodes)
	err = dec.Decode(&t.Total)
	if err != nil {
		return err
	}
	err = dec.Decode(&t.Freq)
	if err != nil {
		return err
	}
	return nil
}

func newTrie(fileName string) (*Trie, error) {
	var filePath string
	var trie *Trie
	if filepath.IsAbs(fileName) {
		filePath = fileName
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		filePath = filepath.Clean(filepath.Join(pwd, fileName))
	}

	fi, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	log.Printf("Building Trie..., from %s\n", filePath)
	h := fmt.Sprintf("%x", md5.Sum([]byte(filePath)))
	cacheFileName := fmt.Sprintf("jieba.%s.cache", h)
	cacheFilePath := filepath.Join(os.TempDir(), cacheFileName)
	isDictCached := true
	cacheFileInfo, err := os.Stat(cacheFilePath)

	if err != nil {
		isDictCached = false
	}

	if isDictCached {
		isDictCached = cacheFileInfo.ModTime().After(fi.ModTime())
	}

	var cacheFile *os.File
	if isDictCached {
		cacheFile, err = os.Open(cacheFilePath)
		if err != nil {
			isDictCached = false
		}
		defer cacheFile.Close()
	}
	if isDictCached {
		dec := gob.NewDecoder(cacheFile)
		err = dec.Decode(&trie)
		if err != nil {
			isDictCached = false
		} else {
			log.Printf("loaded model from cache %s\n", cacheFilePath)
		}
	}

	if !isDictCached {
		trie = &Trie{Nodes: mapset.NewSet(), Total: 0.0,
			Freq: make(map[string]float64)}

		file, openError := os.Open(filePath)
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
			trie.addWord(word, freq)
		}
		if scanErr := scanner.Err(); scanErr != nil {
			return nil, scanErr
		}

		// dump trie
		cacheFile, err = os.OpenFile(cacheFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return trie, err
		}
		defer cacheFile.Close()
		enc := gob.NewEncoder(cacheFile)
		err := enc.Encode(trie)
		if err != nil {
			return trie, err
		} else {
			log.Printf("dumped model from cache %s\n", cacheFilePath)
		}
	}
	return trie, nil
}

func (t *Trie) addWord(word string, freq float64) {
	t.Freq[word] = freq
	t.Total += freq
	runes := []rune(word)
	count := len(runes)
	for i := 0; i < count; i++ {
		t.Nodes.Add(string(runes[:i+1]))
	}
}

func addWord(word string, freq float64, tag string) {
	if len(tag) > 0 {
		UserWordTagTab[word] = strings.TrimSpace(tag)
	}
	T.addWord(word, freq)
}

func LoadUserDict(filePath string) error {
	file, openError := os.Open(filePath)
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

func SetDictionary(dict_path string) (err error) {
	T, err = newTrie(dict_path)
	return
}
