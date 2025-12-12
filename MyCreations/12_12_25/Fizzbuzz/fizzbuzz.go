package main

import "fmt"

func main() {
	fizz_1()
	fizz_2()
	fizz_3()
}

func fizz_1() {
	for i := 1; i < 100; i++ {
		if i%5 == 0 && i%7 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("fizz")
		} else if i%7 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

	}
}
func fizz_2() {
	for i := 1; i < 100; i++ {
		if i%5 == 0 && i%7 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("fizz")
			continue
		}
		if i%7 == 0 {
			fmt.Println("buzz")
			continue
		}
		fmt.Println(i)
	}
}
func fizz_3() {
	for i := 1; i < 101; i++ {
		print_i := true
		if i%5 == 0 {
			fmt.Print("fizz")
			print_i = false
		}
		if i%7 == 0 {
			fmt.Print("buzz")
			print_i = false
		}
		if print_i {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
