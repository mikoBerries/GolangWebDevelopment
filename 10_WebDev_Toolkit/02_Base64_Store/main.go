/*
Base64 encode for storing string
https://www.rfc-editor.org/rfc/rfc4648
Base64 Encoding used for storing data that banned some special char
Base64url Encoding with URL and Filename Safe Alphabet

Base encoding visually hides otherwise easily recognized information
such as passwords, but does not provide any computational confidentiality
Example : storing hashed cookie ,URL
*/
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

func main() {
	//example in cookie hash value
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, "")
	temp := h.Sum(nil)
	fmt.Println(temp)
	b := base64.StdEncoding.EncodeToString(temp)
	fmt.Println(b)
	c, err := base64.StdEncoding.DecodeString(b)

	if err != nil {
		log.Fatalln("err decode :", err)
	}
	fmt.Println(c)

}
