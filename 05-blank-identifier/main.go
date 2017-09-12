package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// func invalidCode() {
// 	a := "a"
// 	b := "b"
// 	fmt.Println("a - ", a)
// 	// b is not used, invalid code
// }

// handle the error
func fetchWithErr() {
	res, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}

// ignore the error returned
//https://golang.org/doc/effective_go.html#blank
func fetchNoErr() {
	res, _ := http.Get("https://www.google.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}

func main() {
	// invalidCode()
	fetchWithErr()
	fetchNoErr()
}
