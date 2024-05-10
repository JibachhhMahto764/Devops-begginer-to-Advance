-> function that accepts variable number of arguments
-> it is possible to pass a varying number of arguments of the same type  as referenced in the function signature
-> to declare a variadic functions the type of the final parameter is preceded by a ellipsis "..."


syntax =>
          func<func_name>(param_1 type,param_2,param_3 ... type)<return_type>

          such as => 1. func SumNumbers(numbers...int) int
                     2. func SumNumbers(str string, numbers...int)
                     
