A variable is only available from inside the region it is created. This is called scope.
1. Global scope => Global variables are available from within any scope, global and local.
2. Local scope  => A variable created inside a function belongs to the local scope of that function, and can only be used inside that function. 

example=    
         calculation_to_hours= 24
         name_of_unit = "hours"
         my_var ="variable outside function"

    def days_to_units(num_of_days,custom_message):
      print(f"{num_of_days} days are {num_of_days * calculation_to_hours} {name_of_unit}")
      print(custom_message)
  
days_to_units(20,"Awesome!")
days_to_units(35,"Looks good!")


def scope_check(num_of_days):
 my_var ="variable inside function"
  print(name_of_unit)
  print(num_of_days)
  print(my_var)
  
  
  
scope_check(20)

print(my_var)
