package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "<h1>Text is here ...</h1>")
	//fmt.Fprintf(w, "<p>This is para</p>")
	//fmt.Fprintf(w, "<p> ... and next para</p>")
	//fmt.Fprintf(w, "<p>data: %s</p>", "<strong>variables</strong>")

	fmt.Fprintf(w, `
	<h1>Text is here ...</h1>
	<p>This is para</p>
	`)
}
func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here is about me ...")
}

func basic_http() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)

	fmt.Println("Look at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
