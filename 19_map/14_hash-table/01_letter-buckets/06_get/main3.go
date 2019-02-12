// Just testing prining directly from link
// Not working

package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print(http.Get("https://www.gutenberg.org/files/30407/30407-0.txt"))
}
