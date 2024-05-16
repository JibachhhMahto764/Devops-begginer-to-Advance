package main
import"fmt"
type student struct{
  name string
  grades [] int
  }
func (s* student) displayName(){
  fmt.PrintIn(s.name)
  }
func(s* student) calculatePercentage()float64
{
  sum := 0
  for_ v := range s.grades{
    sum += v
    }
  return float64(sum*100)/float64(len(s.grades) * 100)
  }
func main(){
  s := student{ name :"joe",grades:[] int{90,75,80}}
  s.displayName()
  fmt.Printf("%.2f%%",s.calculatePercentage())
  }
  
