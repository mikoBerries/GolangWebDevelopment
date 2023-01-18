/*
Passing value using GET url / POST From Method

3 encoding method
● application/x-www-form-urlencoded (the default)
● multipart/form-data ==> for file type
● text/plain == for debuging

*/package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var myTemplate *template.Template

func init() {
	myTemplate = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
	http.HandleFunc("/", foo)
	// http.HandleFunc("/main", mainPages)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	if req.Method == http.MethodPost {
		//post file type data
		//producing file,header,err
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("file:", f)
		fmt.Println("file header:", h)
		fmt.Println("err:", err)

		//read
		byte, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//get string from file
		s = string(byte)

		//storing to server
		nf, err := os.Create(filepath.Join("./Client_file/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer nf.Close()
		//write new file with byte
		_, err = nf.Write(byte)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// body := make([]byte, req.ContentLength)
	// _, err := req.Body.Read(body)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	log.Fatalln(err)
	// }
	// io.WriteString(w, string(body))

	// for post method
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := myTemplate.ExecuteTemplate(w, "index.gohtml", s)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
	// io.WriteString(w, `
	// <form action="/" method="POST" enctype="multipart/form-data">
	//  <input type="file" name="q">
	//  <input type="submit">
	// </form>
	// <br>`+s)

}
