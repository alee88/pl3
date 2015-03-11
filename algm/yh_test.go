package algm

import (
	"testing"
)

func TestYhFilt(t *testing.T) {
	a := YhFilter{"6"}
	if PASS == a.Filt("879") {
		t.Errorf("yh filt judge error1.")
	}

	if FAIL == a.Filt("258") {
		t.Errorf("yh filt judge error2.")
	}
}
