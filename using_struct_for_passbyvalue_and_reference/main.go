package main

import (
	"fmt"
)

//methodology is same as mentioned in the pass_by_value_vs_pass_by_reference file but only change is using struct
type pass struct{
	h int
	b string
}

func main() {
	b1:=pass{b:"hotragn"}
	b2:=pass{b:"fullstack"}
	fmt.Println("pass-by-value:")
	fmt.Println("value in main for pass by value is", b1.b)//b1.b refers to variable stored in struct pass-b1 with b
	fmt.Println("address in main for pass by value is", &b1.b)
	passbyvalue(b1)
	fmt.Println("pass-by-reference")
	fmt.Println("value in main for pass by ref is", b2.b)
	fmt.Println("address in main for pass by ref is", &b2.b)
	passbyref(&b2)//passing the address of the original value so that its the value in function is pointed to this
}
func passbyvalue(b1 pass) {
	b1.b=b1.b+"Pettugani"
	fmt.Println("the value of b1 in pass by value", b1.b)
	fmt.Println("the address in pass by value is ", &b1.b)
}
func passbyref(b2 *pass) {
	b2.b=b2.b+"dev"
	fmt.Println("the value of  in pass by ref", b2.b)
	fmt.Println("the address in pass by ref is ", &b2.b)
}

