//undestanding of basic mux server and routing
package main

import (
	"fmt"
	"net/http"
)

type hotdog int
type cat int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "this is route for hotdog")
}
func (c cat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "this is route for hotdog")
}

func main() {
	//few way to handling routing using mux /default mux
	MakeServerUsingMuxHandler()
	MakeServerUsingDefaultMux()
	MakeServerUsingHandleFunc()
	MakeServerUsingHandle()
}

func MakeServerUsingHandle() {
	// http.Handle consume (string,http.Handler)
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	//type casting "d" and "c" to type http.HandlerFunc
	//because ==> type HandlerFunc func(ResponseWriter, *Request)
	//and have==> func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

	http.ListenAndServe(":8080", nil) //use nil so it'll use default
}

func MakeServerUsingHandleFunc() {
	// http.HandleFunc consume (string,func (res http.ResponseWriter, req *http.Request))
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)
	http.ListenAndServe(":8080", nil) //use nil so it'll use default
}

func MakeServerUsingDefaultMux() {
	var c cat
	var d hotdog
	// path : /dog/file/thisdog.html
	http.Handle("/dog/", d)
	//path : /cat
	http.Handle("/cat", c)
	http.ListenAndServe(":8080", nil) //use nil so it'll use default
}

func MakeServerUsingMuxHandler() {
	//dog and car are http.Handler
	var c cat
	var d hotdog
	//call new *mux
	mux := http.NewServeMux()
	// path : /dog/file/thisdog.html
	mux.Handle("/dog/", d)
	//path : /cat
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", mux)
}

func d(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "this is route for hotdog")
}
func c(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "this is route for hotdog")
}
