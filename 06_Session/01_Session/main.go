/*
	Undestanding Session
	common use id in session using UUID
*/
package main

import (
	"html/template"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))

}

func main() {
	// http.HandleFunc("/set", set)
	// http.HandleFunc("/readCookies", readCookies)
	// http.HandleFunc("/expire", expire)

	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {

	myCookie, err := req.Cookie("session") //finding cookies in header
	if err != nil {

		ID := uuid.NewV4()
		myCookie = &http.Cookie{
			Name:  "session",
			Value: ID.String(),
			// Secure: true,
			// HttpOnly: true,
			// Path:     "/",
		}
		http.SetCookie(w, myCookie)
	}
	// if the user exists already, get user
	var us user
	if un, ok := dbSessions[myCookie.Value]; ok { //search in dbsessions
		us = dbUsers[un] //pick user struct from Dbusers map

	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		us = user{un, f, l}
		dbSessions[myCookie.Value] = un
		dbUsers[un] = us
	}
	tpl.ExecuteTemplate(w, "index.gohtml", us)
}

func bar(w http.ResponseWriter, req *http.Request) {
	var us user
	//check cookie "session" in header
	sessionCookies, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//check map dbSession using cookie "session" value
	un, ok := dbSessions[sessionCookies.Value]
	if !ok { //search in dbsessions
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	us = dbUsers[un] //pick user struct from Dbusers map
	tpl.ExecuteTemplate(w, "bar.gohtml", us)

}

// func set(w http.ResponseWriter, req *http.Request) {
// 	http.SetCookie(w, &http.Cookie{
// 		Name:  "session",
// 		Value: "some value",
// 		Path:  "/",
// 	})
// 	fmt.Fprintln(w, `<h1><a href="/read">read</a></h1>`)
// }

// func read(w http.ResponseWriter, req *http.Request) {
// 	c, err := req.Cookie("session")
// 	if err != nil {
// 		http.Redirect(w, req, "/set", http.StatusSeeOther)
// 		return
// 	}

// 	fmt.Fprintf(w, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
// }

// func expire(w http.ResponseWriter, req *http.Request) {
// 	c, err := req.Cookie("session")
// 	if err != nil {
// 		http.Redirect(w, req, "/set", http.StatusSeeOther)
// 		return
// 	}

// 	// MaxAge=0 means no 'Max-Age' attribute specified.
// 	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
// 	// MaxAge>0 means Max-Age attribute present and given in seconds
// 	c.MaxAge = -1 // delete cookie
// 	http.SetCookie(w, c)
// 	http.Redirect(w, req, "/asdasd", http.StatusSeeOther)
// }
