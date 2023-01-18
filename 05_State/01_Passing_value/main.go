/*
Passing value using GET url / POST From Method
*/package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var myTemplate *template.Template

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func init() {
	myTemplate = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/main", mainPages)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainPages(w http.ResponseWriter, req *http.Request) {
	q := req.FormValue("q")

	// for post method
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`+q)

	//for get method
	// io.WriteString(w, `
	// <form method="GET">
	//  <input type="text" name="q">
	//  <input type="submit">
	// </form>
	// <br>`+q)
}

func foo(w http.ResponseWriter, req *http.Request) {
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"
	err := myTemplate.ExecuteTemplate(w, "index.gohtml", person{f, l, s})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
