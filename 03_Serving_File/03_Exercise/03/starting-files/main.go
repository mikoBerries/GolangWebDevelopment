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

	fs := http.FileServer(http.Dir("./public")) //serve static dir file
	http.Handle("/pics/", fs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func mainPages(w http.ResponseWriter, req *http.Request) {
	err := myTemplate.Execute(w, "index.gohtml")
	if err != nil {
		log.Panic("err execute:", err)
	}
}
