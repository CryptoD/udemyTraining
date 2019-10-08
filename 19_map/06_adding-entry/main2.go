package main

import "fmt"

func main() {
	mapka := make(map[string]string)

	mapka["Firma 1"] = "HSBC"
	mapka["Firma 1"] = "Chase"
	mapka["Firma 2"] = "HSBC"
	mapka["Firma 3"] = "DB"
	fmt.Println(mapka)
}
