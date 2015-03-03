package algm

import (
	"strconv"
	"strings"
)

type KdFilter struct {
	KdSet string
}

func (f KdFilter) Filt(s string) bool {
	if f.KdSet == "" {
		f.KdSet = allCond
	}

	max, _ := strconv.Atoi(string(s[0]))
	for i, _ := range s {
		t, _ := strconv.Atoi(string(s[i]))
		if max < t {
			max = t
		}
	}

	min, _ := strconv.Atoi(string(s[0]))
	for i, _ := range s {
		t, _ := strconv.Atoi(string(s[i]))
		if min > t {
			min = t
		}
	}

	if strings.Contains(f.KdSet, strconv.Itoa(max-min)) {
		return PASS
	}

	return FAIL
}
