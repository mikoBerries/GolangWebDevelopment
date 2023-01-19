/*
	Giving cookies to client side
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// var myTemplate *template.Template

// func init() {
// 	myTemplate = template.Must(template.ParseGlob("templates/*.gohtml"))
// }
func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/readCookies", readCookies)
	http.HandleFunc("/expire", expire)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	// Cookies.path "/" meaning universal cookies passing to all web path
	// Cookies.path "/set" meaning when cookies passed too "/set" called
	// Cookies.path creating specifik cookies for specifik path
	myCookie := http.Cookie{
		Name:  "Cookie-name",
		Value: "cookies-value",
		Path:  "/",
	}
	http.SetCookie(w, &myCookie)
	//setting multiple cookies
	myCookie = http.Cookie{
		Name:  "cookies2",
		Value: "10-123-1239123",
		Path:  "/",
	}
	http.SetCookie(w, &myCookie)
}

func readCookies(w http.ResponseWriter, req *http.Request) {
	//getting all cookies []
	//req.cookie(Cookie.Name) for secific cookies name
	cookies := req.Cookies()
	io.WriteString(w, "total cookies :"+strconv.Itoa(len(cookies))+"\n")
	//reading all cookies
	for _, val := range cookies {
		io.WriteString(w, val.Name+" : "+val.Value+"\n")
		fmt.Println(val.Name + " : " + val.Value + "\n")
	}
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(w, `<h1><a href="/read">read</a></h1>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	c.MaxAge = -1 // delete cookie
	http.SetCookie(w, c)
	http.Redirect(w, req, "/asdasd", http.StatusSeeOther)
}
