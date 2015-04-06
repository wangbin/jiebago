package posseg

import (
	"testing"
)

var defaultRoute []Tag

func init() {
	var t Tag
	t, _ = NewTag("B", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("M", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("E", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("S", "v")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("B", "v")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("E", "v")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("B", "n")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("M", "n")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("E", "n")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("S", "d")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("S", "v")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("S", "n")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("B", "v")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("E", "v")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("B", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("M", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("M", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("M", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("E", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = NewTag("S", "zg")
	defaultRoute = append(defaultRoute, t)
}

func TestViterbi(t *testing.T) {
	ss := "李小福是创新办主任也是云计算方面的专家;"
	route := viterbi([]rune(ss))
	if len(route) != len(defaultRoute) {
		t.Fatal(len(route))
	}
	for index := range route {
		if route[index] != defaultRoute[index] {
			t.Fatal(route[index])
		}
	}
}

func BenchmarkViterbi(b *testing.B) {
	ss := "李小福是创新办主任也是云计算方面的专家;"
	for i := 0; i < b.N; i++ {
		viterbi([]rune(ss))
	}
}
