package jiebago

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var T *Trie

type Trie struct {
	Total float64
	Freq  map[string]float64
}

func (t Trie) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(t.Total)
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
	err := dec.Decode(&t.Total)
	if err != nil {
		return err
	}
	err = dec.Decode(&t.Freq)
	if err != nil {
		return err
	}
	return nil
}

func newTrie(dictFileName string) (*Trie, error) {
	dictFilePath, err := DictPath(dictFileName)
	if err != nil {
		return nil, err
	}

	dictFileInfo, err := os.Stat(dictFilePath)
	if err != nil {
		return nil, err
	}

	log.Printf("Building Trie..., from %s\n", dictFilePath)
	h := fmt.Sprintf("%x", md5.Sum([]byte(dictFilePath)))
	cacheFileName := fmt.Sprintf("jieba.%s.cache", h)
	cacheFilePath := filepath.Join(os.TempDir(), cacheFileName)
	isDictCached := true

	cacheFileInfo, err := os.Stat(cacheFilePath)
	if err != nil {
		isDictCached = false
	}

	if isDictCached {
		isDictCached = cacheFileInfo.ModTime().After(dictFileInfo.ModTime())
	}

	var cacheFile *os.File
	if isDictCached {
		cacheFile, err = os.Open(cacheFilePath)
		if err != nil {
			isDictCached = false
		}
		defer cacheFile.Close()
	}

	var trie *Trie

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
		trie = &Trie{Total: 0.0, Freq: make(map[string]float64)}

		dictFile, err := os.Open(dictFilePath)
		if err != nil {
			return nil, err
		}
		defer dictFile.Close()

		wtfs, err := ParseDictFile(dictFile)
		if err != nil {
			return nil, err
		}

		for _, wtf := range wtfs {
			trie.addWord(wtf)
		}
		// dump trie
		cacheFile, err = os.OpenFile(cacheFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return trie, err
		}
		defer cacheFile.Close()
		enc := gob.NewEncoder(cacheFile)
		err = enc.Encode(trie)
		if err != nil {
			return trie, err
		} else {
			log.Printf("dumped model from cache %s\n", cacheFilePath)
		}
	}
	return trie, nil
}

func (t *Trie) addWord(wtf *WordTagFreq) {
	t.Freq[wtf.Word] = wtf.Freq
	t.Total += wtf.Freq
	runes := []rune(wtf.Word)
	count := len(runes)
	for i := 0; i < count; i++ {
		wfrag := string(runes[0 : i+1])
		if _, ok := t.Freq[wfrag]; !ok {
			t.Freq[wfrag] = 0.0
		}
	}
}
func addWord(wtf *WordTagFreq) {
	if len(wtf.Tag) > 0 {
		UserWordTagTab[wtf.Word] = strings.TrimSpace(wtf.Tag)
	}
	T.addWord(wtf)
}

func LoadUserDict(dictFilePath string) error {
	dictFile, err := os.Open(dictFilePath)
	if err != nil {
		return err
	}
	defer dictFile.Close()

	wtfs, err := ParseDictFile(dictFile)
	if err != nil {
		return err
	}
	for _, wtf := range wtfs {
		addWord(wtf)
	}
	return nil
}

func SetDictionary(dict_path string) (err error) {
	T, err = newTrie(dict_path)
	return
}
