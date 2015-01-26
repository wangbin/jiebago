package posseg

import (
	"testing"
)

var (
	route1 = []StateTag{
		StateTag{'B', "nr"},
		StateTag{'M', "nr"},
		StateTag{'E', "nr"},
		StateTag{'S', "v"},
		StateTag{'B', "v"},
		StateTag{'E', "v"},
		StateTag{'B', "n"},
		StateTag{'M', "n"},
		StateTag{'E', "n"},
		StateTag{'S', "d"},
		StateTag{'S', "v"},
		StateTag{'S', "n"},
		StateTag{'B', "v"},
		StateTag{'E', "v"},
		StateTag{'B', "nr"},
		StateTag{'M', "nr"},
		StateTag{'M', "nr"},
		StateTag{'M', "nr"},
		StateTag{'E', "nr"},
		StateTag{'S', "zg"}}
)

func TestViterbi(t *testing.T) {
	ss := "李小福是创新办主任也是云计算方面的专家;"
	prob, route := Viterbi([]rune(ss))
	if prob != MIN_FLOAT {
		t.Error(prob)
	}
	if len(route) != len(route1) {
		t.Error(len(route))
	}
	for index, _ := range route {
		if route[index] != route1[index] {
			t.Error(route[index])
		}
	}
}
