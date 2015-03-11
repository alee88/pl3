package algm

import (
	"strconv"
	"strings"
)

type LjFilter struct {
	LjSet string
}

func (f LjFilter) Filt(s string) bool {
	if f.LjSet == "" {
		f.LjSet = allCond
	}

	lj := 0
	for i, _ := range s {
		t, _ := strconv.Atoi(string(s[i]))
		lj += t
	}

	for lj > 9 {
		lj = lj/10 + lj%10
	}
	
	if strings.Contains(f.LjSet, strconv.Itoa(lj)) {
		return PASS
	}

	return FAIL
}
