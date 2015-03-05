package algm

import (
	"testing"
)

func TestKdFilt(t *testing.T) {
	a := KdFilter{"1"}
	if PASS == a.Filt("879") {
		t.Errorf("kd filt judge error1.")
	}

	if FAIL == a.Filt("898") {
		t.Errorf("kd filt judge error2.")
	}
}
