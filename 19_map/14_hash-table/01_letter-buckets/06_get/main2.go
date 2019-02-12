package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, res2 := http.Get("https://www.gutenberg.org/files/30407/30407-0.txt")
	if res2 != nil {
		log.Fatal(res2)
	}
	bs, res2 := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if res2 != nil {
		log.Fatal(res2)
	}
	fmt.Printf("%s", bs)
}
