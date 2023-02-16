package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {

	http.Handle("/test", new(countHandle))

	http.HandleFunc("/bar", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("test")
		fmt.Fprintf(writer, "hello, %q", request.URL.Path)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type countHandle struct {
	mu sync.Mutex //guard n
	n  int
}

func newCon() *countHandle {
	return &countHandle{
		n: 0,
	}
}

func (c *countHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.n++
	fmt.Fprintf(w, "Count is %d\n", c.n)
}
