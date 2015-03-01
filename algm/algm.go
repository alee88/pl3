package algm

import (
	"sort"
)

const allCond = "0123456789"

const PASS = true

const FAIL = false

type Filter interface {
	Filt(s string) bool
}

func Filt(seq []Filter, candSet []string) []string {
	rst := make([]string, 0, 1000)
	pass := true
	for _, c := range candSet {
		pass = true
		for _, f := range seq {
			if f.Filt(c) != PASS {
				pass = false
				break
			}
		}
		if pass {
			rst = append(rst, c)
		}
	}

	return rst
}

type SortRunes []rune

func (s SortRunes) Len() int {
	return len(s)
}

func (s SortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(SortRunes(r))
	return string(r)
}
