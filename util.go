package jiebago

import (
	//	"bufio"
	//	"crypto/md5"
	//	"encoding/gob"
	//	"fmt"
	"os"
	"path/filepath"
	"regexp"
	//	"strconv"
	//	"strings"
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

/*
func cachePath(dictPath string) string {
	return filepath.Join(os.TempDir(),
		fmt.Sprintf("jieba.%x.cache", md5.Sum([]byte(f.dictFilePath))))
}

func fileInfo(filePath string, missingOk bool) (*os.FileInfo, err) {
	fileInfo, err := os.Stat(filePath)
	if missingOk && err.Err == os.ErrNotExist {
		return fileInfo, nil
	}
	return fileInfo, err
}

func isCached(dictPath, cachePath string) (bool, error) {
	dictFileInfo, err := fileInfo(dictPath, false)
	if err != nil {
		return false, err
	}
	cacheFileInfo, err := fileInfo(cachePath, true)
	if err != nil {
		return false, err
	}
	return cacheFileInfo.ModTime().After(dictFileInfo.ModTime()), nil
}

func load(cachePath string, d DictLoader) error {
	dec := gob.NewDecoder(cacheFile)
	return dec.Decode(&d)
}

func read(dictPath, d DictLoader, pos bool) error {
	dictFile, err := os.Open(dictFilePath)
	if err != nil {
		return err
	}
	defer dictFile.Close()
	scanner := bufio.NewScanner(dictFile)
	var token *Token
	var line string
	var fields []string
	for scanner.Scan() {
		line = scanner.Text()
		fields = strings.Split(line, " ")
		token = &Token{Term: strings.Replace(fields[0], "\ufeff", "", 1)}
		if length := len(fields); length > 1 {
			token.Freq, err = strconv.ParseFloat(fields[1], 64)
			if err != nil {
				return err
			}
			if pos && length > 2 {
				token.Pos = fields[2]
			}
		}
		d.Add(token)
	}
	return scanner.Err()
}

func dump(cachePath string, d DictLoader) error {
	cacheFile, err = os.OpenFile(cachePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer cacheFile.Close()
	enc := gob.NewEncoder(cacheFile)
	return enc.Encode(d)
}

func SetDict(s Segmenter, dictName string, pos bool) error {
	dictPath, err := DictPath(dictName)
	if err != nil {
		return err
	}
	cachePath = cachePath(dictPath)
	cached, err := isCached(dictPath, cachePath)
	if err != nil {
		return err
	}

	if cached {
		err = load(cachePath, s)
		if err == nil {
			return nil
		}
		cached = false
	}

	err = read(dictPath, s, pos)
	if err != nil {
		return err
	}

	err = dump(cachePath, s)
	if err != nil {
		return err
	}
}

func LoadUserDict(dictName string, s Segmenter, pos bool) error {
	dictPath, err := DictPath(dictName)
	if err != nil {
		return err
	}
	return read(dictPath, s, pos)
}
*/

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
