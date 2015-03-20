package cache

type FileBasedCache struct {
	freq                map[[]rune]float64
	total               float64
	dictPath, cachePath string
}

func (f *FileBasedCache) SetDict(dictFileName string) error {
	if len(f.freq) > 0 {
		f.freq = make(map[[]rune]flot64)
	}

	if err := f.setFilePath(dictFileName); err != nil {
		return err
	}

	cached, err := f.cached()
	if err != nil {
		return err
	}

	if cached {
		if err := f.load(cacheFilePath); err == nil {
			return nil
		} // TODO: logging?
	}

	if err := f.read(dictFilePath); err != nil {
		return err
	}

	if err := f.dump(cacheFilePath); err != nil {
		return err
	}
}

func (f *FileBasedCache) LoadUserDict(userDictFileName string) error {
	return nil
}

func (f *FileBasedCache) AddWord(word, tag string, freq float64) {

}

func (f *FileBasedCache) Get(key string) (float64, bool) {
	val, ok := f.freq[[]rune(key)]
	return val, ok
}

func (f *FileBasedCache) Total() float64 {
	return f.total
}

func (f *FileBasedCache) setFilePath() (err error) { // TODO: specify the temp dir
	f.dictFilePath, err = DictPath(dictFileName)
	if err != nil {
		return
	}
	f.cacheFilePath = filepath.Join(os.TempDir(),
		fmt.Sprintf("jieba.%x.cache", md5.Sum([]byte(f.dictFilePath))))
	return
}

func (f *FileBasedCache) cached() (bool, error) {
	dictFileInfo, err := os.Stat(f.dictFilePath) // TODO: logging
	if err != nil {
		return false, err
	}

	cacheFileInfo, err := os.Stat(cacheFilePath)
	if err != nil { // TODO: logging
		return false, nil
	}

	return cacheFileInfo.ModTime().After(dictFileInfo.ModTime()), nil
}

func (f *FileBasedCache) load(cacheFilePath string) error {
	cacheFile, err := os.Open(cacheFilePath)
	if err != nil {
		return err
	}
	defer cacheFile.Close()

	dec := gob.NewDecoder(cacheFile)
	return dec.Decode(&f)
}

func (f *FileBasedCache) read(dictFilePath string) error {
	wtfs, err := ParseDictFile(dictFilePath)
	if err != nil {
		return err
	}

	for _, wtf := range wtfs {
		t.addWord(wtf) // TODO: add word, ignore frequency
	}

	return nil
}

func (f *FileBasedCache) dump(cacheFilePath string) error {
	cacheFile, err := os.OpenFile(cacheFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer cacheFile.Close()

	enc := gob.NewEncoder(cacheFile)
	if err := enc.Encode(f); err != nil {
		return err
	}
	return nil
}
