package main

import "fmt"


//For Loop
func forLoop() {
	sum:= 0;
	for i:= 1; i<=10;i++ {
		fmt.Println(i);
		sum+=i;
	}
	fmt.Println(sum);
}

//While Loop
func forLoop1() {
	sum:= 0;
	for sum<100 {
		sum+= sum;
	}
	fmt.Println(sum);
}

//for range
func forLoop2() {

	for i:= range 10 {
		fmt.Println(i)
	}

	var arr = []int{1,2,3,4,5};
	for i:= range arr {
		fmt.Println(i)
	}
}



func test() {
	fmt.Println("Hello Dev!! This is Loop");
}


