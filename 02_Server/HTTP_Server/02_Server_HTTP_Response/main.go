/*leraning about basic fundamental of packakge net/http
and function serveHTTP http.Request
*/
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type hotdog int

//ServeHTTP(w http.ResponseWriter, req *http.Request) interface of http.Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("myHeader", "this is some key etc")
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "<h1>this is anythign you want to put</h1>")
	//execute template with data request
	// err = indexTemplate.ExecuteTemplate(w, "index.gohtml", nil)
	// if err != nil {
	// 	log.Panic("err ExecuteTemplate:", err)
	// }
}

var indexTemplate *template.Template

func init() {
	// indexTemplate = template.Must(template.ParseFiles("./Resource_Template/index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
