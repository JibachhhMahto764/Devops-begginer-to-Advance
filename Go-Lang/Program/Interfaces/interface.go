package main
import"fmt"
type shape interface {
  area() float64
  perimeter() float64
  }
type square struct {
  side float64
  }
func(s square) area() float64 {
  return s.side * s.side
  }
func (s square) perimeter() float64 {
  return 4*s.side
  }
type rect struct {
  length,breath float64
  }
func(r rect) area() float64 {
  return r.lenght*r.breath
  }
func(r rect) perimeter() float64{
  return 2*r.lenght + 2*r.breath
  }
func printData(s shape) {
  fmt.PrintIn(s)
  fmt.PrintIn(s.area())
  fmt.PrintIn(s.perimeter())
  }
func main(){
  r := rect{length:3,breath:4}
  c:= square{side:5}
PrintData(r)
  PrintData(c)
  }
  
