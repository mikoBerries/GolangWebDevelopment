/*
Redirecting client
	StatusMultipleChoices  = 300 // RFC 9110, 15.4.1

	StatusMovedPermanently = 301 // RFC 9110, 15.4.2
	301 - Will informing to client url are moved
	StatusFound            = 302 // RFC 9110, 15.4.3

	StatusSeeOther         = 303 // RFC 9110, 15.4.4
	303 - changes method to get (always GET)

	StatusNotModified      = 304 // RFC 9110, 15.4.5
	StatusUseProxy         = 305 // RFC 9110, 15.4.6

	StatusTemporaryRedirect = 307 // RFC 9110, 15.4.8
	307 - keeps same method
	StatusPermanentRedirect = 308 // RFC 9110, 15.4.9

	common user is 302 & 307  status code

	// RFC 7231: Hypertext Transfer Protocol (HTTP/1.1)
	https://datatracker.ietf.org/doc/html/rfc7231
*/
/*
Passing value using GET url / POST From Method

3 encoding method
● application/x-www-form-urlencoded (the default)
● multipart/form-data ==> for file type
● text/plain == for debuging

*/package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var myTemplate *template.Template

func init() {
	myTemplate = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "this is your method at foo :"+req.Method)
	fmt.Println("this is your method at foo :" + req.Method)
	myTemplate.ExecuteTemplate(w, "index.gohtml", nil)
}

func bar(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "this is your method at bar :"+req.Method)
	fmt.Println("this is your method at bar :" + req.Method)
	if req.Method == http.MethodPost {
		w.Header().Add("Location", "/") //redirecting to "/" location
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}

func barred(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "this is your method at barred :"+req.Method)

	fmt.Println("this is your method at barred :" + req.Method)
	//http.redirect are func to writing redirect header
	http.Redirect(w, req, "/", http.StatusSeeOther)
	//http.StatusSeeOther make method POST "/barred" become method GET "/"
	myTemplate.ExecuteTemplate(w, "index.gohtml", nil)
}
