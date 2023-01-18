/*third party github julienschmidt - httprouter lib
ps httprouter.Params are map[]params value from request
strcut params => key string, value string ( just like map[key]value )
*/package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var MyTemplate *template.Template

//init - Parsing all template
func init() {
	MyTemplate = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/about", about)
	router.GET("/contact", contact)
	router.GET("/user/:name", user)
	router.GET("/apply", apply)
	router.POST("/apply", applyProcess)
	router.GET("/blog/:category/:article", blogRead)
	router.POST("/blog/:category/:article", blogWrite)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := MyTemplate.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}
func about(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := MyTemplate.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := MyTemplate.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}
func apply(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := MyTemplate.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := MyTemplate.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}
func user(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello , %s \n", ps.ByName("user"))
}

func blogRead(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "READ :Category. %s \n", ps.ByName("category"))
	fmt.Fprintf(w, "READ :Article. %s \n", ps.ByName("article"))
}
func blogWrite(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "WRITE :Category. %s \n", ps.ByName("category"))
	fmt.Fprintf(w, "WRITE :Article. %s \n", ps.ByName("article"))
}

//handling err in every handler func()
func HandleError(w http.ResponseWriter, e error) {
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		log.Println(e)
	}
}
