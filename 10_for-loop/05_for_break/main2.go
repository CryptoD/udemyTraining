package main
import "fmt"
func main() {
  a := 100

  for {
    a++
    fmt.Println(a)
    if a == 200 {
      break
    }
  }
}
