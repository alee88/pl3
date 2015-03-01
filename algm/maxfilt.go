package algm

import (
	"strconv"
	"strings"
)

type MaxFilter struct {
	MaxSet string
}

func (f MaxFilter) Filt(s string) bool {
	if f.MaxSet == "" {
		f.MaxSet = allCond
	}
	max, _ := strconv.Atoi(string(s[0]))
	for i, _ := range s {
		t, _ := strconv.Atoi(string(s[i]))
		if max < t {
			max = t
		}
	}

	if strings.Contains(f.MaxSet, strconv.Itoa(max)) {
		return PASS
	}

	return FAIL
}
