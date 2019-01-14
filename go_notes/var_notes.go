package go_notes

import "fmt"

// go seems to be a cross between c and python. in C you need to declare a variable before assignment
// something like int c_var = 100; in python you can do python_var = 100 the python case wins the speed and
// ease of use case. go uses go_var :=100. You cant do the python style in go, go_var=100, it treats go_var as
// an undeclared var like what we see in C.

func main() {
	x := 100
	fmt.Println(x)
	//cant do y=1000. Will generate error messages
}
