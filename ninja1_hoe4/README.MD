Types: 
go static types: 
go dynamic types: 

go user defined types are structs. struct member functions or behavior is supported through
interfaces.  C/C++ has structs and classes where classes support virtual
functions and structs dont. C/C++ has a pointer for this and supports class member functions
which don't exist in structs but you can include function pointers in a struct. 

go struct:

type Person struct{
  name string
  age int
  city,phone string
}

var p Person
p.name="propername"
p.age=100
p.city="here"
p.phone="23-33-333"

Now for the sucky part, there are many ways to initialize the p instance. 
add....


go interfaces:
