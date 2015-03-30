package jiebago

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Entry struct {
	Word string
	Flag string
	Freq float64
}

type Loader interface {
	AddEntry(Entry)
}

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

func LoadDict(l Loader, dictFileName string, usingFlag bool) error {
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
