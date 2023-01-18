# 01  Create a basic server using TCP.

The server should use net.Listen to listen on port 8080. 

Remember to close the listener using defer.

Remember that from the "net" package you first need to LISTEN, then you need to ACCEPT an incoming connection.
 
Now write a response back on the connection.

Use io.WriteString to write the response: I see you connected.

Remember to close the connection.

Once you have all of that working, run your TCP server and test it from telnet (telnet localhost 8080).


# 02 Building upon the code from the previous exercise:


In that previous exercise, we WROTE to the connection.

Now I want you to READ from the connection.

You can READ and WRITE to a net.Conn as a connection implements both the reader and writer interface.

Use bufio.NewScanner() to read from the connection.

After all of the reading, include these lines of code:

fmt.Println("Code got here.")
io.WriteString(c, "I see you connected.")

Launch your TCP server.

In your **web browser,** visit localhost:8080.

Now go back and look at your terminal.

Can you answer the question as to why "I see you connected." is never written?


# 03 Building upon the code from the previous exercise:

We are now going to get "I see you connected" to be written.

When we used bufio.NewScanner(), our code was reading from an io.Reader that never ended.

We will now break out of the reading.

Package bufio has the Scanner type. The Scanner type "provides a convenient interface for reading data". When you have a Scanner type, you can call the SCAN method on it. Successive calls to the Scan method will step through the tokens (piece of data). The default token is a line. The Scanner type also has a TEXT method. When you call this method, you will be given the text from the current token. Here is how you will use it:

```
scanner := bufio.NewScanner(conn)
for scanner.Scan() {
	ln := scanner.Text()
	fmt.Println(ln)
}
```

Use this code to READ from an incoming connection and print the incoming text to standard out (the terminal).

When your "ln" line of text is equal to an empty string, break out of the loop.

Run your code and go to localhost:8080 in **your browser.**

What do you find?

# 04 Building upon the code from the previous problem:

Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".

Pass the connection of type net.Conn as an argument into this function.

Add "go" in front of the call to "serve" to enable concurrency and multiple connections.



# 05 Building upon the code from the previous problem:

Before we WRITE our RESPONSE, let's WRITE to our RESPONSE the STATUS LINE and some RESPONSE HEADERS. Remember the request line and status line:

REQUEST LINE
GET / HTTP/1.1
method SP request-target SP HTTP-version CRLF
https://tools.ietf.org/html/rfc7230#section-3.1.1

RESPONSE (STATUS) LINE
HTTP/1.1 302 Found
HTTP-version SP status-code SP reason-phrase CRLF
https://tools.ietf.org/html/rfc7230#section-3.1.2

Write the following strings to the response - use io.WriteString for all of the following except the second and third:

"HTTP/1.1 200 OK\r\n"

fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))

fmt.Fprint(c, "Content-Type: text/plain\r\n")

"\r\n"

Look in your browser "developer tools" under the network tab. Compare the RESPONSE HEADERS from the previous file with the RESPONSE HEADERS in your new solution.

# 06 Building upon the code from the previous problem:

Print to standard out (the terminal) the REQUEST method and the REQUEST URI from the REQUEST LINE.

Add this data to your REPONSE so that this data is displayed in the browser.

# 07 Building upon the code from the previous problem:

Change your RESPONSE HEADER "content-type" from "text/plain" to "text/html"

Change the RESPONSE from "CHECK OUT THE RESPONSE BODY PAYLOAD" (and everything else it contained: request method, request URI) to an HTML PAGE that prints "HOLY COW THIS IS LOW LEVEL" in <h1> tags.

# 08 Building upon the code from the previous problem:

Add code to respond to the following METHODS & ROUTES:
	GET /
	GET /apply
	POST /apply

