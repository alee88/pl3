package algm

import (
	"testing"
)

func TestMinFilt(t *testing.T) {
	a := MinFilter{"6"}
	if PASS == a.Filt("879") {
		t.Errorf("min filt judge error1.")
	}

	if FAIL == a.Filt("786") {
		t.Errorf("min filt judge error2.")
	}
}
