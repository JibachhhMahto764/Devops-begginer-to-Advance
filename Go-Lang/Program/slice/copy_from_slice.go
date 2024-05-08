syntax => 
       func copy(dst ,src[] type)int
num := copy (dest_slice,src_slice)



example => 
                 package main
                 import("fmt")
                func main(){
                  src_slice :=[]int{10,20,30,40,50}
                  dest_slice := make([] int,3)
                  num := copy(dest_slice,src_slice)
                  fmt.PrintIn(dest_slice)
                  fmt.PrintIn("number of element > copied:", num)
                  }
