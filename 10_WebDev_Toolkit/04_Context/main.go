/*
	understanding usage of context to signaling some process to end
	so we'll not having a leak go routine
*/
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

	//example for ctx usage
	//context.backgorund() returning new empty context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // make sure all paths cancel the context to avoid context leak

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 { //after <-chan output are 5
			cancel()
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

//generating <-chan value of n++
func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return // avoid leaking of this goroutine when ctx is done.
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}

func foo(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()

	//// Use context Values only for request-scoped data that transits processes and
	// APIs, not for passing optional parameters to functions.
	// context.
	ctx = context.WithValue(ctx, "mykey", 100000)
	ctx = context.WithValue(ctx, "another data", "Bond")

	ctx, cancelctx := context.WithCancel(ctx)
	defer cancelctx() // releases resources if slowOperation completes before timeout elapses

	log.Println(ctx)
	fmt.Fprintln(w, ctx)
	//printed context
	//result:
	//context.Background
	//.WithValue(type *http.contextKey, val <not Stringer>)
	//.WithValue(type *http.contextKey, val [::1]:8080)
	//.WithCancel.WithCancel
	// WithValue(type string, val <not Stringer>)
	// .WithValue(type string, val Bond)
}

func dbAccess(ctx context.Context, c context.CancelFunc) (int, error) {
	defer c()
	ch := make(chan int)
	go func() {
		// ridiculous long running task
		uid := ctx.Value("mykey").(int)
		time.Sleep(10 * time.Second)

		// check to make sure we're not running in vain
		// if ctx.Done() has close
		if ctx.Err() != nil {
			return
		}

		ch <- uid
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = context.WithValue(ctx, "mykey", 100000)
	ctx, cancelctx := context.WithTimeout(ctx, 3*time.Second)

	results, err := dbAccess(ctx, cancelctx)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Fprintln(w, results)
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
