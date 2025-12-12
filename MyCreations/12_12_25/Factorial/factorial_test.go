package main

import "fmt"

func ExampleFactorial() {
	fmt.Println(Factorial(2))
	fmt.Println(Factorial(4))
	fmt.Println(Factorial(6))
	fmt.Println(Factorial(7))
	fmt.Println(Factorial(15))
	fmt.Println(Factorial(16))

	//Output:
	//2
	//24
	//720
	//5040
	//1307674368000
	//20922789888000
}
func ExampleCursedFactorial() {
	fmt.Println(CursedFactorial(2))
	fmt.Println(CursedFactorial(4))
	fmt.Println(CursedFactorial(6))
	fmt.Println(CursedFactorial(7))
	fmt.Println(CursedFactorial(15))
	fmt.Println(CursedFactorial(16))

	//Output:
	//2
	//24
	//720
	//5040
	//1307674368000
	//20922789888000
}
