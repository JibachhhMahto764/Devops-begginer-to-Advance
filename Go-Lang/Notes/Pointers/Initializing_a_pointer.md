First we have to know about pointers declaration ..
declaring a pointer syntax:
   var<pointer_name>*<data_type)
   such as; var ptr_i * int 


   Intializing a pointer....

    1. syntax1:  var<pointer_name>*<data_type> = &<variable_name>
     such as :-   
               i:=10
               var ptr_i * int = &i

    2. syntax2:- var<pointer_name> =& <variable_name>
    such as:-  
              s:= "hello"
              var ptr_s = &s

              shorthand declaration......
              <pointer_name>:=&<variable_name>
              s:="hello"
              ptr_s := &s
