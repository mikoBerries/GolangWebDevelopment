/*Serving file using http.serveContent() & HTTP.servefile() &http.Dir(path string)
 */
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//serving all file from http.Dir(path string) ()directory
	//including code .go
	http.Handle("/", http.FileServer(http.Dir(".")))
	//StripPrefix replacing /content and serving to specify dir
	http.Handle("/content/", http.StripPrefix("/content", http.FileServer(http.Dir("../assets"))))

	http.Handle("/cat", http.HandlerFunc(cat))
	http.Handle("/dog", http.HandlerFunc(dog))

	http.Handle("/pict/", http.HandlerFunc(d))
	http.Handle("/toby.jpg", http.HandlerFunc(dogpict))
	http.Handle("/dogpict", http.HandlerFunc(dp))

	//http.NotFoundHandler() for unsuported path/file will return 404
	http.Handle("/favicon.ico", http.NotFoundHandler())

	/*since http.ListenAndServe returning Err we can use log.fatal*/
	log.Fatal(http.ListenAndServe(":8080", nil))

	/*
		for making a static sever using this line of code
		http.Dir(path string) will not displaying special file "index.html"
	*/
	// http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))

	//using http.Error(w, "file not found", http.StatusNotFound) func as err massage
	// http.Error func will write us standart header and status code to given"writer" http.ResponseWriter
}

//for http.StripPrefix
func cat(w http.ResponseWriter, req *http.Request) {
	body := `
	<!--image serve using StripPrefix-->
	<img src="/content/cat.jpg">
	`
	//content/cat.jpg  =>>  asset/cat.jpg
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, body)
}

//for http.StripPrefix
func dog(w http.ResponseWriter, req *http.Request) {
	body := `
	<!--image serve using StripPrefix-->
	<img src="/content/dog.jpg">
	`
	// /content/dog.jpg  =>>  /asset/dog.jpg
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, body)
}

func d(w http.ResponseWriter, req *http.Request) {
	body := `
	<img src="../toby.jpg">
	`
	//get from /toby.jpg handler
	//must have http.Handle("/toby.jpg", http.HandlerFunc(dogpict))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, body)
}

func dogpict(w http.ResponseWriter, req *http.Request) {
	//writer request string path to pict
	http.ServeFile(w, req, "toby.jpg")
}

func dp(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer f.Close()
	//f.retruning file info struct , err
	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	//writer,request,file name,file modification time,file
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
