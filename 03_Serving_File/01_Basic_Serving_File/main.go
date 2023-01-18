/*basic Serving file in server
 */
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/notServingFromServer", http.HandlerFunc(notServingFromServer))
	http.Handle("/servingFromServer", http.HandlerFunc(servingFromServer))
	http.Handle("/servingFromServer2", http.HandlerFunc(servingFromServer2))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
func notServingFromServer(w http.ResponseWriter, req *http.Request) {
	body := `
	<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, body)
}
func servingFromServer(w http.ResponseWriter, req *http.Request) {
	body := `
	<!--image doesn't serve-->
	<img src="/toby.jpg">
	`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, body)
}

func servingFromServer2(w http.ResponseWriter, req *http.Request) {
	// body := ``
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 400)
		return
	}
	defer f.Close()

	//io. copy from file to writter
	io.Copy(w, f)
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")

}
