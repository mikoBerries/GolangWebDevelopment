/*leraning about basic fundamental of packakge net/http
and function serveHTTP http.Request
*/
package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

//ServeHTTP(w http.ResponseWriter, req *http.Request) interface of http.Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	//store data request to struct
	data := struct {
		Method        string
		Submissions   url.Values
		URL           *url.URL
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method, //method (GET/POST/PUT)
		req.Form,   //value from map[string][]string
		req.URL,    //get url from request (/path/index.html)
		req.Header, //get Header from  incoming request (life line)
		req.Host,
		req.ContentLength, //length of body content in byte
	}
	//execute template with data request
	err = indexTemplate.ExecuteTemplate(w, "index.gohtml", data)
	if err != nil {
		log.Panic("err ExecuteTemplate:", err)
	}
}

var indexTemplate *template.Template

func init() {
	indexTemplate = template.Must(template.ParseFiles("./Resource_Template/index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
