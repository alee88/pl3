package algm

import (
	"strconv"
	"strings"
)

type MinFilter struct {
	MinSet string
}

func (f MinFilter) Filt(s string) bool {
	if f.MinSet == "" {
		f.MinSet = allCond
	}
	min, _ := strconv.Atoi(string(s[0]))
	for i, _ := range s {
		t, _ := strconv.Atoi(string(s[i]))
		if min > t {
			min = t
		}
	}

	if strings.Contains(f.MinSet, strconv.Itoa(min)) {
		return PASS
	}

	return FAIL
}
