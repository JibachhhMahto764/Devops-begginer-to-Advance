#!/bin/bash

#simple Password Generator
echo"Welcome to Password Generator"

sleep 2        #this is the sleep program for 2 sec

echo"Please Enter The Length Of The Password:-"
read PASS_LENGTH    #input from user

for p in $(seq 1 );  # for loop and $ is used for varying the data 
do
   openssl rand -base64 48 | cut -c1-$PASS_LENGTH    # this is for generating the random password with base 64
done
