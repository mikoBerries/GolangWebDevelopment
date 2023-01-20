/*
	all func for session management to database
*/
package main

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

//simple data storange replace with DB conn
func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		//create session
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:   "session",
			Value:  sID.String(),
			MaxAge: sessionLeght,
		}

	}
	c.MaxAge = sessionLeght
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if s, ok := dbSessions[c.Value]; ok {
		//update db session lastActivity
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s

		u = dbUsers[s.uname]
	}
	return u
}

/*simple data storange replace with DB
return true if data found in session and user(logged)
return false if not found (not logged)
*/
func alreadyLogIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	// refresh session max age
	c.MaxAge = sessionLeght
	http.SetCookie(w, c)
	s, ok := dbSessions[c.Value]
	if ok {
		//update db session lastActivity
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.uname] //found in db user
	return ok
}

//func to deleting not used session in DB
func cleanSomeSession() {
	for k, s := range dbSessions {
		//now + sessionlenght in sec is after s.last active time return bool
		if time.Now().Add(sessionLeght).After(s.lastActivity) {
			delete(dbSessions, k)
		}
	}
	// dbSessionsCleaned = time.Now()
}
