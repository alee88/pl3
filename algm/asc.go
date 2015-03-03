package algm

import (
	"sort"
)

type AscFilter struct {
}

func (f AscFilter) Filt(s string) bool {
	r := []rune(s)
	if sort.IsSorted(SortRunes(r)) {
		return PASS
	}

	return FAIL
}
