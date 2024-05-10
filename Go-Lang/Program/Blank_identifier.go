blank identifier(underscore)=> '_'


package main
import "fmt"
func f() (int, int) {
    return 42, 53
}
func main() {
    a, b := f()
    fmt.Println(a, b)
  // print the results of f() using v1_
    v1_a, v1_b := f()
    fmt.Println(v1_a, v1_b)
}
