package algm

import (
	"testing"
)

func TestHwFilt(t *testing.T) {
	a := HwFilter{"2"}
	if PASS == a.Filt("012") {
		t.Errorf("hw filt judge error1.")
	}

	if FAIL == a.Filt("011") {
		t.Errorf("hw filt judge error2.")
	}
}
