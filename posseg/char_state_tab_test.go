package posseg

import (
	"testing"
)

func TestGet(t *testing.T) {
	result := charStateTab.get('\u8000')
	if len(result) != 17 {
		t.FailNow()
	}
	result = charStateTab.get('\uaaaa')
	if len(result) == 17 {
		t.FailNow()
	}
}
