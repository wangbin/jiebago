package jiebago

import (
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Jieba struct {
	Total float64
	Freq  map[string]float64
}

func (j *Jieba) load(dictFileName string) error {
	dictFilePath, err := DictPath(dictFileName)
	if err != nil {
		return err
	}

	dictFileInfo, err := os.Stat(dictFilePath)
	if err != nil {
		return err
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

	if isDictCached {
		dec := gob.NewDecoder(cacheFile)
		err = dec.Decode(&j)
		if err != nil {
			isDictCached = false
		} else {
			log.Printf("loaded model from cache %s\n", cacheFilePath)
		}
	}

	if !isDictCached {
		err = LoadDict(j, dictFilePath, false)
		if err != nil {
			return err
		}
		// dump trie
		cacheFile, err = os.OpenFile(cacheFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer cacheFile.Close()
		enc := gob.NewEncoder(cacheFile)
		err = enc.Encode(j)
		if err != nil {
			return err
		} else {
			log.Printf("dumped model from cache %s\n", cacheFilePath)
		}
	}
	return nil
}

func (j *Jieba) AddEntry(entry *Entry) {
	j.Add(entry.Word, entry.Freq)
}

func (j *Jieba) Add(word string, freq float64) {
	j.Freq[word] = freq
	j.Total += freq
	runes := []rune(word)
	for i := 0; i < len(runes); i++ {
		frag := string(runes[0 : i+1])
		if _, ok := j.Freq[frag]; !ok {
			j.Freq[frag] = 0.0
		}
	}
}

// Load user specified dictionary file.
func (j *Jieba) LoadUserDict(dictFilePath string) error {
	return LoadDict(j, dictFilePath, false)
}

// Set the dictionary, could be absolute path of dictionary file, or dictionary
// name in current directory. This function must be called before cut any
// sentence.
func NewJieba(dictFileName string) (*Jieba, error) {
	j := &Jieba{Total: 0.0, Freq: make(map[string]float64)}
	err := j.load(dictFileName)
	return j, err
}
