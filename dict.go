package jiebago

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ParseDictFile(dictFilePath string) ([]*Entry, error) {
	dictFile, err := os.Open(dictFilePath)
	if err != nil {
		return nil, err
	}
	defer dictFile.Close()
	entries := make([]*Entry, 0)
	scanner := bufio.NewScanner(dictFile)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		length := len(fields)
		word := fields[0]
		word = strings.Replace(word, "\ufeff", "", 1)
		entry := NewEntry()
		entry.Word = word
		if length > 1 {
			entry.Freq, err = strconv.ParseFloat(fields[1], 64)
			if err != nil {
				return nil, err
			}
		}
		if length > 2 {
			entry.Flag = fields[2]
		}
		entries = append(entries, entry)
	}
	return entries, scanner.Err()
}
