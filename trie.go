package jiebago

type Jieba struct {
	Total float64
	Freq  map[string]float64
}

func (j *Jieba) AddEntry(entry *Entry) {
	j.Add(entry.Word, entry.Freq)
}

func (j *Jieba) Add(word string, freq float64) {
	j.Freq[word] = freq
	j.Total += freq
	runes := []rune(word)
	for i := 0; i < len(runes); i++ {
		frag := string(runes[0 : i+1])
		if _, ok := j.Freq[frag]; !ok {
			j.Freq[frag] = 0.0
		}
	}
}

// Load user specified dictionary file.
func (j *Jieba) LoadUserDict(dictFilePath string) error {
	return LoadDict(j, dictFilePath, false)
}

// Set the dictionary, could be absolute path of dictionary file, or dictionary
// name in current directory. This function must be called before cut any
// sentence.
func NewJieba(dictFileName string) (*Jieba, error) {
	j := &Jieba{Total: 0.0, Freq: make(map[string]float64)}
	err := SetDict(j, dictFileName, false)
	return j, err
}
