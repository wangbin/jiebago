package jiebago

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type WordTagFreq struct {
	Word, Tag string
	Freq      float64
}

func ParseDictFile(dictFilePath string) (wtfs []*WordTagFreq, err error) {
	var dictFile *os.File
	dictFile, err = os.Open(dictFilePath)
	if err != nil {
		return
	}
	defer dictFile.Close()
	scanner := bufio.NewScanner(dictFile)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		length := len(fields)
		word := fields[0]
		word = strings.Replace(word, "\ufeff", "", 1)
		wtf := &WordTagFreq{Word: word}
		if length > 1 {
			wtf.Freq, err = strconv.ParseFloat(fields[1], 64)
			if err != nil {
				return
			}
		}
		if length > 2 {
			wtf.Tag = fields[2]
		}
		wtfs = append(wtfs, wtf)
	}
	err = scanner.Err()
	return
}
