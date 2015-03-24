package jiebago

type Pair struct {
	Word string
	Flag string
}

type Token struct {
	*Pair
	Freq float64
}

type DictLoader interface {
	Add(*Token)
}
