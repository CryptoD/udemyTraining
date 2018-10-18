package main
import "fmt"


func funkcjapodrzedna(number int) (int, bool) {
return number/2, number%2 == 0
}

func main() {
  fmt.Println(funkcjapodrzedna(5))
}
