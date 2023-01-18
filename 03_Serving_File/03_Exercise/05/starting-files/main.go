package main

import (
	"html/template"
	"log"
	"net/http"
)

var myTemplate *template.Template

func init() {
	myTemplate = template.Must(template.ParseFiles("./templates/index.gohtml"))
}
func main() {
	http.HandleFunc("/", mainPages)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func mainPages(w http.ResponseWriter, req *http.Request) {
	err := myTemplate.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Panic("err execute:", err)
	}
}
