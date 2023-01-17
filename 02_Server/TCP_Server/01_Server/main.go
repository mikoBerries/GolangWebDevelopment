/*This main.go file to understanding server on golang
//TCP Transmision Control Protocol - HTTP Hyper Text Transfer Protocol
//HTTP are proctocol that run using TCP
//HTTP Request => request line,Headers,Message body (optional)
request line =>Method SP Request-URI SP HTTP-Version CRLF (GET - path/to/file/index.html - /HTTP/1.0 )
//HTTP Response => Status line,Headers,MEssageBody (optional)
status line => HTTP-Version SP Status-Code SP Reason-Phrase CRLF (HTTP/1.0 - 200 - OK))
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var connTimeout time.Duration = 10

func main() {
	listener, err := net.Listen("tcp", ":8080") //tcp network , address port":8080"
	if err != nil {
		log.Panic("err net.Listen :", err)
	}
	defer listener.Close()

	for {
		//Cmd => telnet localhost 8080
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("err listener.Accept :", err)
		}
		//send write data to connection
		go sendStringToConection(conn)
		//listen data from connection
		go handler(conn)

	}

}

func sendStringToConection(conn net.Conn) {
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n") // status line
	// fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) 	// header
	fmt.Fprint(conn, "Content-Type: text/plain\r\n") // header
	// io.WriteString(conn, "\r\n") 							// blank line; CRLF; carriage-return line-feed
	// io.WriteString(conn, body) 								// body, aka, payload

	io.WriteString(conn, "\nHello from TCP server")
	fmt.Fprintln(conn, "how is your day")
	fmt.Fprintf(conn, "%v", "WEL, i hope!")

	// conn.Close()
}

//scanning connection
func handler(conn net.Conn) {
	// s := "this is pe ,two twe\n \t asdasd a\n asdas"
	// scanner := bufio.NewScanner(strings.NewReader(s))
	// scanner.Split(bufio.ScanRunes)
	scanner := bufio.NewScanner(conn)

	//set connection timeout
	err := conn.SetDeadline(time.Now().Add(connTimeout * time.Second))
	if err != nil {
		log.Fatalln("err connection timeout :", err)
	}
	for scanner.Scan() {
		line := scanner.Text()
		//printing in server side
		fmt.Println("Scanned from conn:", line)
		//printing in conection side (client)
		fmt.Fprintf(conn, "I heard you say :%s \n", line)
	}
	conn.Close()
	fmt.Println("exit handler func()")
}
