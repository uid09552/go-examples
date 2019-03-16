package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var serverURL = "http://goexamplefrontendservice:8080"

func handleCounter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET params were:", r.URL.Query())
	cQuery := r.URL.Query().Get("count")
	counter := 1
	counter, err := strconv.Atoi(cQuery)

	if err != nil {
		counter = 1
	}
	respChan := make(chan string)
	result := []string{}
	client := http.Client{
		Timeout: time.Duration(500 * time.Millisecond),
	}
	for i := 0; i < counter; i++ {
		go func(c chan string) {
			resp, err := client.Get(serverURL)
			if err != nil {
				fmt.Println("error")
				errStr := err.Error()
				c <- fmt.Sprintf("error: %s", errStr)
			}
			if resp != nil {

				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)
				c <- string(body)
			}
		}(respChan)
		r := <-respChan
		result = append(result, r)
	}
	io.WriteString(w, fmt.Sprintf("Query Params: count %d Calling frontend %d time(s) with (timeout %d ms) response \n%s", counter, counter, client.Timeout/time.Millisecond, strings.Join(result, "\nresult: ")))

}

func main() {
	if os.Getenv("GO_EXAMPLE_SERVICE") != "" {
		serverURL = os.Getenv("GO_EXAMPLE_SERVICE")
		fmt.Println("server URL:", serverURL)
	}
	http.HandleFunc("/", handleCounter)
	http.ListenAndServe(":8080", nil)
	fmt.Println("server started")
}
