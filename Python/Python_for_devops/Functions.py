functions is defined using def keyword
block of code which only runs when it called
calling  a function = execution a function


calculation_to_units = 24
name_of_units = "hours"

def days_to_units(days):
    return days * calculation_to_units

print(f"20 days are {days_to_units(20)} {name_of_units}")
print("All good!")
