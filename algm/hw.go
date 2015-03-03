package algm

import (
	"strconv"
	"strings"
)

type HwFilter struct {
	HwSet string
}

func (f HwFilter) Filt(s string) bool {
	if f.HwSet == "" {
		f.HwSet = allCond
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		t, _ := strconv.Atoi(string(s[i]))
		sum += t
	}
	ssum := strconv.Itoa(sum)

	if strings.Contains(f.HwSet, string(ssum[len(ssum)-1])) {
		return PASS
	}

	return FAIL
}
