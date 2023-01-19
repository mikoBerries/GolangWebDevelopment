/*
	Undestanding Session
	common use id in session using UUID
*/
package main

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", foo)
	// http.HandleFunc("/set", set)
	// http.HandleFunc("/read", read)
	// http.HandleFunc("/readCookies", readCookies)
	// http.HandleFunc("/expire", expire)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	ID := uuid.NewV4()
	myCookie := http.Cookie{
		Name:  "session",
		Value: ID.String(),
		// Secure: true,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &myCookie)
}

// func readCookies(w http.ResponseWriter, req *http.Request) {
// 	//getting all cookies []
// 	//req.cookie(Cookie.Name) for secific cookies name
// 	cookies := req.Cookies()
// 	io.WriteString(w, "total cookies :"+strconv.Itoa(len(cookies))+"\n")
// 	//reading all cookies
// 	for _, val := range cookies {
// 		io.WriteString(w, val.Name+" : "+val.Value+"\n")
// 		fmt.Println(val.Name + " : " + val.Value + "\n")
// 	}
// }

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
