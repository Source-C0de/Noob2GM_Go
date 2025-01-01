
package main
import "fmt"

func greatestNumber() {
	var a,b,c int;
	a = 10;
	b = 20;
	c = 5;
	if a>b && a>c {
		fmt.Printf("%d is greatest number",a);
	}else if b>a && b>c {
		fmt.Printf("%d is greatest number",b);

	}else{
		fmt.Printf("%d is greatest number",c);
	}
}

func evenOdd() {
	n := 101
	if n%2==0 {
		fmt.Printf("%d is even number",n);
	}else{
		fmt.Printf("%d is odd number",n);
	}
}

func currentTime() {

}
func main_test() {
	evenOdd();
	greatestNumber();
}