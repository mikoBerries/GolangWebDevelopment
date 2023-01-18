//Hand on Exercise 02 (Readmme.md)
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error on net.Listen", err)
	}
	//closing listener
	defer listener.Close()
	for {
		//accept connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Error on listener.Accept", err)
		}
		// scanner := bufio.NewScanner(conn)

		// for scanner.Scan() {

		// 	ln := scanner.Text()
		// 	fmt.Println(ln)
		// 	if ln == "" { // when ln is empty, header is done
		// 		break
		// 	}
		// }

		// fmt.Println("Code got here.")
		// io.WriteString(conn, "I see you connected")
		// conn.Close()

		go callServe(conn)
	}
}

func callServe(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	//scanning header
	/*Example header
	GET / HTTP/1.1 +> first line method uri httpver
	Host: localhost:8080
	Connection: keep-alive
	Cache-Control: max-age=0
	sec-ch-ua: "Not_A Brand";v="99", "Google Chrome";v="109", "Chromium";v="109"
	sec-ch-ua-mobile: ?0
	sec-ch-ua-platform: "Windows"
	Upgrade-Insecure-Requests: 1
	User-Agent: Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36
	Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,;q=0.8,application/signed-exchange;v=b3;q=0.9
	Sec-Fetch-Site: cross-site
	Sec-Fetch-Mode: navigate
	Sec-Fetch-User: ?1
	Sec-Fetch-Dest: document
	Accept-Encoding: gzip, deflate, br
	Accept-Language: en-US,en;q=0.9

	*/
	i := 1

	var rMethod, rURI string
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 1 {
			split := strings.Split(ln, " ")
			rMethod = split[0]
			rURI = split[1]
			fmt.Println(">>>>>>>Method   :", rMethod)
			fmt.Println(">>>>>>>URI      :", rURI)
			fmt.Println(">>>>>>>HTTP ver :", split[2])
		} else {
			fmt.Println(ln)
			if ln == "" { // when ln is empty string, header is done
				break
			}
		}
		i++
	}
	// c.Write([]byte("this is your responds :" + time.Now().String()))
	// fmt.Fprintf(c, "this is your responds :"+time.Now().String())
	// io.WriteString(c, "this is your responds"+time.Now().String())

	switch {
	case rMethod == "GET" && rURI == "/":
		// handleIndex(c) do something get /
	case rMethod == "GET" && rURI == "/apply":
		// handleApply(c) do something get /apply
	case rMethod == "POST" && rURI == "/apply":
		// handleApplyPost(c) do something post /apply
	default:
		writeStatusLine(c)
	}
	// writeStatusLine(c)
}

//func to wrtie status line response
func writeStatusLine(c net.Conn) {
	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Code Gangsta</title>
	</head>
	<body>
		<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
	</body>
	</html>
`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	// fmt.Fprint(c, "Content-Type: text/plain\r\n")//response header for text/plain type
	fmt.Fprint(c, "Content-Type: text/html\r\n") //response header for html/plain type
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)

}
