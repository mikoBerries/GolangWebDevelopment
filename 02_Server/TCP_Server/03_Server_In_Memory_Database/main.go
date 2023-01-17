//a simple server that storing data in memory using map[string]value
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Panic(err)
	}

	handler(conn)
	fmt.Println("server exit program")
}

func handler(conn net.Conn) {
	fmt.Fprintln(conn, "Command: \n1.Get key\n2.SET key value\n3.Delete key")
	defer conn.Close()
	datas := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		inputdata := scanner.Text()
		//separate string to []string by white space
		mySliceString := strings.Fields(inputdata)

		switch strings.ToLower(mySliceString[0]) {
		case "get":
			fmt.Fprintln(conn, datas[mySliceString[1]])
		case "set":
			if len(mySliceString) >= 3 {
				datas[mySliceString[1]] = mySliceString[2]
			} else {
				fmt.Fprintln(conn, "Invalid input")
			}
		case "del":
			delete(datas, mySliceString[1])
		default:
			fmt.Fprintln(conn, "Invalid command")
		}
	}
}
