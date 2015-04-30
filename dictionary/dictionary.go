package dictionary

import (
	"bufio"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type Dictionary struct {
	total, logTotal float64
	freqMap         map[string]float64
	posMap          map[string]string
	sync.RWMutex
}

func New() *Dictionary {
	return &Dictionary{
		freqMap: make(map[string]float64),
		posMap:  make(map[string]string)}
}

func (d *Dictionary) addToken(token Token) {
	d.freqMap[token.text] = token.frequency
	d.total += token.frequency
	runes := []rune(token.text)
	if len(token.pos) > 0 {
		d.posMap[token.text] = token.pos
	}
	n := len(runes)
	for i := 0; i < n; i++ {
		frag := string(runes[:i+1])
		if _, ok := d.freqMap[frag]; !ok {
			d.freqMap[frag] = 0.0
		}
	}
}

func (d *Dictionary) updateLogTotal() {
	d.logTotal = math.Log(d.total)
}

func (d *Dictionary) AddToken(token Token) {
	d.Lock()
	d.addToken(token)
	d.updateLogTotal()
	d.Unlock()
}

func (d Dictionary) Total() float64 {
	return d.total
}

func (d Dictionary) LogTotal() float64 {
	return d.logTotal
}

func (d Dictionary) Frequency(key string) (float64, bool) {
	d.RLock()
	freq, ok := d.freqMap[key]
	d.RUnlock()
	return freq, ok
}

func (d Dictionary) Pos(key string) (string, bool) {
	d.RLock()
	pos, ok := d.posMap[key]
	d.RUnlock()
	return pos, ok
}

func (d *Dictionary) LoadDictionary(fileName string) error {
	return d.loadDictionary(fileName, false)
}

func (d *Dictionary) LoadUserDictionary(fileName string) error {
	return d.loadDictionary(fileName, true)
}

func (d *Dictionary) loadDictionary(fileName string, isUserDictionary bool) error {
	filePath, err := dictPath(fileName)
	if err != nil {
		return err
	}
	dictFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer dictFile.Close()

	scanner := bufio.NewScanner(dictFile)
	var token Token
	var line string
	var fields []string

	d.Lock()
	defer d.Unlock()

	if !isUserDictionary && len(d.freqMap) > 0 {
		d.freqMap = make(map[string]float64)
		d.posMap = make(map[string]string)
		d.total = 0.0
		d.logTotal = 0.0
	}
	for scanner.Scan() {
		line = scanner.Text()
		fields = strings.Split(line, " ")
		token.text = strings.TrimSpace(strings.Replace(fields[0], "\ufeff", "", 1))
		if length := len(fields); length > 1 {
			token.frequency, err = strconv.ParseFloat(fields[1], 64)
			if err != nil {
				return err
			}
			if length > 2 {
				token.pos = strings.TrimSpace(fields[2])
			}
		}
		d.addToken(token)
	}
	d.updateLogTotal()
	if err = scanner.Err(); err != nil {
		return err
	}
	return nil
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
