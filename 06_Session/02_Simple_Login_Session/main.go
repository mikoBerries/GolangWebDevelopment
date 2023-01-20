/*
	Undestanding Session
	common use id in session using UUID
*/
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName string
	First    string
	Last     string
	Role     string
	Password []byte
}
type session struct {
	uname        string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = map[string]user{}       // user ID, user
var dbSessions = map[string]session{} // session ID, user ID
var dbSessionsCleaned time.Time       //last time cleaned (for go routine)

const sessionLeght = 600 //max age int second

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.gohtml"))

}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/logout", logout)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	us := getUser(w, req)

	tpl.ExecuteTemplate(w, "index.gohtml", us)
}

func bar(w http.ResponseWriter, req *http.Request) {
	us := getUser(w, req)
	if !alreadyLogIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if us.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", us)

}

func signup(w http.ResponseWriter, req *http.Request) {
	var us user
	if alreadyLogIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {

		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		p := req.FormValue("password")
		by, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Fatalln(err.Error())
		}

		// check duplicate username
		_, ok := dbUsers[un]
		if ok { // if found
			http.Error(w, "Duplicate name", http.StatusForbidden)
			return
		}

		//store to dbUsers ( un as key us as value)
		us = user{un, f, l, r, by}
		dbUsers[un] = us

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLogIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		err := req.ParseForm()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		un := req.FormValue("username")
		ps := req.FormValue("password")
		if p, ok := dbUsers[un]; ok {
			//bcrypt.CompareHashAndPassword returnin nil if equal and return Error if no equal
			err = bcrypt.CompareHashAndPassword(p.Password, []byte(ps))
			if err != nil {
				http.Redirect(w, req, "/login", http.StatusSeeOther)
				return
			}
			ID := uuid.NewV4()
			sessionCookie := &http.Cookie{
				Name:   "session",
				Value:  ID.String(),
				MaxAge: sessionLeght,
				// HttpOnly: true,
				// Path:     "/",
			}

			http.SetCookie(w, sessionCookie)
			fmt.Println("cookies", sessionCookie.String())
			// store to dbSessions ( UUID as key un as value)
			dbSessions[sessionCookie.Value] = session{un, time.Now()}
		}
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLogIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(dbSessions, c.Value)
	// remove the cookie
	c.MaxAge = -1
	c.Value = ""
	http.SetCookie(w, c)

	//best practice to cleaning session are using separate goroutine
	go cleanSomeSession()

	http.Redirect(w, req, "/", http.StatusSeeOther)
}
