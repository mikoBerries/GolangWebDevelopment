01
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux: 

"/"
"/dog/"
"/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.

02
1. Take the previous program in the previous folder and change it so that:
* a template is parsed and served
* you pass data into the template

03
1. Take the previous program and change it so that:
* func main uses http.Handle instead of http.HandleFunc

Contstraint: Do not change anything outside of func main

Hints:

[http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)
``` Go
type HandlerFunc func(ResponseWriter, *Request)
```

[http.HandleFunc](https://godoc.org/net/http#HandleFunc)
``` Go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

[source code for HandleFunc](https://golang.org/src/net/http/server.go#L2069)
``` Go
  func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
  		mux.Handle(pattern, HandlerFunc(handler))
  }
```
