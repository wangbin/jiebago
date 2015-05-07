package posseg

import (
	"testing"
)

var defaultRoute []tag

func init() {
	var t tag
	t, _ = newTag("B", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("M", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("E", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("S", "v")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("B", "v")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("E", "v")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("B", "n")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("M", "n")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("E", "n")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("S", "d")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("S", "v")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("S", "n")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("B", "v")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("E", "v")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("B", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("M", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("M", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("M", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("E", "nr")
	defaultRoute = append(defaultRoute, t)
	t, _ = newTag("S", "zg")
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
