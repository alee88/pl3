package algm

import (
	"testing"
)

func TestLjFilt(t *testing.T) {
	a := LjFilter{"15"}
	if PASS == a.Filt("879") {
		t.Errorf("lj filt judge error1.")
	}

	if FAIL == a.Filt("257") {
		t.Errorf("lj filt judge error2.")
	}
	
	if PASS == a.Filt("254") {
		t.Errorf("lj filt judge error3.")
	}
	
	if FAIL == a.Filt("253") {
		t.Errorf("lj filt judge error2.")
	}
}
