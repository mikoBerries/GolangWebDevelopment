/*mongo DB connection*/
package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"m.test/11_MongoDB/01_MVC/controler"
)

// var uc userControler.userControler

func main() {
	uc := controler.NewUserControler()

	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/createUser", uc.CreateUser)
	r.DELETE("/deleteUser", uc.DeleteUser)
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "done")
}
