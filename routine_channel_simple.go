package main

import "fmt"

func messageRoutine(c chan string) {
	m := <-c
	fmt.Println(m)
}

func main() {
	x := make(chan string)
	go messageRoutine(x)
	x <- "hello"

	//go messageRoutine(x)

	close(x)

}
