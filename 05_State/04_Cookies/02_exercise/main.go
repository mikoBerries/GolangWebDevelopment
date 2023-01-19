package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie") // returning nil and new error if cookies not found
	var counter int
	if err != nil {
		counter = 1
	} else if c != nil {
		counter, err = strconv.Atoi(c.Value)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Fatalln(err.Error())
		}
		counter++
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: strconv.Itoa(counter),
		Path:  "/read",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}

// Using cookies, track how many times a user has been to your website domain.
