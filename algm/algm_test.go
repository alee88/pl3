package algm

import (
	"fmt"
	"testing"
)

func TestFilt(t *testing.T) {
	cands := make([]string, 0)
	for i := 0; i < 1000; i++ {
		s := fmt.Sprintf("%03d", i)
		cands = append(cands, s)
	}

	filters := make([]Filter, 0)
	filters = append(filters,
		AscFilter{},
		KdFilter{KdSet: "1"},
		HwFilter{HwSet: "2"},
		MaxFilter{MaxSet: "1"},
		MinFilter{MinSet: "0"})

	rst := Filt(filters, cands)

	if len(rst) != 1 || rst[0] != "011" {
		t.Errorf("result is :%s, should be 011", rst[0])
	}

}

func TestSortString(t *testing.T) {
	str := SortString("987")

	if str != "789" {
		t.Errorf("result is :%s, should be 789", str)
	}
}
