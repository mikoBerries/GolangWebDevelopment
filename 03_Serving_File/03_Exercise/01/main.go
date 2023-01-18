//Exercise read "Readme.md"
package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var myTemplate *template.Template

func main() {
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/dog.jpg", http.HandlerFunc(dogpict))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}
func dog(w http.ResponseWriter, req *http.Request) {
	myTemplate, err := myTemplate.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatal("err parseFiles :", err)
	}
	myTemplate.ExecuteTemplate(w, "dog.gohtml", nil)
}
func dogpict(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}
