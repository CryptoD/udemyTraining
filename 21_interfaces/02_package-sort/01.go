package main
import "fmt"
import "sort"

type sortingjob []string
func (a sortingjob)      Len()          int      { return len(a) }
func (a sortingjob)      Swap(i, j int)          { a[i], a[j] = a[j], a[i] }
func (a sortingjob)      Less(i, j int) bool     { return a[i] < a[j] }

func main() {
var groupofnames = sortingjob{"Spiderman", "Batman", "Zorro", "Janosik"}
fmt.Println(groupofnames)
sort.Sort(groupofnames)
fmt.Println(groupofnames)

}
