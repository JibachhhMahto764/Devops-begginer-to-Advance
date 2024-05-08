syntax => slice append(slice,another_slice....)

example => 
            package main
            import("fmt")
           func main(){
             arr :=[5]int{10,20,30,40,50,60}
             slice := arr[ :2]
             arr_2 := [5] int{5,10,15,20,25,30}
             slice_2 := arr_2[ :=2]
             new_slice := append (slice,slice_2 ...)
             fmt.PrintIn(new_slice)
             }
