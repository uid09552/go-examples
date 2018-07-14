package main

import (
	"fmt"
	"net/http"
	"time"
)

type list []struct {
}

func main() {
	links := []string{"https://google.com", "https://facebook.com", "https://amazon.com", "https://gmx.de"}
	cn := make(chan string)
	for _, link := range links {
		fmt.Println(link)
		go checkLink(link, cn)
	}
	for l := range cn {
		go func(l string) {
			time.Sleep(time.Second)
			checkLink(l, cn)
		}(l)
	}
	for {
		checkLink(<-cn, cn)
	}

}

func checkLink(link string, cn chan string) {
	fmt.Println("getting link: ", link)
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down")
		cn <- link
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println(link + " statuscode seems not to be 200")
		cn <- link
		return
	}
	fmt.Printf("%+v", resp)
	cn <- link
}
