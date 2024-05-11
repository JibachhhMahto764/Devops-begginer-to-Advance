1.what is pointer?
-> A pointer is a variable that holds memory address of another variables
-> they point where the memory is allocated and provide way to find or even change the value located at the memory location

#// Address and De-reference operators

1. & operator -> The address of a variable can be obtained by preceding the name of a variable with an ampersand sign(&), known as address-of operator.

2. * operator -> It is known as the de-reference operator. when placed before an address ,it returns the value at that address.
  
                                      memory address     memory
X := 77                                  0x0301           77

 &x = 0x0301

 *0x0301 = 77 
