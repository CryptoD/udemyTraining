package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	a, err := http.Get("https://www.google.com/")

	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(a.Body)
	a.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%t", b)
	fmt.Printf("%s", b)
}
