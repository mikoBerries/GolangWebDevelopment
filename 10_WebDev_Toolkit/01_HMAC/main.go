/*
HMAC ->hash message authentication code
Hashed message using special "code" stored in server side only not in client
*/
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
)

//specialHashKey are a secred string used to hashed our message
var specialHashKey string = "ourkey"

func main() {
	c := getCode("test@example.com")
	fmt.Println(c)
	//453159a62580804892cc90a27dfb8bfdf2309107336445dcfd0186674111ee71

	c = getCode("test@exampl.com")
	fmt.Println(c)
	//7e35ec99f8dfdd96d54c5185a81b25f662e2f786c36e737d4b65f903fe4862bc

	//a diffrent string will returning a diffrent result
	//so server authenticate message from this specify client/user
	//store real Hmac in database on group session data

	http.HandleFunc("/", foo)
	http.HandleFunc("/authenticate", auth)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	//check session cookie
	c, err := req.Cookie("session")
	if err != nil { // if err make new cookie to c
		c = &http.Cookie{
			Name:  "session",
			Value: "",
		}
	}

	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		c.Value = e + `|` + getCode(e)
		//set hashed email string to client cookie
	}

	// write new cookie with c
	http.SetCookie(w, c)

	io.WriteString(w, `<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	      <input type="email" name="email">
	      <input type="submit">
	    </form>
	    <a href="/authenticate">Validate This `+c.Value+`</a>
	  </body>
	</html>`)

}

//aut
func auth(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("session")
	if err != nil { //no cookie name "session" found (user undefined)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if c.Value == "" { //cookie name "session" value value nil
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//split cookie to get hashed email from client cookie
	xs := strings.Split(c.Value, "|")
	email := xs[0]
	codeRcvd := xs[1]

	//invalid hmac code
	//codeCheck := getCode(email + "s")

	//valid hmac code
	codeCheck := getCode(email)

	//check hash code from client side and server side if equal or not
	if codeRcvd != codeCheck {
		fmt.Println("HMAC codes didn't match")
		fmt.Println(codeRcvd)
		fmt.Println(codeCheck)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	//pass auth check
	io.WriteString(w, `<!DOCTYPE html>
	<html>
	  <body>
	  	<h1>`+codeRcvd+` - RECEIVED </h1>
	  	<h1>`+codeCheck+` - RECALCULATED </h1>
	  </body>
	</html>`)
}

//get hashed message with code specialHashKey ("ourkey")
func getCode(s string) string {
	h := hmac.New(sha256.New, []byte(specialHashKey))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
