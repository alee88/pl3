// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"pl3/algm"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

type Page struct {
	Title string
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

	filters := make([]algm.Filter, 0)
	filters = append(filters,
		algm.AscFilter{},
		algm.KdFilter{KdSet: kd},
		algm.HwFilter{HwSet: hw},
		algm.MaxFilter{MaxSet: max},
		algm.MinFilter{MinSet: min})

	cands := make([]string, 0)
	for i := 0; i < 1000; i++ {
		s := fmt.Sprintf("%03d", i)
		cands = append(cands, s)
	}
	rst := algm.Filt(filters, cands)

	p := &Page{Title: "结果", Body: []byte(strings.Join(rst, ";"))}
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

	if *addr {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}

	http.ListenAndServe(":8080", nil)
}
