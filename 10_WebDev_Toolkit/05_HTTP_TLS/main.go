/*
how to use HTTPS instead HTTP
creating certificate and key
*/
package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

//SSL TLS HTTPS process explained in 7 minutes
//https://www.youtube.com/watch?v=4nGrOpo0Cuc&ab_channel=JohannesBickel
//TLS Hello
//https://www.cloudflare.com/learning/ssl/what-happens-in-a-tls-handshake/

// Go to https://localhost:10443/ or https://127.0.0.1:10443/
// list of TCP ports:
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers

// Generate unsigned certificate
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com
// for example
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost

// WINDOWS
// windows may have issues with go env GOROOT
// go run %(go env GOROOT)%/src/crypto/tls/generate_cert.go --host=localhost

// instead of go env GOROOT
// you can just use the path to the GO SDK
// wherever it is on your computer
