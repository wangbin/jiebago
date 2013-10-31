package finalseg

import (
	"math"
	"testing"
)

func TestViterbi(t *testing.T) {
	obs := "我们是程序员"
	states := []byte{'B', 'M', 'E', 'S'}
	prob, path := viterbi([]rune(obs), states)
	if math.Abs(prob+39.68824128493802) > 1e-10 {
		t.Error(prob)
	}
	for index, state := range []byte{'B', 'E', 'S', 'B', 'M', 'E'} {
		if path[index] != state {
			t.Error(path)
		}
	}
}

func TestCutHan(t *testing.T) {
	obs := "我们是程序员"
	result := cutHan(obs)
	if len(result) != 3 {
		t.Error(result)
	}
	if result[0] != "我们" {
		t.Error(result[0])
	}
	if result[1] != "是" {
		t.Error(result[1])
	}
	if result[2] != "程序员" {
		t.Error(result[2])
	}
}

func TestCut(t *testing.T) {
	sentence := "我们是程序员"
	result := Cut(sentence)
	if len(result) != 3 {
		t.Error(len(result))
	}
	if result[0] != "我们" {
		t.Error(result[0])
	}
	if result[1] != "是" {
		t.Error(result[1])
	}
	if result[2] != "程序员" {
		t.Error(result[2])
	}
	result2 := Cut("I'm a programmer!")
	if len(result2) != 8 {
		t.Error(result2)
	}
	result3 := Cut("程序员average年龄28.6岁。")
	if len(result3) != 6 {
		t.Error(result3)
	}

}
