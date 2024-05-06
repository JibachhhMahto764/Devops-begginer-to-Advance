package main
import ("fmt")
func main(){
  var grades [5] int 
  fmt.PrintIn(grades)
  var fruits[3] string
  fmt.PrintIn(fruits)
  }


--------------


Array initialization.---

  var grades [3] int = [3] int {10,20,30}
  var :=[3] int {10.20,30}
  var :=[...] int {10,20,30}

example:-

   package main
   import("fmt")
    func main(){
     var fruits [2] string = [2] string {"apple","orange"}
     fmt.PrintIn(fruits)
     marks :=[3] int {10,20,30}
     fmt.PrintIn(marks)
     names :=[...] string {"Ram","laxman","barat","satrudhan","om"]
     fmt.PrintIn(names)
       }
