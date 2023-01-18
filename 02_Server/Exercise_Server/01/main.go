//Hand on Exercise 01 (Readme.md)
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var myTemplate *template.Template

type User struct {
	Name string
}

func (u User) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, u.Name)
}
func main() {
	human := User{"Agus Budi"}
	http.Handle("/", http.HandlerFunc(i))
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/me/", human)
	http.Handle("/something/", http.HandlerFunc(something))
	http.ListenAndServe(":8080", nil)
}

func d(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "this is dog page")
}

func i(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "this is index page")
}

func something(w http.ResponseWriter, req *http.Request) {
	myTemplate, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}
	humans := []User{
		User{"User 1"},
		User{"User 2"},
		User{"User 3"},
	}
	myTemplate.ExecuteTemplate(w, "something.gohtml", humans)
	if err != nil {
		log.Fatalln("error Execute template ", err)
	}
}
