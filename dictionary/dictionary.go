package dictionary

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type DictLoader interface {
	Load(<-chan Token)
	AddToken(Token)
}

func loadDictionary(file *os.File) (<-chan Token, <-chan error) {
	tokenCh, errCh := make(chan Token), make(chan error)

	go func() {
		defer close(tokenCh)
		defer close(errCh)
		scanner := bufio.NewScanner(file)
		var token Token
		var line string
		var fields []string
		var err error
		for scanner.Scan() {
			line = scanner.Text()
			fields = strings.Split(line, " ")
			token.text = strings.TrimSpace(strings.Replace(fields[0], "\ufeff", "", 1))
			if length := len(fields); length > 1 {
				token.frequency, err = strconv.ParseFloat(fields[1], 64)
				if err != nil {
					errCh <- err
					return
				}
				if length > 2 {
					token.pos = strings.TrimSpace(fields[2])
				}
			}
			tokenCh <- token
		}

		if err = scanner.Err(); err != nil {
			errCh <- err
		}
	}()
	return tokenCh, errCh

}

func LoadDictionary(dl DictLoader, fileName string) error {
	filePath, err := dictPath(fileName)
	if err != nil {
		return err
	}
	dictFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer dictFile.Close()
	tokenCh, errCh := loadDictionary(dictFile)
	dl.Load(tokenCh)

	return <-errCh

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
