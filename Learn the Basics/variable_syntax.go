package main //must need package declaration

import (
	"fmt"
) // every go program must have some object.

var name string = "fahim";

func sum( a int,  b int) int {
	return a+b;
}

func main() {

	fmt.Println("Hello, WOrld of Go")
	fmt.Println(name);
	n:=10
	fmt.Println(n);
	fmt.Println(sum(10,20));
}