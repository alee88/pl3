package algm

import (
	"testing"
)

func TestMaxFilt(t *testing.T) {
	a := MaxFilter{"6"}
	if PASS == a.Filt("879") {
		t.Errorf("max filt judge error1.")
	}

	if PASS == a.Filt("869") {
		t.Errorf("max filt judge error2.")
	}

	if FAIL == a.Filt("456") {
		t.Errorf("max filt judge error3.")
	}

}
