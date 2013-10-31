package posseg

import (
	"testing"
)

func TestGet(t *testing.T) {
	result := CharStateTab.Get('\u8000')
	if len(result) != 17 {
		t.FailNow()
	}
	result = CharStateTab.Get('\uaaaa')
	if len(result) == 17 {
		t.FailNow()
	}
}
