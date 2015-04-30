package util

import (
	"regexp"
	"testing"
)

func TestRegexpSplit(t *testing.T) {
	result := RegexpSplit(regexp.MustCompile(`\p{Han}+`),
		"BP神经网络如何训练才能在分类时增加区分度？", -1)
	if len(result) != 2 {
		t.Fatal(result)
	}
	result = RegexpSplit(regexp.MustCompile(`(\p{Han})+`),
		"BP神经网络如何训练才能在分类时增加区分度？", -1)
	if len(result) != 3 {
		t.Fatal(result)
	}
	result = RegexpSplit(regexp.MustCompile(`([\p{Han}#]+)`),
		",BP神经网络如何训练才能在分类时#增加区分度？", -1)
	if len(result) != 3 {
		t.Fatal(result)
	}
}
