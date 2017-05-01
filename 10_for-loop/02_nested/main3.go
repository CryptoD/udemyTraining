package main // Lots of errors in this file. Ignore AM/PM, it sucks

import "fmt"

func main() {

	for h := 0; h <= 24; h++ {
		for m := 0; m <= 60; m++ {
			for s := 0; s <= 60; s++ {

				a := "A"

				if h <= 12 {
					var a = "AM"

				} else {
					var a = "PM"

				}
				fmt.Printf("\r %02d:%02d:%02d %s", h, m, s, a)

			}
		}
	}
}