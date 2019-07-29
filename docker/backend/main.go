package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//Data my example data
type Data struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func handleBackend(rw http.ResponseWriter, rq *http.Request) {
	dt := Data{Name: "hello", Value: "myvalue "}
	out, err := json.Marshal(dt)
	if err != nil {
		fmt.Println("Error in converting output")
	}
	fmt.Println(string(out))
	io.WriteString(rw, string(out))
}

func main() {
	fmt.Println("Starting Backend")
	http.HandleFunc("/", handleBackend)
	http.ListenAndServe(":8081", nil)
}
