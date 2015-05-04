package posseg

import (
	"math"
	"sync"

	"github.com/wangbin/jiebago/dictionary"
)

type Dictionary struct {
	total, logTotal float64
	freqMap         map[string]float64
	posMap          map[string]string
	sync.RWMutex
}

func (d *Dictionary) Load(ch <-chan dictionary.Token) {
	d.Lock()
	for token := range ch {
		d.addToken(token)
	}
	d.Unlock()
	d.updateLogTotal()
}

func (d *Dictionary) AddToken(token dictionary.Token) {
	d.Lock()
	d.addToken(token)
	d.Unlock()
	d.updateLogTotal()
}

func (d *Dictionary) addToken(token dictionary.Token) {
	d.freqMap[token.Text()] = token.Frequency()
	d.total += token.Frequency()
	runes := []rune(token.Text())
	n := len(runes)
	for i := 0; i < n; i++ {
		frag := string(runes[:i+1])
		if _, ok := d.freqMap[frag]; !ok {
			d.freqMap[frag] = 0.0
		}
	}
	if len(token.Pos()) > 0 {
		d.posMap[token.Text()] = token.Pos()
	}
}

func (d *Dictionary) updateLogTotal() {
	d.logTotal = math.Log(d.total)
}

func (d *Dictionary) Frequency(key string) (float64, bool) {
	d.RLock()
	freq, ok := d.freqMap[key]
	d.RUnlock()
	return freq, ok
}

func (d *Dictionary) Pos(key string) (string, bool) {
	d.RLock()
	pos, ok := d.posMap[key]
	d.RUnlock()
	return pos, ok
}

func (d *Dictionary) loadDictionary(fileName string) error {
	return dictionary.LoadDictionary(d, fileName)
}
