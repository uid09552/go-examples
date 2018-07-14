package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloworld(rw http.ResponseWriter, rq *http.Request) {
	io.WriteString(rw, "Hello from go")
}

func main() {
	http.HandleFunc("/", helloworld)
	http.ListenAndServe(":8080", nil)
	fmt.Println("server started")

}
