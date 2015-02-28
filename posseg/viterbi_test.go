package posseg

import (
	"testing"
)

var (
	route1 = []stateTag{
		stateTag{'B', "nr"},
		stateTag{'M', "nr"},
		stateTag{'E', "nr"},
		stateTag{'S', "v"},
		stateTag{'B', "v"},
		stateTag{'E', "v"},
		stateTag{'B', "n"},
		stateTag{'M', "n"},
		stateTag{'E', "n"},
		stateTag{'S', "d"},
		stateTag{'S', "v"},
		stateTag{'S', "n"},
		stateTag{'B', "v"},
		stateTag{'E', "v"},
		stateTag{'B', "nr"},
		stateTag{'M', "nr"},
		stateTag{'M', "nr"},
		stateTag{'M', "nr"},
		stateTag{'E', "nr"},
		stateTag{'S', "zg"}}
)

func TestViterbi(t *testing.T) {
	ss := "李小福是创新办主任也是云计算方面的专家;"
	prob, route := viterbi([]rune(ss))
	if prob != MinFloat {
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
