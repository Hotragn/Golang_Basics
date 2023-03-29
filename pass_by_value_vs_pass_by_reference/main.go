package main

import (
	"fmt"
)

func main() {
	var h1=4 //h1 is initialized to use for pass by value
	h2:= "fullstack" //h2 string is used for pass by reference
	updatedh2:=&h2 // points to memory address of original value
	//in pass by value the value is copied and stored for another memory location(address)
	fmt.Println("Pass-By-Value:")
	fmt.Println("value in main for-pass by value:", h1,"& memory location in main-pass by value:", &h1)
	passbyvalue(h1)// passing the value of h1 to h passbyvalue function which will copy it
	//everytime we pass a value to a function go creates a copy of it

	//in pass by reference the value is not passed rather the address is passed such that whenever something happens to that variable it is happened to original variable too
	fmt.Println("Pass-By-Reference:")
	fmt.Println("value in main for before-pass by ref:", h2,"& memory location in main-pass by ref:", &h2)
	passbyref(updatedh2)
	fmt.Println("value in main after-pass by ref:", h2,"& memory location in main after-pass by ref:", &h2) //original value is updated 
}

func passbyvalue(h int) {
	p := 6
	h = h + p
	fmt.Println("value pass-by-value :", h,"& memory location pass-by-value: ", &h)
}

func passbyref(h2 *string) {// makes a copy and points it to original value
    var p = "dev"
	*h2 = *h2 + p//so this will be pointing to updatedh2 and when we update value the original is also updated
	fmt.Println("value pass-by-ref :", h2,"& memory location pass-by-ref: ", &h2)
}