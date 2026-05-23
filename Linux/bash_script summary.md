**List of learned commands:

chmod u+x hello_world.sh- give the file execution rights

whereis bash- view the location of the bash interpreter

whereis python3- view location of python3 interpreter

echo $0- display the name of the current shell

echo $SHELL- output the path to the current shell file

cat /etc/shells- list shells on this machine



Script text:

hello_world.sh file:

echo "Who am I? - " $USER

echo "Where am I? - " $(pwd)

echo "Where? - " $(uname)

echo "Where am I from? - " $HOME



file calendar.py:

#! /usr/bin/python3

# Program to display calendar of the given month and year

# importing calendar module

import calendar

yy = 2035 # year

mm = 11 # month

# To take month and year input from the user

# yy = int(input("Enter year: "))

# mm = int(input("Enter month: "))

# display the calendar

print(calendar.month(yy, mm)):**
