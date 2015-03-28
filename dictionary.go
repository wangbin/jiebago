package jiebago

type Pair struct {
	Word string
	Flag string
}

type Entry struct {
	*Pair
	Freq float64
}

func NewEntry() *Entry {
	return &Entry{new(Pair), 0.0}
}

type DictLoader interface {
	AddEntry(*Entry)
}

type Cacher interface {
	CacheNameFormat() string
}
