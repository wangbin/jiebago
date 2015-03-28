package jiebago

import (
	"bufio"
	"crypto/md5"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func DictPath(dictFileName string) (string, error) {
	if filepath.IsAbs(dictFileName) {
		return dictFileName, nil
	}
	var dictFilePath string
	cwd, err := os.Getwd()
	if err != nil {
		return dictFilePath, err
	}
	dictFilePath = filepath.Clean(filepath.Join(cwd, dictFileName))
	return dictFilePath, nil
}

func LoadDict(l DictLoader, dictFilePath string, usingFlag bool) error {
	log.Printf("Building Trie..., from %s\n", dictFilePath)

	dictFile, err := os.Open(dictFilePath)
	if err != nil {
		return err
	}
	defer dictFile.Close()
	scanner := bufio.NewScanner(dictFile)
	var entry *Entry
	var line string
	var fields []string
	for scanner.Scan() {
		line = scanner.Text()
		fields = strings.Split(line, " ")
		entry = NewEntry()
		entry.Word = strings.Replace(fields[0], "\ufeff", "", 1)
		if length := len(fields); length > 1 {
			entry.Freq, err = strconv.ParseFloat(fields[1], 64)
			if err != nil {
				return err
			}
			if usingFlag && length > 2 {
				entry.Flag = fields[2]
			}
		}
		l.AddEntry(entry)
	}
	return scanner.Err()
}

func cachePath(dictPath string) string {
	return filepath.Join(os.TempDir(),
		fmt.Sprintf("jieba.%x.cache", md5.Sum([]byte(dictPath))))
}

func cached(dictPath, cachePath string) (bool, error) {
	dictFileInfo, err := os.Stat(dictPath)
	if err != nil {
		return false, err
	}
	cacheFileInfo, err := os.Stat(cachePath)
	if err != nil {
		return false, nil
	}
	return cacheFileInfo.ModTime().After(dictFileInfo.ModTime()), nil
}

func load(l DictLoader, cachePath string) error {
	cacheFile, err := os.Open(cachePath)
	if err != nil {
		return err
	}
	defer cacheFile.Close()

	dec := gob.NewDecoder(cacheFile)
	return dec.Decode(l)
}

func dump(l DictLoader, cachePath string) error {
	cacheFile, err := os.OpenFile(cachePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer cacheFile.Close()
	enc := gob.NewEncoder(cacheFile)
	return enc.Encode(l)
}

func SetDict(l DictLoader, dictName string, pos bool) error {
	dictPath, err := DictPath(dictName)
	if err != nil {
		return err
	}
	cachePath := cachePath(dictPath)
	cached, err := cached(dictPath, cachePath)
	if err != nil {
		return err
	}

	if cached {
		err = load(l, cachePath)
		if err == nil {
			log.Printf("loaded model from cache %s\n", cachePath)
			return nil
		}
		cached = false
	}

	err = LoadDict(l, dictPath, pos)
	if err != nil {
		return err
	}

	err = dump(l, cachePath)
	if err == nil {
		log.Printf("dumped model from cache %s\n", cachePath)
		return nil
	}
	return err
}

// Split sentence using regular expression.
func RegexpSplit(r *regexp.Regexp, sentence string) chan string {
	result := make(chan string)
	go func() {
		locs := r.FindAllStringIndex(sentence, -1)
		lastLoc := 0
		for _, loc := range locs {
			if loc[0] == lastLoc {
				result <- sentence[loc[0]:loc[1]]
			} else {
				result <- sentence[lastLoc:loc[0]]
				result <- sentence[loc[0]:loc[1]]
			}
			lastLoc = loc[1]
		}
		if lastLoc < len(sentence) {
			result <- sentence[lastLoc:]
		}
		close(result)
	}()
	return result
}
