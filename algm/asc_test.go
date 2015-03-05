package algm

import (
	"testing"
)

func TestAscFilt(t *testing.T) {
	a := AscFilter{}
	if PASS == a.Filt("879") {
		t.Errorf("879 is not asc,so it should not Pass the asc filt.")
	}

	if FAIL == a.Filt("789") {
		t.Errorf("789 is asc,so it should Pass the asc filt.")
	}
}
