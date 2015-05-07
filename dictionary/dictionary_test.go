package dictionary

import (
	"sync"
	"testing"
)

type Dict struct {
	freqMap map[string]float64
	posMap  map[string]string
	sync.RWMutex
}

func (d *Dict) Load(ch <-chan Token) {
	d.Lock()
	for token := range ch {
		d.freqMap[token.Text()] = token.Frequency()
		if len(token.Pos()) > 0 {
			d.posMap[token.Text()] = token.Pos()
		}
	}
	d.Unlock()
}

func (d *Dict) AddToken(token Token) {
	d.Lock()
	d.freqMap[token.Text()] = token.Frequency()
	if len(token.Pos()) > 0 {
		d.posMap[token.Text()] = token.Pos()
	}
	d.Unlock()
}

func TestLoadDictionary(t *testing.T) {
	d := &Dict{freqMap: make(map[string]float64), posMap: make(map[string]string)}
	err := LoadDictionary(d, "../userdict.txt")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(d.freqMap) != 7 {
		t.Fatalf("Failed to load userdict.txt, got %d tokens with frequency, expected 7",
			len(d.freqMap))
	}
	if len(d.posMap) != 6 {
		t.Fatalf("Failed to load userdict.txt, got %d tokens with pos, expected 6", len(d.posMap))
	}
}

func TestAddToken(t *testing.T) {
	d := &Dict{freqMap: make(map[string]float64), posMap: make(map[string]string)}
	LoadDictionary(d, "../userdict.txt")
	d.AddToken(Token{"好用", 99, "a"})
	if d.freqMap["好用"] != 99 {
		t.Fatalf("Failed to add token, got frequency %f, expected 99", d.freqMap["好用"])
	}
	if d.posMap["好用"] != "a" {
		t.Fatalf("Failed to add token, got pos %s, expected \"a\"", d.posMap["好用"])
	}
}
