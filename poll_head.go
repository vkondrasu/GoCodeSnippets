package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://www.google.co.in",
	"http://golang.org",
	"http://deliveroo.com",
}

func main() {
	for _, url := range urls {
		resp, err := http.Head(url)

		if err != nil {
			fmt.Println("error: ", url, err.Error())
		} else {
			fmt.Println(url, " : ", resp.Status)
		}

	}
}
