package main

import (
	"fmt"

	"pl3/algm"
)

func main() {
	cands := make([]string, 0)
	for i := 0; i < 1000; i++ {
		s := fmt.Sprintf("%03d", i)
		cands = append(cands, s)
	}

	var kd, hw, max, min string

	fmt.Println("请输入所有可能的跨度：")
	fmt.Scanf("%s", &kd)
	fmt.Println("请输入所有可能的和尾：")
	fmt.Scanf("%s", &hw)
	fmt.Println("请输入所有可能的最大值：")
	fmt.Scanf("%s", &max)
	fmt.Println("请输入所有可能的最小值：")
	fmt.Scanf("%s", &min)

	fmt.Print("您输入的条件为：")
	fmt.Printf("跨度[%s]；和尾[%s]；最大值[%s]；最小值[%s]\n", kd, hw, max, min)

	filters := make([]algm.Filter, 0)
	filters = append(filters,
		algm.AscFilter{},
		algm.KdFilter{KdSet: kd},
		algm.HwFilter{HwSet: hw},
		algm.MaxFilter{MaxSet: max},
		algm.MinFilter{MinSet: min})

	rst := algm.Filt(filters, cands)
	fmt.Println("符合条件的组合共有：", len(rst), "组，分别为:", rst)
}
