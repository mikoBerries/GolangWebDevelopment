//serving all files
package main

import (
	"log"
	"net/http"
)

func main() {

	//code 1
	// http.Handle("/", http.FileServer(http.Dir(".")))
	// log.Fatal(http.ListenAndServe(":8080", nil))

	//OR
	//code 2

	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))

}
