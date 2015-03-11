package algm

import (
	"strconv"
	"strings"
)

const yhCond = "0123456"

type YhFilter struct {
	YhSet string
}

func (f YhFilter) Filt(s string) bool {
	if f.YhSet == "" {
		f.YhSet = yhCond
	}

	yh := 0
	for i, _ := range s {
		t, _ := strconv.Atoi(string(s[i]))
		yh += t % 3
	}

	if strings.Contains(f.YhSet, strconv.Itoa(yh)) {
		return PASS
	}

	return FAIL
}
