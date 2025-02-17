#!/bin/bash

#simple Password Generator
echo"Welcome to Password Generator"

sleep 2

echo"Please Enter The Length Of The Password:-"
read PASS_LENGTH

for p in $(seq 1 );
do
   openssl rand -base64 48 | cut -c1-$PASS_LENGTH
done
