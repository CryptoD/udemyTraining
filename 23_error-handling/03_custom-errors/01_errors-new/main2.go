package main

import "errors"
import "log"

func main() {
	_, err := Lols(1)
	if err != nil {
	log.Fatal(err)
}

}

func Lols(lolzzz float64) (float64, error) {

if lolzzz > 0 {
	return 0, 	errors.New("No co jest kurwa?")
}
return 100, nil
}