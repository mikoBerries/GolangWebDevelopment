//third party github julienschmidt - httprouter lib
package main

import "github.com/julienschmidt/httprouter"

func main() {
	mux := httprouter.New()
	mux.GET("/", nil)

}
