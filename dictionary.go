package jiebago

type Entry struct {
	Word string
	Flag string
	Freq float64
}

type DictLoader interface {
	AddEntry(Entry)
}

type Cacher interface {
	CacheNameFormat() string
}
