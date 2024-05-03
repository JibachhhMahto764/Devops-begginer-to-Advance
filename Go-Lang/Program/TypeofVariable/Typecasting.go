type casting => the process of converting one data type to another is known as type casting.


Interger to float.....

package main
import "fmt"
func main() {
var i int = 90
var f float64 = float64(i)
fmt.printf("%2f \n",f)
  }


float to interger

package main
import "fmt"
func main(){
  var f float64 = 45.89
  var i int = int(f)
  fmt.printf("%v \n",i)
  }

Strcon...package

Itoa()
-> converts interger to string
-> returns one value-string formed with the given integer

pacakge main
import"fmt"
import"strconv"
func main(){
  var i int = 42 // into to string
  var s string = strconv.Itoa(i)
  fmt.printf("%q",S)
  }

Atoi()
-> converts string to interger
-> return two value values -the corresppanding
interger, error(if any)


package main
import(
  "fmt"
  "strconv")
func main(){
  var s string = "200"
  i,err := strconv.Atoi(s)
  fmt.printf("%v,%t \n",i,i)
  fmt.printf("%v,%t ", err, err)
  }


String to interger

package main
import 9
"fmt"
"str")
func main(){
  var s string = "200abc"
  i,err := strconv.Atoi(s)
  fmt.printf("%v,%T \n",i,i)
  fmt.printf("%v,%v",err,err)
  }
