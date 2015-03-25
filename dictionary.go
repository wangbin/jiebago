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

type Loader interface {
	AddEntry(Entry)
	CachePath(string) string
}
