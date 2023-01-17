//Making tcp server client to dial tcp server in 01_server
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	//dialing to listener server in 01_Server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic("error  net.Listen")
	}
	defer conn.Close()
	//writing to tcp server on localhost:8080
	fmt.Fprint(conn, "this is massage from 02_server")

	//reading all 01_server writing in conn
	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Panic("ioutil.ReadAll")
	}
	fmt.Println(string(bs))
}
