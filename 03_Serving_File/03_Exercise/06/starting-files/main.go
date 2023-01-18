package main

import (
	"html/template"
	"log"
	"net/http"
)

var myTemplate *template.Template

func init() {
	myTemplate = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", mainPages)
	http.HandleFunc("/about", about)
	http.HandleFunc("/apply", apply)
	http.HandleFunc("/applyprocess", apply)
	http.HandleFunc("/contact", contact)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func mainPages(w http.ResponseWriter, req *http.Request) {
	err := myTemplate.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)

}
func about(w http.ResponseWriter, req *http.Request) {
	err := myTemplate.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}
func contact(w http.ResponseWriter, req *http.Request) {
	err := myTemplate.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}
func apply(w http.ResponseWriter, req *http.Request) {
	var err error
	if req.Method == http.MethodPost {
		err = myTemplate.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	} else {
		err = myTemplate.ExecuteTemplate(w, "apply.gohtml", nil)
	}
	HandleError(w, err)
}
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
