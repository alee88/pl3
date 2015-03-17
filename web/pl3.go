// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"strconv"
	"strings"

	"pl3/algm"
)

var (
	port = flag.Int("port", 80, "set server port.")
)

type Page struct {
	Title string
	Tips  string
	Total string
	Body  []byte
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "首页"}
	renderTemplate(w, "index", p)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	kd := r.FormValue("kd")
	hw := r.FormValue("hw")
	max := r.FormValue("max")
	min := r.FormValue("min")
	yh := r.FormValue("yh")
	lj := r.FormValue("lj")

	filters := make([]algm.Filter, 0)
	filters = append(filters,
		algm.AscFilter{},
		algm.KdFilter{KdSet: kd},
		algm.HwFilter{HwSet: hw},
		algm.MaxFilter{MaxSet: max},
		algm.MinFilter{MinSet: min},
		algm.YhFilter{YhSet: yh},
		algm.LjFilter{LjSet: lj})

	cands := make([]string, 0)
	for i := 0; i < 1000; i++ {
		s := fmt.Sprintf("%03d", i)
		cands = append(cands, s)
	}
	rst := algm.Filt(filters, cands)

	tips := fmt.Sprintf("跨度[%s],和尾[%s],最大值[%s],最小值[%s],余数和[%s],数字累加值[%s]\n",
		kd, hw, max, min, yh, lj)
	num := fmt.Sprintf("满足条件的组合共有：%d 组", len(rst))

	p := &Page{Title: "结果", Tips: tips, Total: num, Body: []byte(strings.Join(rst, ";"))}
	renderTemplate(w, "result", p)
}

var templates = template.Must(template.ParseFiles("index.html", "result.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/result", resultHandler)

	if *port <= 0 || *port > 65535 {
		fmt.Printf("invalid port:%d!", *port)
		return
	}
	p := strconv.Itoa(*port)

	conn, err := net.Dial("udp", "www.google.com.hk:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	addr := strings.Split(conn.LocalAddr().String(), ":")[0] + ":" + p
	http.ListenAndServe(addr, nil)
}
