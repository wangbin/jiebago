package cache

type Cacher interface {
	SetDict(string) error
	LoadUserDict(string) error
	AddWord(string, string, float64)
	Get(string) (float64, bool)
	Total() float64
}
