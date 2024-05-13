calculation_to_hours = 24
name_of_unit = "hours"
my_var = "variable outside function"

def days_to_units(num_of_days, custom_message):
    print(f"{num_of_days} days are {num_of_days * calculation_to_hours} {name_of_unit}")
    print(custom_message)
    input()

def scope_check(num_of_days):
    my_var = "variable inside function"
    print(name_of_unit)
    print(num_of_days)
    print(my_var)
    input()

# Example usage:
days_to_units(20, "Awesome!")
scope_check(5)
