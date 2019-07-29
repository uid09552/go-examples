package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var requestURL = "http://127.0.0.1:8081"

// Data myexample response
type Data struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func helloworld(rw http.ResponseWriter, rq *http.Request) {
	dt := &Data{}
	resp, err := http.Get(requestURL)
	if err != nil {
		fmt.Println("failed to do request")
		io.WriteString(rw, "failed to do request")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println("failed to read body")
		io.WriteString(rw, "failed to read body")
		return
	}
	json.Unmarshal(body, &dt)
	var name string
	if value, ok := os.LookupEnv("name"); ok {
		name = value
	} else {
		name = "not set"
	}
	if val, err := json.Marshal(&dt); err == nil {
		fmt.Println(string(val))
		fmt.Fprintf(rw, "hello from backend, called from frontend: backend response: %s configmap: %s", val, name)
	} else {
		io.WriteString(rw, "failed in backend")
	}

}

func main() {
	http.HandleFunc("/", helloworld)
	http.ListenAndServe(":8080", nil)
	fmt.Println("server started on port 8080")

}
