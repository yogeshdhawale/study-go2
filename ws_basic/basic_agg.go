package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type NewsAggEg struct {
	Title string
	News  string
}

func index_handler_agg_eg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>Text is here ...</h1>
	<p>This is para</p>
	`)
}
func newsAggHandlerEg(w http.ResponseWriter, r *http.Request) {

	p := NewsAggEg{Title: "1st News", News: "News text"}
	t, _ := template.ParseFiles("template_basic_agg.html")
	err := t.Execute(w, p)
	if err != nil {
		fmt.Println(err)
	}
}

func basic_agg_eg() {
	http.HandleFunc("/", index_handler_agg_eg)
	http.HandleFunc("/agg/", newsAggHandlerEg)
	http.ListenAndServe(":8000", nil)
}
