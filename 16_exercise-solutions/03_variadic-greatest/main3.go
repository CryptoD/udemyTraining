package main
import "fmt"
func calculate(digits ...int) int {
var number1 int

for _, number2 := range digits {
if number2 > number1 {
number1 = number2
  }

}
return number1
}

func main() {
var number3 = calculate(2, 30, 1, 22, 0, 11, 2)
fmt.Println(number3)


}
