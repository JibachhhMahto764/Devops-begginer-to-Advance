-> User-defined datatype 
-> a structure that groups togehter data elements
-> Provide a way to reference a series of grouped values through a single variables name
-> Used when it makes sense to group or associate two or more data variabless.


declaration -->
    type<struct_name> struct{
    // list of fields 
    }

    such as => 
    type Student struct{
    name string
    rollNo int
    marks [] int
    grades map[string] int
    }
