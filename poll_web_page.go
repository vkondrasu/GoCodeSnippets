package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		log.Fatalf("Got : %v", err)
	}
}

func main() {
	flag.Parse()

	res, err := http.Get(flag.Arg(0))
	CheckError(err)

	data, err := ioutil.ReadAll(res.Body)

	CheckError(err)

	fmt.Printf("Got body: %q", string(data))
}
