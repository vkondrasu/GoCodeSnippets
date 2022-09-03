package main

import (
	"fmt"
	"io"
	"net/http"
)

func SimpleHandler(w http.ResponseWriter, req *http.Request) {

	io.WriteString(w, "<h1> hello, from GO web app </h1>")

}

const form = `<html><body><form action="#" method="post" name="bar">
<input type="text" name="in"/>
<input type="submit" value="Submit"/>
</form></html></body>`

func FormHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	switch req.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, "<h5>"+req.FormValue("in")+"</h5>")
	}

}

func main() {
	http.HandleFunc("/test1", SimpleHandler)
	http.HandleFunc("/test2", FormHandler)

	err := http.ListenAndServe("0.0.0.0:3000", nil)

	if err != nil {
		fmt.Println("Problem with starting the server")
	}
}
