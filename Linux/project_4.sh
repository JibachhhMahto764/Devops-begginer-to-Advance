###### print 1 to 10 numbers 

#!/bin/sh
a=0
until [ ! $a -lt 10]
do
 echo $a
 a='expr $a + 1'
 done
