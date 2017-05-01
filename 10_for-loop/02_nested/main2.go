package main

import (
	"fmt"
	"time"
)

func main() {

	for h := 0; h < 60; h++ {

		for m := 0; m < 60; m++ {

			for s := 0; s < 60; s++ {

				fmt.Printf("\r%02d:%02d:%02d", h, m, s)
				time.Sleep(1 * time.Second)
			}
		}
	}

}
