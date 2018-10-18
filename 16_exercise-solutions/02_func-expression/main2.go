package main
import "fmt"
func main() {

n := func(xxx int) (int, bool) {
return xxx/2, xxx%2 == 0 }

fmt.Println(n(5))
}
// Here you use letter "n" and you make it a function.
// It's like func with the name n but inside function.
