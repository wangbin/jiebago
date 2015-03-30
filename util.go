package jiebago

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func dictPath(dictFileName string) (string, error) {
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

func LoadDict(l DictLoader, dictFileName string, usingFlag bool) error {
	dictFilePath, err := dictPath(dictFileName)
	if err != nil {
		return err
	}

	dictFile, err := os.Open(dictFilePath)
	if err != nil {
		return err
	}
	defer dictFile.Close()

	scanner := bufio.NewScanner(dictFile)
	var entry Entry
	var line string
	var fields []string
	for scanner.Scan() {
		line = scanner.Text()
		fields = strings.Split(line, " ")
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
