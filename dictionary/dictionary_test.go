package dictionary

import (
	"math"
	"testing"
)

var d *Dictionary

func init() {
	d = New()
}

func TestLoadDictionary(t *testing.T) {
	if err := d.LoadDictionary("../dict.txt"); err != nil {
		t.Fatalf("Failed to load dict.txt, err = %s", err)
	}
	n := len(d.freqMap)

	d.LoadDictionary("../foobar.txt")
	if len(d.freqMap) == n {
		t.Fatalf("Failed to load foobar.txt")
	}
}

func TestLoadUserDictionary(t *testing.T) {
	err := d.LoadUserDictionary("../userdict.txt")
	if err != nil {
		t.Fatalf("Failed to load userdict.txt, err = %s", err)
	}
	if f, _ := d.Frequency("八一双鹿"); f != 3.0 {
		t.Fatalf("Wrong frequency for word \"八一双鹿\", expect 3.0, got %f", f)
	}
}

func TestTotal(t *testing.T) {
	d.LoadDictionary("../userdict.txt")

	if d.Total() != 319.0 {
		t.Fatalf("Wrong total for userdict.txt, expect 319.0, got %f", d.Total())
	}
	if d.LogTotal() != math.Log(319.0) {
		t.Fatalf("Wrong total for userdict.txt, expect %f, got %f", math.Log(319.0), d.LogTotal())
	}
}
